package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type App struct {
	Upgrader                 websocket.Upgrader
	RandomMatchmakingCh      chan ClientMatchmakingParams
	MatchmakingResponseCh    chan MatchmatckingResponse
	RemoveFromQueueCh        chan ClientMatchmakingParams
	MatchmakingResultCh      chan ClientMatchmakingResult
	RandomMatchmakingManager *RandomMatchmakingManager
	Clients                  sync.Map
	MatchmakingChMutex       sync.Mutex
	Mutex                    sync.Mutex
	ClientsToMatch           chan *Client
}

func NewApp() *App {
	var app = &App{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		RandomMatchmakingCh:      make(chan ClientMatchmakingParams),
		MatchmakingResultCh:      make(chan ClientMatchmakingResult),
		MatchmakingResponseCh:    make(chan MatchmatckingResponse),
		RemoveFromQueueCh:        make(chan ClientMatchmakingParams),
		RandomMatchmakingManager: NewRandomMatchmakingManager(),
		Clients:                  sync.Map{},
		MatchmakingChMutex:       sync.Mutex{},
		Mutex:                    sync.Mutex{},
		ClientsToMatch:           make(chan *Client), // incoming client channels that want to be matched
	}
	return app
}
