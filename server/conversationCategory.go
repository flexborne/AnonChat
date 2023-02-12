package main

import (
	"math/rand"
	"time"
)

// ConversationCategory defines the different conversation categories for clients
type ConversationCategory int

const (
	Casual ConversationCategory = iota
	Serious
	Flirtatious
)

var conversationCategory = [...]string{
	"Casual",
	"Serious",
	"Flirtatious",
}

func (i ConversationCategory) String() string {
	return conversationCategory[i]
}

var allConversationCategories = []ConversationCategory{
	Casual,
	Serious,
	Flirtatious,
}

func randomConversatioCategory() ConversationCategory {
	rand.Seed(time.Now().UnixNano())
	return allConversationCategories[rand.Intn(len(allConversationCategories))]
}
