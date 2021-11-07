package client

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/guglicap/ingotmc.v3/client/callback"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/socket"
	"log"
	"os"
)

type ClientID uuid.UUID

func (c ClientID) String() string {
	return uuid.UUID(c).String()
}

// Client glues everything together.
// It receives packets from the socket, decodes them via the protocol and dispatches the event to the simulation.
// It receives callbacks from the protocol / simulation encodes them through the protocol and sends them to the socket.
type Client struct {
	ID    ClientID
	log   *log.Logger
	ctx   context.Context
	close context.CancelFunc

	socket *socket.Socket //NOTE: should probably make it an interface
	proto  Protocol
}

func NewClient(socket *socket.Socket, protocol Protocol) *Client {
	c := &Client{
		ID: ClientID(uuid.New()),

		socket: socket,
		proto:  protocol,
	}
	c.ctx = context.Background()
	c.ctx = context.WithValue(c.ctx, "client_id", c.ID)
	c.ctx, c.close = context.WithCancel(c.ctx)
	c.log = log.New(os.Stdout, fmt.Sprintf("client %s: ", c.ID), log.LstdFlags | log.Lmsgprefix)
	return c
}

func (c *Client) Run() {
	sbound, cbound := c.socket.Start(c.ctx)
	events := c.proto.Process(c.ctx, sbound)
loop:
	for ev := range events {
		switch t := ev.(type) {
		case proto.EventFatalError:
			dc := callback.Disconnect{t.Err.Error()}
			pkt, err := c.proto.PacketFor(dc)
			if err == nil {
				cbound <- pkt
			}
			break loop
		case proto.EventError:
			c.log.Println(t)
		default:
			fmt.Println("unknown event type (??)")
		}
	}
	c.close()
	c.log.Println("goodbye")
}