package broadcaster

import (
	"bufio"
	"fmt"
	"net"
)

type client chan<- string
type message struct {
	message string
	cli chan string
}

var (
	entering = make(chan chan<- string)
	leaving = make(chan client)
	messages = make(chan message)
)

func Broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
			case msg := <-messages:
				for cli := range clients {
					if msg.cli != cli {
						cli <- msg.message
					}
				}
			case cli := <-entering:
				clients[cli] = true
			case cli := <-leaving:
				delete(clients, cli)
				close(cli)
		}
	}
}

func HandleConn(conn net.Conn) {
	ch := make(chan string) //mensagens de saída do cliente
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- message{ message: who + " has arrived", cli: ch } 
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- message{ message: who + ": " + input.Text(), cli: ch } 
	}

	leaving <- ch
	messages <- message{ message: who + " has left", cli: ch } 
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
