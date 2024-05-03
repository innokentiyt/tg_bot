package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	bot_token := os.Getenv("tg_bot_api")
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

	for {
		err, updates := requestUpdates(bot_token)
		if err != nil {
			fmt.Println("Some error happened:\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		processUpdates(updates, channel_id)

		time.Sleep(5 * time.Second)
	}
}
