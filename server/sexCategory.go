package main

import (
	"math/rand"
	"time"
)

// SexCategory defines the different sex categories for clients
type SexCategory int

const (
	Male SexCategory = iota
	Female
	Undefined
)

var allSexCategories = []SexCategory{
	Male,
	Female,
	Undefined,
}

func randomSexCategory() SexCategory {
	rand.Seed(time.Now().UnixNano())
	return allSexCategories[rand.Intn(len(allSexCategories))]
}
