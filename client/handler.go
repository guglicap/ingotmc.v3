package client

import (
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/event"
)

// handle intercepts internal actions and sends the rest to the simulation
func (c *Client) handle(e action.Action) {
	switch ev := e.(type) {
	case action.NewConnection:
		handleLogin(c, ev)
	default:
		c.actions <- e
	}
}

// handleLogin handles authentication.
// TODO: should be a bigger deal, implement encryption, compression
func handleLogin(cl *Client, nc action.NewConnection) {
	userUUID, err := cl.auth.Authenticate(nc.Username)
	if err != nil {
		_ = cl.ProcessEvent(event.Disconnect{Reason: err})
		return
	}
	_ = cl.ProcessEvent(event.AuthSuccess{
		UUID:     userUUID,
		Username: nc.Username,
	})
	cl.actions <- action.NewPlayer{
		UUID:     userUUID,
		Username: nc.Username,
	}
}
