package main

import (
	"math/rand"
	"time"
)

// AgeCategory defines the different age categories for clients
type AgeCategory int

const (
	Under18 AgeCategory = 1 << iota
	Between18And25
	Between26And35
	Over35
)

func (c AgeCategory) HasAgeCategory(a AgeCategory) bool {
	return c&a == a
}

var allAgeCategories = []AgeCategory{
	Under18,
	Between18And25,
	Between26And35,
	Over35,
}

func randomAgeCategory() AgeCategory {
	rand.Seed(time.Now().UnixNano())
	return allAgeCategories[rand.Intn(len(allAgeCategories))]
}
