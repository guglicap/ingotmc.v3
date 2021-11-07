package main

import (
	"github.com/guglicap/ingotmc.v3/client"
	"github.com/guglicap/ingotmc.v3/kaki"
	"github.com/guglicap/ingotmc.v3/socket"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s, _ := net.Listen("tcp", ":25565")
	clients := make(chan net.Conn)
	go func() {
		defer close(clients)
		for {
			c, err := s.Accept()
			if err != nil {
				return
			}
			clients <- c
		}
	}()
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
loop:
	for {
		select {
		case <-stop:
			break loop
		case c := <-clients:
			s := socket.NewSocket(c)
			p := kaki.New()
			cl := client.NewClient(s, p)
			go cl.Run()
		}
	}
	_ = s.Close()
}
