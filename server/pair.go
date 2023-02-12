package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

// Pair exists only when both clients are connected
type Pair struct {
	Clients      []*websocket.Conn
	BroadcastCh  chan Message
	DisconnectCh chan IdType
	Lock         sync.Mutex
}

func newPair(c1, c2 *Client) *Pair {
	pair := &Pair{
		Clients:      make([]*websocket.Conn, 0),
		BroadcastCh:  make(chan Message),
		DisconnectCh: make(chan IdType),
	}

	pair.Clients = append(pair.Clients, c1.Conn)
	pair.Clients = append(pair.Clients, c2.Conn)

	c1.PairConnection = NewOptional(PairConnection{
		pair.BroadcastCh, pair.DisconnectCh,
	})
	c2.PairConnection = NewOptional(PairConnection{
		pair.BroadcastCh, pair.DisconnectCh,
	})

	return pair
}

func (g *Pair) run() {
	for {
		select {
		case message, ok := <-g.BroadcastCh:
			if !ok {
				return
			}
			for _, c := range g.Clients {
				text, ok := message["text"].(string)
				if !ok {
					continue
				}
				fmt.Println(text)
				b, _ := json.Marshal(message)
				fmt.Println("marshalled", string(b))

				c.WriteJSON(message)
				fmt.Println("message", message)
				fmt.Println("message", message)

			}
		case id := <-g.DisconnectCh:
			for _, c := range g.Clients {
				c.WriteJSON(Message{"type": "event", "event": "disconnect", "id": id})
			}
			return
		}
	}
}
