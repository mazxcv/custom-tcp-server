package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitchan   chan struct{}
	msgchan    chan Message
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitchan:   make(chan struct{}),
		msgchan:    make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}

	defer ln.Close()
	s.ln = ln
	go s.acceptLoop()
	<-s.quitchan
	close(s.msgchan)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		fmt.Println("accept:", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		s.msgchan <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}

		conn.Write([]byte(s.listenAddr + " received\n"))
	}
}

func main() {
	server := NewServer("127.0.0.1:3000")
	go func() {
		for msg := range server.msgchan {
			fmt.Printf("received message from connection (%s): %s\n", msg.from, string(msg.payload))
		}
	}()
	log.Fatal(server.Start())
}
