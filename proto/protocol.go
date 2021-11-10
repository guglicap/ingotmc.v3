package proto

import (
	"context"
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/event"
)

// Protocol describes a protocol implementation.
type Protocol interface {
	// Process reads incoming packets on the packets chan, decodes them and sends the resulting Actions back on the appropriate channel.
	// It should stop processing once the context signals to do so or if the packets chan gets closed.
	// TODO: should this be asynchronous?
	Process(ctx context.Context, packets <-chan []byte) chan action.Action

	// PacketFor accepts a sim event and returns a matching packet if found. It returns an error instead.
	// It may also return an error if packet encoding went wrong.
	PacketFor(callback event.Event) ([]byte, error)
}
