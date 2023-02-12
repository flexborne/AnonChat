package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func badRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func (app *App) HandleConnections(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling new connection")
	// Upgrade initial GET request to a websocket
	ws, err := app.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to websocket:", err)
		badRequestResponse(w, err)
		return
	}

	client := NewClient(ws, ClientAttributes{}, ClientPreferences{})

	// Send the client its unique ID
	initialResponse := Message{
		"type": "id",
		"id":   client.Id,
	}

	ws.WriteJSON(initialResponse)

	go func() {
		defer func() {
			ws.Close()
			app.Clients.Delete(client.Id)
			app.RemoveFromQueueCh <- convertToMatchmakingParams(client)
			client.DisconnectFromChat()
			client = nil
		}()
		for {
			var msg Message
			err := ws.ReadJSON(&msg)

			if err != nil {
				fmt.Println("Error reading json:", err)
				if _, ok := err.(*websocket.CloseError); ok {
					return
				}
				break
			}

			if msg["type"] == "cmd" {
				app.handleServerCommand(w, msg, client)
			}

			if msg["type"] == "chat" {
				err := client.BroadcastMessage(msg)
				if err != nil {
					badRequestResponse(w, err)
				}

				msg["senderId"] = client.Id
			}
		}

	}()
}

type Message map[string]any

func (app *App) handleServerCommand(w http.ResponseWriter, msg Message, client *Client) {
	if msg["request"] == "findMatch" {
		err := client.DecodeJSON(msg)
		if err != nil {
			badRequestResponse(w, err)
		}

		fmt.Println("serverCommand")
		app.ClientsToMatch <- client
	}
}
