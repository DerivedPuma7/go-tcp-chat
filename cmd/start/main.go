package main

import (
	"log"
	"net"

	broadcast "github.com/derivedpuma7/go-tcp-chat/broadcaster"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcast.Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go broadcast.HandleConn(conn)
	}
}
