package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var okReaction string = "👌"
var clownReaction string = "🤡"

func generateReaction(emoji string) []ReactionType {
	return []ReactionType {
		{"emoji", emoji},
	}
}

func processGratzMsg(u Update) {
	user := u.Message.ReplyMsg.From
	send_user := u.Message.From

	if user.ID == send_user.ID {
		err := setMessageReaction(u.Message.MessageID, u.Message.Chat.ID, generateReaction(clownReaction))
		if err != nil {
			fmt.Println("Error sending reaction", err)
		}
		return
	}
	appendGratz(user)
	err := setMessageReaction(u.Message.MessageID, u.Message.Chat.ID, generateReaction(okReaction))
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
