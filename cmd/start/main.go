package main

import (
	"fmt"
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
		fmt.Println("aguardando client")
		conn, err := listener.Accept()
		fmt.Println("client conectado")
		if err != nil {
			log.Print(err)
			continue
		}
		go broadcast.HandleConn(conn)
	}
}
