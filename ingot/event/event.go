package event

import (
	"strings"

	"github.com/cskr/pubsub"
)

type ID string

type Subscription[E Event] chan E 

type Event interface {
	EventID() ID
}

type hub[E Event] struct {
	*pubsub.PubSub[E]

	prefix string
}

func BuildHub[E Event]() hub[E] {
	return hub[E]{
		PubSub: pubsub.New[E](5),
		prefix: "",
	}
}

func buildTopic(prefix string, id string) string {
	return strings.Join([]string{ // topic
		prefix,
		id,
	}, "/")
}

func (h *hub[E]) Pub(ev E) {
	h.PubSub.Pub(
		ev, // msg
		buildTopic(h.prefix, string(ev.EventID())), // topic
	)
}

func (h *hub[E]) Sub(id ID) Subscription[E] {
	evs := h.PubSub.Sub(buildTopic(h.prefix, string(id)))
	return Subscription[E](evs)
}

func (h *hub[E]) UnSub(s Subscription[E]) {
	// check docs for pubsub.PubSub.Unsub()
	go h.PubSub.Unsub(s)
}

