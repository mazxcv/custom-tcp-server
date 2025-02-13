package first

import (
	"fmt"
	"time"
)

type Server struct {
	users  map[string]string
	userch chan string
	quitch chan struct{}
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
running:
	for {
		select {
		case user := <-s.userch:
			s.AddUser(user)
		case <-s.quitch:
			fmt.Println("server quitting")
			break running
		default:
		}
	}
}

func (s *Server) AddUser(user string) {
	s.users[user] = user
}

func serverplaschan() {

	server := NewServer()
	server.Start()

	go func() {
		time.Sleep(1 * time.Second)
		close(server.quitch)
		// server.quitch <- struct{}{}
	}()
	select {}
}
