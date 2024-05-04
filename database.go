package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
)

var my_db Database

// ====== temp database

func readDatabase() error {
	content, err := os.ReadFile("./database.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &my_db)
	if err != nil {
		return err
	}
	return nil
}

func addToDatabase(user User, reactions []ReactionType, message_id int) {
	idx := slices.IndexFunc(my_db.Users, func(u UserData) bool {
		return u.ID == user.ID
	})
	if idx < 0 {
		fmt.Println("Tried searching", user.ID, "but it was unavailable")
		my_db.Users = append(my_db.Users, UserData{
			user.ID,
			user.FirstName,
			[]MessageData {
				{
					reactions,
					message_id,
				},
			},
		})
		saveDatabase()
		return
	}
	r_idx := slices.IndexFunc(my_db.Users[idx].Messages, func(r MessageData) bool {
		return r.MessageId == message_id
	})
	if r_idx < 0 {
		my_db.Users[idx].Messages = append(my_db.Users[idx].Messages, MessageData {reactions, message_id})
		saveDatabase()
		return
	}
	my_db.Users[idx].Messages[r_idx] = MessageData {reactions, message_id}
	saveDatabase()
}

func saveDatabase() {
	json, err := json.Marshal(my_db)
	if err != nil {
		fmt.Println("Error saving db:", err)
		return
	}
	err = os.WriteFile("database.json", json, 0644)
	if err != nil {
		fmt.Println("Error saving db:", err)
	}
}

func getSortedTopListAsString(emoji string) string {
	var result string

	count_map := countEmojiReactions(my_db.Users, emoji)

	type SortStruct struct {
		Key string
		Value int
	}
	var emoji_reactions []SortStruct

	for k,v := range count_map {
		emoji_reactions = append(emoji_reactions, SortStruct{k,v})
	}

	sort.Slice(emoji_reactions, func(i, j int) bool {
		return emoji_reactions[i].Value > emoji_reactions[j].Value
	})
	for i,s := range emoji_reactions {
		result += strconv.Itoa(i+1) + ". " + s.Key + ": " + strconv.Itoa(s.Value) + "\n"
	}

	return result
}

func countEmojiReactions(users []UserData, emoji string) map[string]int {
	emojiReactionCount := make(map[string]int)

	for _, user := range users {
		count := 0
		for _, message := range user.Messages {
			for _, reaction := range message.Reactions {
				if reaction.Type == "emoji" && reaction.Emoji == emoji {
					count++
				}
			}
		}
		emojiReactionCount[user.Name] = count
	}

	return emojiReactionCount
}
