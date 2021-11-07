package client

import (
	"context"
	"github.com/guglicap/ingotmc.v3/client/callback"
	"github.com/guglicap/ingotmc.v3/event"
)

type Protocol interface {
	// Process reads incoming packets on the packets chan, decodes them and sends the resulting Events back on the appropriate channel.
	// It should stop processing once the context signals to do so or if the packets chan gets closed.
	Process(ctx context.Context, packets <-chan []byte) chan event.Event

	PacketFor(callback callback.Callback) ([]byte, error)
}
