package main

func processUpdates(updates []Update, channel_id int) {
	for _, update := range updates {
		if update.MessageReaction.Chat.ID != channel_id {
			continue
		}
		update.MessageReaction.
	}
}
