// Package mc is likely a bad idea.
// It's basically a catch-all package for data-types that don't have their own package yet,
// because they're needed but I haven't implemented everything yet.
// It should be temporary and kept as empty as possible. In theory.
package mc

import "encoding/binary"

type EID int32

func EIDFrom(v []byte) EID {
	l := 4
	if x := len(v); x < l {
		l = x
	}
	return EID(binary.BigEndian.Uint32(v[:l]))
}
