package main

// === telegram structs ===


//
   //{
   //   "update_id": 579421648,
   //   "message": {
   //     "message_id": 1214302,
   //     "from": {
   //       "id": 1683199658,
   //       "is_bot": false,
   //       "first_name": "Андрей",
   //       "last_name": "Солженицын",
   //       "username": "magic_frontier"
   //     },
   //   }
   // },

type Chat struct {
	ID int `json:"id"`
}

type MessageEntity struct {
	Type string `json:"type"`
}

type ReplyMessage struct {
	MessageID int `json:"message_id"`
	From User `json:"from,omitempty"`
}

type Message struct {
	MessageID int `json:"message_id"`
	From User `json:"from,omitempty"`
	Entities []MessageEntity `json:"entities,omitempty"`
	Text string `json:"text,omitempty"`
	Chat Chat `json:"chat"`
	ReplyMsg ReplyMessage `json:"reply_to_message,omitempty"`
}

type User struct {
	ID int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
}

type ReactionType struct {
	Type string `json:"type"`
	Emoji string `json:"emoji"`
}

type MessageReactionUpdated struct {
	Chat Chat `json:"chat"`
	Message_id int `json:"message_id"`
	User User `json:"user,omitempty"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

type Update struct {
	ID int `json:"update_id"`
	MessageReaction MessageReactionUpdated `json:"message_reaction,omitempty"`
	Message Message `json:"message,omitempty"`
}

// ==== local save structs 

type UsersData struct {
	Users map[string]UserInfo `json:"Users"`
}

type UserInfo struct {
	Gratz int    `json:"gratz"`
	Name  string `json:"name"`
}

// ==== LLM structs

type LLM_Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type LLM_Messages struct {
	Model string `json:"model"`
	Messages []LLM_Message `json:"messages"`
}

type LLM_Answer struct {
	Choices [] struct {
		Index int `json:"index"`
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
