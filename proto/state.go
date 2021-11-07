package proto

type State int

func (s State) String() string {
	switch s {
	case Handshaking:
		return "Handshaking"
	case Login:
		return "Login"
	case Play:
		return "Play"
	default:
		return "unknown"
	}
}

const (
	Handshaking State = iota
	Status
	Login
	Play
)
