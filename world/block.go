package world

// Block needs to be implemented.
// TODO: implement Block.
// Note: this probably belongs in world, with actual blocks being implemented in a protocol-specific package
type Block interface {
	ID() int
	Solid() bool
	Properties() struct{} // TODO: implement Block properties
}

