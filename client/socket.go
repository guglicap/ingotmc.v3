package client

import (
	"context"
	"errors"
	"github.com/guglicap/ingotmc.v3/proto/decode"
	"github.com/guglicap/ingotmc.v3/proto/encode"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// the bytes array here are the raw packet data ( + packet id )
type sendFunc func([]byte, io.Writer) error
type recvFunc func(io.Reader) ([]byte, error)

// Socket handles low-level communication with a client.
// Its job is basically only sending and receiving packets, without actually encoding / decoding any of them
// TODO: compression
type Socket struct {
	conn net.Conn
	log  *log.Logger

	send         sendFunc
	clientbound  chan []byte
	writeTimeout time.Duration

	receive     recvFunc
	serverbound chan []byte
	isClosed    bool
}

const defaultWriteTimeout = 2 * time.Second

func NewSocket(conn net.Conn) *Socket {
	s := &Socket{
		conn:         conn,
		isClosed:     false,
		log:          log.New(os.Stdout, "socket: ", log.LstdFlags|log.Lmsgprefix),
		send:         sendPacket,
		writeTimeout: defaultWriteTimeout,
		serverbound:  make(chan []byte),
		clientbound:  make(chan []byte),

		receive: readPacket,
	}
	return s
}

func (s *Socket) Start(socketContext context.Context) (serverbound <-chan []byte, clientbound chan<- []byte) {
	go s.readWorker(socketContext)
	go s.writeWorker(socketContext)
	return s.serverbound, s.clientbound
}

func (s *Socket) readWorker(socketContext context.Context) {
loop:
	for {
		select {
		case <-socketContext.Done():
			s.log.Println("rWorker: stopping from context")
			break loop
		default:
			pkt, err := s.receive(s.conn)
			if err != nil {
				s.log.Println("rWorker: stopping from conn err")
				break loop
			}
			s.serverbound <- pkt
		}
	}
	s.close()
}

func (s *Socket) writeWorker(socketContext context.Context) {
loop:
	for {
		select {
		case <-socketContext.Done():
			s.log.Println("wWorker: stopping from context")
			break loop
		case pkt, ok := <-s.clientbound:
			if !ok {
				break loop
			}
			s.conn.SetWriteDeadline(time.Now().Add(s.writeTimeout)) // NOTE: handle error
			err := s.send(pkt, s.conn)
			if err != nil {
				s.log.Println("wWorker: stopping from conn err")
				break loop
			}
		}
	}
	s.close()
}

func (s *Socket) close() {
	if s.isClosed {
		return
	}
	s.isClosed = true
	close(s.serverbound)
	s.conn.Close()
	s.log.Println("goodbye")
}

// TODO: nicer errors, timeouts...
func readPacket(r io.Reader) (pkt []byte, err error) {
	length, err := decode.VarInt(r)
	if err != nil {
		return
	}
	pkt = make([]byte, length)
	n, err := r.Read(pkt)
	if err != nil {
		return
	}
	if n != int(length) {
		return nil, errors.New("socket: read less bytes")
	}
	return
}

func sendPacket(pkt []byte, w io.Writer) (err error) {
	l := len(pkt)
	err = encode.VarInt(int32(l), w)
	if err != nil {
		return
	}
	n, err := w.Write(pkt)
	if err != nil {
		return
	}
	if n != l {
		return errors.New("socket: wrote less bytes")
	}
	return
}
