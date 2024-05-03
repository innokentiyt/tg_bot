package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func requestReactions(bot_key string) (error, []Update) {
	var buf bytes.Buffer
	offset := 0
	params := map[string]string {
		"offset": strconv.Itoa(offset),
	}
	allowed_updates := []string {
		"message_reaction",
		//"message_reaction_count",
	}
	data, err := json.Marshal(allowed_updates)
	if err != nil {
		fmt.Println("Error parsing json:", err)
	}
	params["allowed_updates"] = string(data)

	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(params)
	if err != nil {
		return err, nil
	}
	url := "https://api.telegram.org/bot"+ bot_key +"/getUpdates"
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
	fmt.Println(string(body))
	return nil, tg_response.Result
}
