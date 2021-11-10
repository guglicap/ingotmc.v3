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
