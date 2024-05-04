package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var bot_token string

func main() {
	bot_token = os.Getenv("tg_bot_api")
	channel_id, err := strconv.Atoi(os.Getenv("tg_channel_id"))
	if err != nil {
		fmt.Println("Error parsing env variable:", err)
		return
	}
	err = readDatabase()
	if err != nil {
		fmt.Println("Error reading temp database:", err)
		saveDatabase()
	}
	offset := 0
	for {
		err, updates := requestUpdates(offset)
		if err != nil {
			fmt.Println("Some error happened:\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if len(updates) > 0 {
			processUpdates(updates, channel_id)
			last_index := len(updates) - 1
			last_element := updates[last_index]
			offset = last_element.ID + 1
		}
		time.Sleep(5 * time.Second)
	}
}
