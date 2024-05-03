package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	bot_key := os.Getenv("tg_bot_api")
	channel_id, err := strconv.Atoi(os.Getenv("tg_channel_id"))
	if err != nil {
		fmt.Println("Error parsing env variable:", err)
		return
	}
	for {
		err, updates := requestReactions(bot_key)
		if err != nil {
			fmt.Println("Some error happened:\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		processUpdates(updates, channel_id)

		time.Sleep(5 * time.Second)
	}
}
