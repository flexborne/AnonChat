package main

import (
	"fmt"
	"sync"
	"time"
)

// Provides mechanism of "random" matchmaking: connects clients with the longest waiting time
type RandomMatchmakingManager struct {
	Clients map[SexCategory]map[ConversationCategory]map[AgeCategory][]ClientMatchmakingParams
	Mutex   sync.Mutex
	Id      IdType
}

func (q *RandomMatchmakingManager) push(x ClientMatchmakingParams) {
	fmt.Println(x.Attributes.Sex, x.Attributes.Category, x.Attributes.Age)
	q.Clients[x.Attributes.Sex][x.Attributes.Category][x.Attributes.Age] =
		append(q.Clients[x.Attributes.Sex][x.Attributes.Category][x.Attributes.Age], x)
}

func (q *RandomMatchmakingManager) delete(x ClientMatchmakingParams) {
	for i, client := range q.Clients[x.Attributes.Sex][x.Attributes.Category][x.Attributes.Age] {
		if client.Id == x.Id {
			q.Clients[x.Attributes.Sex][x.Attributes.Category][x.Attributes.Age] = append(q.Clients[x.Attributes.Sex][x.Attributes.Category][x.Attributes.Age][:i],
				q.Clients[x.Attributes.Sex][x.Attributes.Category][x.Attributes.Age][i+1:]...)
			break
		}
	}
}

// newQueue creates a new queue with specific age and conversation categories
func NewRandomMatchmakingManager() *RandomMatchmakingManager {
	var queue *RandomMatchmakingManager
	queue = &RandomMatchmakingManager{
		Clients: make(map[SexCategory]map[ConversationCategory]map[AgeCategory][]ClientMatchmakingParams),
		Mutex:   sync.Mutex{},
		Id:      generateUID(),
	}
	for _, sex := range allSexCategories {
		queue.Clients[sex] = make(map[ConversationCategory]map[AgeCategory][]ClientMatchmakingParams)
		for _, conv := range allConversationCategories {
			queue.Clients[sex][conv] = make(map[AgeCategory][]ClientMatchmakingParams)
			for _, age := range allAgeCategories {
				queue.Clients[sex][conv][age] = make([]ClientMatchmakingParams, 0)
			}
		}
	}

	return queue
}

func (q *RandomMatchmakingManager) run(clientMatchCh <-chan ClientMatchmakingParams,
	matchmakingResponseCh <-chan MatchmatckingResponse,
	matchmakingResultCh chan<- ClientMatchmakingResult,
	removeCh <-chan ClientMatchmakingParams) {
	for {
		select {
		case c := <-clientMatchCh:
			fmt.Println(c.Attributes.Sex, c.Attributes.Age, c.Attributes.Category,
				c.Preferences.Sex, c.Preferences.AgeCategories)
			q.Mutex.Lock()
			isAnyMatch := false
			var bestMatch ClientMatchmakingParams
			for _, ageCategory := range allAgeCategories {
				if c.Preferences.AgeCategories.HasAgeCategory(ageCategory) {
					for _, possibleMatch := range q.Clients[c.Preferences.Sex][c.Attributes.Category][ageCategory] {
						if match(c, possibleMatch) && c.Id != possibleMatch.Id &&
							(possibleMatch.WaitingStartTime.Before(bestMatch.WaitingStartTime) ||
								bestMatch.WaitingStartTime.IsZero()) {
							bestMatch = possibleMatch
							isAnyMatch = true
							break
						}
					}
				}
			}

			if isAnyMatch {
				matchmakingResultCh <- ClientMatchmakingResult{c.Id, bestMatch.Id}
				response := <-matchmakingResponseCh // we should explicitly wait for response in order to avoid deadlocks
				if response.HasValue() {
					q.delete(response.Value.(ClientMatchmakingParams))
				}
			} else {
				c.WaitingStartTime = time.Now()
				q.push(c)
			}
			q.Mutex.Unlock()
		case c := <-removeCh:
			q.Mutex.Lock()
			q.delete(c)
			q.Mutex.Unlock()
		}
	}
}

type ClientMatchmakingParams struct {
	Attributes       ClientAttributes
	Preferences      ClientPreferences
	Id               IdType
	WaitingStartTime time.Time
}

func convertToMatchmakingParams(client *Client) ClientMatchmakingParams {
	return ClientMatchmakingParams{
		Attributes:  client.Attributes,
		Preferences: client.Preferences,
		Id:          client.Id,
	}
}

func match(c1, c2 ClientMatchmakingParams) bool {
	if c1.Attributes.Sex != c2.Preferences.Sex ||
		c2.Attributes.Sex != c1.Preferences.Sex {
		return false
	}

	if c1.Attributes.Category != c2.Attributes.Category { // clients should have the same conv category
		return false
	}

	if !c1.Preferences.AgeCategories.HasAgeCategory(c2.Attributes.Age) ||
		!c2.Preferences.AgeCategories.HasAgeCategory(c1.Attributes.Age) {
		return false
	}

	return true
}

type ClientMatchmakingResult struct {
	FinderId    IdType
	BestMatchId IdType
}

type MatchmatckingResponse = Optional // has ClientMatchmakingParams to delete if connection is successful
