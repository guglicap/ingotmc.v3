package server

import (
	"errors"
	"net"

	"github.com/cskr/pubsub"
	"github.com/fatih/color"
)

type eventID string

type event interface {
	id() eventID
}

type server struct {
	events *pubsub.PubSub[event]
	l net.Listener
}

func (s *server) listen(id eventID, f func (event)) {
	evs := s.events.Sub(string(id))
	go func() {
		for ev := range evs {
			f(ev)
		}
	}()
}

func New() *server {
	s := &server{
		events: pubsub.New[event](5),
	}
	s.listen("closing", func(e event) {
		color.Red("closing with reason: %s", e.(*closing).reason)
	})
	s.listen("newConn", func (e event)  {
			color.Cyan("newConn %v", e)
	})
	s.listen("starting", func (e event)  {
		color.Green("starting: %s", e.(starting).message)
	})
	return s
}

func (s server) emit(e event) {
	s.events.Pub(e, string(e.id()))
}

type closing struct {
	reason string
}

func (c closing) id() eventID {
	return "closing"
}

func (s server) closingWith(reason string) {
	s.emit(&closing{reason})
}

type starting struct{
	message string
}

func (s starting) id() eventID {
	return "starting"
}

func (s *server) Start() {
	s.emit(starting{"starting the server"})
	var err error
	s.l, err = net.Listen("tcp", ":51726")
	if err != nil {
		s.closingWith("error: " + err.Error())
		return
	}
	for {
		var cl net.Conn
		cl, err = s.l.Accept()
		if err != nil {
			switch {
			case errors.Is(err, net.ErrClosed):
				break
			default:
				s.closingWith("error: " + err.Error())
				s.l.Close()
			}
			break
		}
		s.emit(&newConn{cl})
	}
}

func (s *server) untilClose(f func ()) {
	for {
		select {
		case <-s.events.SubOnce("close"):
			return
		default:
			f()
		}
	}
}

func (s *server) Close(reason string) {
	if s.l == nil {
		return
	}
	s.closingWith(reason)
	s.l.Close()
}