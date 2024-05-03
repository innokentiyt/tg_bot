package main

// === telegram structs ===

type Chat struct {
	ID int `json:"id"`
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
	MessageReaction MessageReactionUpdated `json:"message_reaction"`
}

// ==== local save structs 

type ReactionData struct {
	Type string `json:"type"`
	Count int `json:"count"`
	ID int `json:"id"`
}

type UserData struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Reactions []ReactionData `json:"reactions"`
}

type UsersSave struct {
	Users []UserData `json:"users"`
}
