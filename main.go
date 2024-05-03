package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	bot_key := os.Getenv("tg_bot_api")
	for {
		err, updates := requestReactions(bot_key)
		if err != nil {
			fmt.Printf("Some error happened:\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		processUpdates(updates)
	}
}
