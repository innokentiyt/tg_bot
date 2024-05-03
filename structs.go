package main

type Chat struct {
	ID int64 `json:"id"`
}

type User struct {
	ID int64 `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
}

type ReactionType struct {
	Type string `json:"type"`
	Emoji string `json:"emoji"`
}

type MessageReactionUpdated struct {
	Chat Chat `json:"chat"`
	Message_id int64 `json:"message_id"`
	User User `json:"user,omitempty"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

type Update struct {
	ID int64 `json:"update_id"`
	MessageReaction MessageReactionUpdated `json:"message_reaction"`
}
