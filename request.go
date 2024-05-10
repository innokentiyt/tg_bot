package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func requestUpdates(offset int) (error, []Update) {
	var buf bytes.Buffer
	params := map[string]string {
		"offset":  strconv.Itoa(offset),
		"timeout": strconv.Itoa(10),
	}
	allowed_updates := []string {
		"message",
	}
	data, err := json.Marshal(allowed_updates)
	if err != nil {
		fmt.Println("Error parsing json:", err)
		return err, nil
	}
	params["allowed_updates"] = string(data)

	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(params)
	if err != nil {
		return err, nil
	}
	url := "https://api.telegram.org/bot"+ bot_token +"/getUpdates"
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return err, nil
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	if resp.StatusCode != 200 {
		return err, nil
	}
	var tg_response struct {
		Result []Update
	}
	err = json.Unmarshal(body, &tg_response)
	if err != nil {
		fmt.Println("Error serializing response:\n", err)
		return err, nil
	}
	//fmt.Println(string(body))
	return nil, tg_response.Result
}

func setMessageReaction(message_id int, chat_id int, reaction []ReactionType) error {
	var buf bytes.Buffer
	params := map[string]string {
		"chat_id":  strconv.Itoa(chat_id),
		"message_id": strconv.Itoa(message_id),
	}
	data, err := json.Marshal(reaction)
	if err != nil {
		return err
	}
	params["reaction"] = string(data)

	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(params)
	if err != nil {
		return err
	}
	url := "https://api.telegram.org/bot"+ bot_token +"/setMessageReaction"
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func sendMessage(chat_id int, text string) error {
	var buf bytes.Buffer
	params := map[string]string {
		"chat_id":  strconv.Itoa(chat_id),
		"text": text,
	}
	encoder := json.NewEncoder(&buf)
	err := encoder.Encode(params)
	if err != nil {
		return err
	}
	url := "https://api.telegram.org/bot"+ bot_token +"/sendMessage"
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func sendLLMAnswer(msg Message) {
	url := "https://api.openai.com/v1/chat/completions"

	msgs := LLM_Messages{
		Model: "gpt-3.5-turbo",
	}
	msgs.Messages = append(msgs.Messages, LLM_Message{Role: "system", Content: "Ты в чате друзей. Отвечай как своим друзьям."})
	msgs.Messages = append(msgs.Messages, LLM_Message{Role: "user", Content: msg.Text})

	json_data, err := json.Marshal(msgs)
	if err != nil {
		fmt.Println("[sendLLMAnswer] Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println("[sendLLMAnswer] Error creating HTTP request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("openai_api_key"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[sendLLMAnswer] Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[sendLLMAnswer] Error reading HTTP request:", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("[sendLLMAnswer] Request Status != 200", string(body))
		return
	}
	var response LLM_Answer
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("[sendLLMAnswer] Error serializing response:", err)
		return
	}
	if len(response.Choices) == 0 {
		fmt.Println("[sendLLMAnswer] Choice is empty:")
		return
	}

	output := response.Choices[0].Message.Content

	// TODO: reply to a message
	err = sendMessage(msg.Chat.ID, output)
	if err != nil {
		fmt.Println("[sendLLMAnswer] Error sending msg to tg:", err)
	}
}
