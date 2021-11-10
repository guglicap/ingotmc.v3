package main

import (
	client2 "github.com/guglicap/ingotmc.v3/client"
	kaki2 "github.com/guglicap/ingotmc.v3/cmd/ingot/internal/kaki"
	"github.com/guglicap/ingotmc.v3/cmd/ingot/internal/simulation"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	l    net.Listener
	quit chan os.Signal
	sim  *simulation.Simulation
}

func (s *server) run() {
	clients := s.acceptClients()
loop:
	for {
		select {
		case <-s.quit:
			break loop
		case c := <-clients:
			sock := client2.NewSocket(c)
			p := kaki2.New()
			cl := client2.NewClient(sock, p, kaki2.OfflineAuthenticator)
			go cl.Run()
			go s.sim.SpawnPlayerFor(cl)
		}
	}
	_ = s.l.Close()
}

func (s *server) acceptClients() chan net.Conn {
	s.l, _ = net.Listen("tcp", ":25565")
	clients := make(chan net.Conn)
	go func() {
		defer close(clients)
		for {
			c, err := s.l.Accept()
			if err != nil {
				return
			}
			clients <- c
		}
	}()
	return clients
}

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	server := server{quit: quit, sim: simulation.New()}
	server.run()
}
