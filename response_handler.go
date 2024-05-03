package main

import (
	"fmt"
	"slices"
)

func processUpdates(updates []Update, channel_id int) {
	var messages_to_check []MessageQuery
	for _, update := range updates {
		if update.MessageReaction.Chat.ID != channel_id {
			continue
		}
		query := MessageQuery{
			update.MessageReaction.Message_id,
			update.MessageReaction.NewReaction,
		}
		messages_to_check = append(messages_to_check, query)
	}

	for _, query := range messages_to_check {
		idx := slices.IndexFunc(updates, func(u Update) bool {
			return u.Message.MessageID == query.message_id
		})
		if idx < 0 {
			fmt.Println("Tried searching", query.message_id, "but it was unavailable")
			continue
		}
		user := updates[idx].Message.From
		addToDatabase(user, query.reactions, query.message_id)
	}

	for _, update := range updates {
		if update.Message.Chat.ID != channel_id {
			continue
		}
		idx := slices.IndexFunc(update.Message.Entities, func(m MessageEntity) bool {
			return m.Type == "bot_command"
		})
		if idx < 0 {
			continue
		}
		fmt.Println(update.Message.Text)
	}
}
