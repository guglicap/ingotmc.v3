// Package client handles the specifics of connections with minecraft clients.
// It doesn't directly deal with packet encoding / decoding, which was left to implement via the Protocol interface
// but rather implements the logic for authentication, event dispatch to the simulation and event processing.
//
package client

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/proto"
	"log"
	"os"
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
	ctx   context.Context
	close context.CancelFunc

	socket *Socket //NOTE: should probably make it an interface
	proto  proto.Protocol
	auth   proto.Authenticator

	// these are kept here to allow handling at shutdown
	cbound chan<- []byte // packets to client
	sbound <-chan []byte // packets from client

	actions chan action.Action // simulation bound actions
}

// NewClient creates a new client.
// TODO: better defaults, functional parameters
func NewClient(socket *Socket, protocol proto.Protocol, authenticator proto.Authenticator) *Client {
	c := &Client{
		id: ID(uuid.New()),

		socket:  socket,
		proto:   protocol,
		auth:    authenticator,
		actions: make(chan action.Action),
	}
	c.ctx = context.Background()
	c.ctx = context.WithValue(c.ctx, "client_id", c.id)
	c.ctx, c.close = context.WithCancel(c.ctx)
	c.log = log.New(os.Stdout, fmt.Sprintf("client %s: ", c.id), log.LstdFlags|log.Lmsgprefix)
	return c
}

// Run starts the client.
// TODO: separate internal event processing and sim connection logic (channel setup rn), maybe rethink it
func (c *Client) Run() {
	c.sbound, c.cbound = c.socket.Start(c.ctx)
	c.actions = c.processActions()
}

func (c *Client) processActions() chan action.Action {
	simboundActions := make(chan action.Action)
	process := func() {
	loop:
		for {
			select {
			case <-c.ctx.Done():
				break loop
			case pkt, ok := <-c.sbound:
				if !ok {
					break loop
				}
				c.handlePacket(pkt)
			}
		}
		close(simboundActions)
		close(c.cbound)
		c.close()
		c.log.Println("goodbye")
	}
	go process()
	return simboundActions
}

// handlePacket asks protocol to decode a packets and forwards action handling to an appropriate function
func (c *Client) handlePacket(pkt []byte) {
	act, err := c.proto.ActionFor(pkt)
	if err != nil {
		c.log.Println(fmt.Errorf("error getting action: %e", err))
		return
	}
	if act == nil { // action was handled internally by protocol
		return
	}
	c.handle(act)
}

// ProcessEvent implements the simulation.Client interface.
// It asks the protocol to generate a packet for the event v and sends it to the client.
func (c *Client) ProcessEvent(v event.Event) error {
	pkt, err := c.proto.PacketFor(v)
	if err != nil {
		return err
	}
	c.cbound <- pkt
	return nil
}

// Actions implements the simulation.Client interface.
// Returns a chan of outbound actions, i.e. actions that the client shouldn't process internally.
func (c *Client) Actions() <-chan action.Action {
	return c.actions
}
