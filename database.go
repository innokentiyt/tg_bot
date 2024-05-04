package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var my_db UsersData

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

func appendGratz(u User) {
	str_id := strconv.Itoa(u.ID)
	val, ok := my_db.Users[str_id]
	if !ok {
		my_db.Users[str_id] = UserInfo{1, u.FirstName}
		saveDatabase()
		return
	}
	val.Gratz += 1
	val.Name = u.FirstName
	my_db.Users[str_id] = val
	saveDatabase()
}
