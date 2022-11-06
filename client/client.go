// Package client is responsible for connection management.
// Specifically, it reads incoming packets from a Socket and decodes them through Protocol.
package client

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

// ID is an identifier associate with a unique client.
type ID uuid.UUID

func (c ID) String() string {
	return uuid.UUID(c).String()
}

// Client handles the logic of client authentication, player creation, event dispatch and event processing.
// Clients have two separate Action channels: the first one, which it gets directly from protocol is checked
// to see if actions are generated that should be processed internally, i.e authentication / encryption or anything
// else connection related. Otherwise, the actions are forwarded to the simulation via another channel, which the simulation
// can access via Actions()
type Client struct {
	id    ID
	log   *log.Logger
}

// NewClient creates a new client.
// TODO: better defaults, functional parameters
func NewClient() *Client {
	c := &Client{
		id: ID(uuid.New()),
	}
	c.log = log.New(os.Stdout, fmt.Sprintf("client %s: ", c.id), log.LstdFlags|log.Lmsgprefix)
	return c
}