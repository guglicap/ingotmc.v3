package proto

import (
	"fmt"
)

var ErrorUnsupportedCallback = fmt.Errorf("event not supported")

func ErrorUnsupportedState(s State) error {
	return fmt.Errorf("unsupported state: %s", s)
}

func ErrorUnsupportedPacket(s State, id int32) error {
	return fmt.Errorf("unsupported packet for state %s: %x", s, id)
}

func ErrorMismatchedProtocol(have, want int32) error {
	return fmt.Errorf("mismatched protocol versions: have %d want %d", have, want)
}

// EventFatalError signals that something happened in the protocol and future requests cannot be served.
type EventFatalError struct {
	Err error
}

//EventError signals that something went wrong in the protocol, but future requests may still be served
type EventError struct {
	Err error
}

func (ee EventError) Error() string {
	return "protocol error: " + ee.Err.Error()
}
