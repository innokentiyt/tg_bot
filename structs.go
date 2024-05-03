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

type Message struct {
	MessageID int `json:"message_id"`
	From User `json:"from,omitempty"`
	Entities []MessageEntity `json:"entities,omitempty"`
	Text string `json:"text,omitempty"`
	Chat Chat `json:"chat"`
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
	ID int64 `json:"update_id"`
	MessageReaction MessageReactionUpdated `json:"message_reaction,omitempty"`
	Message Message `json:"message,omitempty"`
}

// ==== local save structs 

type MessageData struct {
	Reactions []ReactionType `json:"reactions"`
	MessageId int `json:"message_id"`
}

type UserData struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Messages []MessageData `json:"messages"`
}

type Database struct {
	Users []UserData `json:"users"`
}

// ====== internal

type MessageQuery struct {
	message_id int
	reactions []ReactionType
}

