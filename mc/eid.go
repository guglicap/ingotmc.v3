// Package mc contains type defitions related to objects which permeate the whole minecraft domain.
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
