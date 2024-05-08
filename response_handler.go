package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var reaction_ok string = "👌"
var reaction_clown string = "🤡"
var twenty_last_messages [20]Message
var twenty_last_messages_iterator int = 0

func generateReaction(emoji string) []ReactionType {
	return []ReactionType {
		{"emoji", emoji},
	}
}

func processGratzMsg(u Update) {
	user := u.Message.ReplyMsg.From
	send_user := u.Message.From

	if user.ID == send_user.ID {
		err := setMessageReaction(u.Message.MessageID, u.Message.Chat.ID, generateReaction(reaction_clown))
		if err != nil {
			fmt.Println("Error sending reaction", err)
		}
		return
	}
	appendGratz(user)
	err := setMessageReaction(u.Message.MessageID, u.Message.Chat.ID, generateReaction(reaction_ok))
	if err != nil {
		fmt.Println("Error sending reaction", err)
	}
}

func processTopMsg(u Update) {
	var msg string
	type Top struct {
		UserName string
		Amount int
	}
	var users []Top
	for _, user := range my_db.Users {
		users = append(users, Top{user.Name, user.Gratz})
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Amount > users[j].Amount
	})

	for i, user := range users {
		msg += strconv.Itoa(i) + ".  " + user.UserName + ": " + strconv.Itoa(user.Amount) + "\n"
	}
	sendMessage(u.Message.Chat.ID, msg)
}

func processUpdates(updates []Update, channel_id int) {
	for _, update := range updates {
		if update.Message.Chat.ID != channel_id {
			continue
		}
		
		fmt.Println(update.Message.Text)
		if update.Message.Text == "грац" {
			processGratzMsg(update)
			return
		}
		idx := slices.IndexFunc(update.Message.Entities, func(m MessageEntity) bool {
			return m.Type == "bot_command"
		})
		if idx < 0 {
			processNonCommandUpdate(update.Message)
			continue
		}
		if strings.Contains(update.Message.Text, "gratz") {
			processGratzMsg(update)
			return
		}
		if strings.Contains(update.Message.Text, "top") {
			processTopMsg(update)
			return
		}
	}
}

func processNonCommandUpdate(msg Message) {
	if len(msg.Text) == 0 {
		return
	}
	pushToTwentyLastMessages(msg)
	if msg.ReplyMsg.MessageID != 0 && msg.ReplyMsg.From.IsBot {
		// this is a reply to a bot
		sendLLMAnswer(msg)
		return
	}
	var probability float32 = 0.01 // 1% chance
	random_value := rand.Float32()
	if random_value >= probability {
		return
	}
	sendLLMAnswer(msg)
}

func pushToTwentyLastMessages(msg Message) {
	twenty_last_messages_iterator += 1
	if twenty_last_messages_iterator > len(twenty_last_messages) - 1 {
		twenty_last_messages_iterator = 0
	}
	twenty_last_messages[twenty_last_messages_iterator] = msg
}
