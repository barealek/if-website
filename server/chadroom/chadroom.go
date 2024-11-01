package chadroom

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Chadroom struct {
	Client    *websocket.Conn
	Responder *websocket.Conn
	Messages  []Message
	State     State
	ctx       context.Context
	cancel    context.CancelFunc
}

func NewChadroom() *Chadroom {
	ctx, cancel := context.WithCancel(context.Background())
	return &Chadroom{
		Messages: []Message{},
		State:    StateNotStarted,
		ctx:      ctx,
		cancel:   cancel,
	}
}

type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

func (c *Chadroom) SetState(state State) {
	fmt.Printf("Setting state to %s\n", state)
	c.State = state
	c.Client.WriteJSON(map[string]string{
		"system": "state",
		"state":  string(state),
	})
	c.Responder.WriteJSON(map[string]string{
		"system": "state",
		"state":  string(state),
	})

}

func (c *Chadroom) StartChatting() error {
	if c.Client == nil || c.Responder == nil {
		return errors.New("client or responder is nil")
	}

	c.Client.WriteJSON(map[string]string{
		"system": "start",
	})
	c.SetState(StateWaitingUser)

	go c.monitorConnections()

	return nil
}

func (c *Chadroom) monitorConnections() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			if c.Client == nil || c.Responder == nil {
				c.cancel()
				log.Println("Client or Responder connection lost")
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func (c *Chadroom) ClientSendMessage(cnt string) {
	var message = Message{
		Author:  "user",
		Content: cnt,
	}
	c.Messages = append(c.Messages, message)
	c.UpdateMessages()
}

func (c *Chadroom) ResponderSendMessage(cnt string) {
	var message = Message{
		Author:  "chadgpt",
		Content: cnt,
	}
	c.Messages = append(c.Messages, message)
	c.UpdateMessages()
}

func (c *Chadroom) UpdateMessages() {
	c.Client.WriteJSON(map[string]interface{}{
		"system":   "update",
		"messages": c.Messages,
	})

	c.Responder.WriteJSON(map[string]interface{}{
		"system":   "update",
		"messages": c.Messages,
	})
}

func (c *Chadroom) ClientSendHint(cnt string) {
	c.Client.WriteJSON(map[string]string{
		"system": "hint",
		"hint":   cnt,
	})
}
