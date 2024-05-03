package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
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
