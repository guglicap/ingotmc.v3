package world

// Gamemode represents a game mode.
// TODO: implement hardcore flag.
// TODO: implement Stringer.
type Gamemode uint8

const (
	Survival Gamemode = iota
	Creative
	Adventure
	Spectator
)
