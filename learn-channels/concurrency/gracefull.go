package concurrency

import (
	"fmt"
	"time"
)

type Message struct {
	From     string
	Payloads []byte
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {
free:
	for {
		select {
		case msg := <-s.msgch:
			fmt.Printf("message received from: %s, payload: %s \n", msg.From, msg.Payloads)
		case <-s.quitch:
			fmt.Println("server is doing a gracefull shutdown")
			// logic for the graceful shutdown

			break free
		default:
		}
	}

	fmt.Println("server is shutdown")
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:     "YoeByDem",
		Payloads: []byte(payload),
	}
	msgch <- msg
}

func gracefullShutdown(s *Server) {
	close(s.quitch)
}

// message communication
func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.StartAndListen()

	go func() {
		time.Sleep(1 * time.Second)
		sendMessageToServer(s.msgch, "Hello, Server.")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		gracefullShutdown(s)
	}()

	select {}
}
