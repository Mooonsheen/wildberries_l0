package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"wildberries/internal/server"
)

func HandleInterrupt(s *server.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\r")
		s.Down()
		os.Exit(0)
	}()
}

func main() {
	server, err := server.NewServer("config.json")
	if err != nil {
		panic(err)
	}
	HandleInterrupt(server)
	if err := server.Up(); err != nil {
		panic(err)
	}
}
