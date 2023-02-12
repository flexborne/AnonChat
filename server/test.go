package main

import (
	"fmt"
	"sync"
)

func addRandomClientsToMatchChannel(app *App) {
	mu := sync.Mutex{}
	for i := 0; i < 5000; i++ {
		mu.Lock()
		attributes := ClientAttributes{
			Sex:      randomSexCategory(),
			Category: randomConversatioCategory(),
			Age:      randomAgeCategory(),
		}
		preferences := ClientPreferences{
			Sex:           randomSexCategory(),
			AgeCategories: randomAgeCategory(),
		}
		client := NewClient(nil, attributes, preferences)
		app.ClientsToMatch <- client

		fmt.Println("last client added ", i)
		mu.Unlock()
	}
}
