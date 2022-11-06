package mc

import (
	"crypto/sha256"
	"encoding/binary"
)

type Seed string

func (s Seed) GetHash() uint64 {
	hash := sha256.Sum256([]byte(s))
	return binary.BigEndian.Uint64(hash[:8])
}
