package server

import "net"

type newConn struct {
	net.Conn
}

func (n *newConn) id() eventID {
	return eventID("newConn")
}
