package mc

import (
	"encoding/hex"
	"strings"
)

// UUID is a user unique id.
type UUID [16]byte

// String implements the stringer interface. Note this returns an hyphenated hex string.
func (u UUID) String() string {
	s := hex.EncodeToString(u[:])
	s = strings.Join([]string{s[:8], s[8:12], s[12:16], s[16:20], s[20:]}, "-")
	return s
}
