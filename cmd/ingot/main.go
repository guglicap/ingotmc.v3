package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/guglicap/ingotmc.v3/cmd/ingot/internal/server"
)

func main() {
	srv := server.New()
	go srv.Start()
	time.AfterFunc(15*time.Second, func() {
		srv.Close("timeout")
	})
	for {
		var cmd string
		_, err := fmt.Scanf("%s\n", &cmd)
		if err != nil {
			continue
		}
		switch cmd {
		default:
			color.Red("unknown command <%s>", cmd)
		case "close":
			srv.Close("user requested")
		case "start":
			go srv.Start()
		}
	}
}
