package client

import (
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/proto"
)

// handle intercepts internal actions and sends the rest to the simulation
func (c *Client) handle(e action.Action) {
	switch ev := e.(type) {
	case proto.EventFatalError:
		_ = c.ProcessEvent(event.Disconnect{Reason: ev.Err})
	case action.NewConnection:
		handleLogin(c, ev)
	default:
		c.actions <- e
	}
}

// handleLogin handles authentication.
// TODO: should be a bigger deal, implement encryption, compression
func handleLogin(c *Client, nc action.NewConnection) {
	userUUID, err := c.auth.Authenticate(nc.Username)
	if err != nil {
		_ = c.ProcessEvent(event.Disconnect{Reason: err})
		return
	}
	_ = c.ProcessEvent(event.AuthSuccess{
		UUID:     userUUID,
		Username: nc.Username,
	})
	c.actions <- action.NewPlayer{
		UUID:     userUUID,
		Username: nc.Username,
	}
}
