package main

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type ClientAttributes struct {
	Sex      SexCategory          `json:"sex"`
	Category ConversationCategory `json:"category"`
	Age      AgeCategory          `json:"age"`
}

func (ca *ClientAttributes) decodeJSON(rawMap Message) error {
	if v, ok := rawMap["sex"].(float64); ok {
		ca.Sex = SexCategory(int(v))
	} else {
		return errors.New("field 'category' must be a float64")
	}

	if v, ok := rawMap["conversationCategory"].(float64); ok {
		ca.Category = ConversationCategory(int(v))
	} else {
		return errors.New("field 'category' must be a float64")
	}

	if v, ok := rawMap["ageCategory"].(float64); ok {
		ca.Age = AgeCategory(int(v))
	} else {
		return errors.New("field 'age' must be a float64")
	}

	return nil
}

type ClientPreferences struct {
	Sex           SexCategory
	AgeCategories AgeCategory
}

type PairConnection struct {
	MessagesCh   chan<- Message
	DisconnectCh chan<- IdType
}

func (cp *ClientPreferences) decodeJSON(rawMap Message) error {
	if v, ok := rawMap["sexPreferences"].(float64); ok {
		cp.Sex = SexCategory(int(v))
	} else {
		return errors.New("field 'age' must be a float64")
	}

	if v, ok := rawMap["agePreferences"].(float64); ok {
		cp.AgeCategories = AgeCategory(int(v))
	} else {
		return errors.New("field 'age' must be a float64")
	}

	return nil
}

type Client struct {
	Conn             *websocket.Conn
	Attributes       ClientAttributes
	Preferences      ClientPreferences
	Id               IdType
	PairConnection   Optional
	WaitingStartTime time.Time
	Mutex            sync.Mutex
}

func (client *Client) DecodeJSON(rawMap Message) error {
	client.Attributes.decodeJSON(rawMap)
	client.Preferences.decodeJSON(rawMap)

	return nil
}

func (client *Client) BroadcastMessage(msg Message) error {
	if !client.PairConnection.HasValue() {
		return errors.New("there is no connection")
	}

	client.PairConnection.Value.(PairConnection).MessagesCh <- msg
	return nil
}

func (client *Client) DisconnectFromChat() {
	if client.PairConnection.HasValue() {
		client.PairConnection.Value.(PairConnection).DisconnectCh <- client.Id
	}
}

func NewClient(conn *websocket.Conn, attributes ClientAttributes, preferences ClientPreferences) *Client {
	return &Client{
		Conn:           conn,
		Attributes:     attributes,
		Preferences:    preferences,
		PairConnection: Optional{Valid: false},
		Mutex:          sync.Mutex{},
		Id:             generateUID(),
	}
}
