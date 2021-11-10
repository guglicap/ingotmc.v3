package proto

import (
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/event"
)

// Protocol describes a protocol implementation.
type Protocol interface {
	// ActionFor generates an Action which matches the given packet.
	// Note: can return (nil, nil) if the packet is a noop (it only needs to be handled internally by the protocol implementation)
	// Returns an error if no valid Actions are found or if packet decoding went wrong.
	ActionFor(packet []byte) (action.Action, error)

	// PacketFor accepts a sim event and returns a matching packet if found. If not, it returns an error instead.
	// It may also return an error if packet encoding went wrong.
	PacketFor(callback event.Event) ([]byte, error)
}
