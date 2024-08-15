package server

import (
	"chat-app/helper"
	"chat-app/pkg/peer"
	"fmt"
	"sync"
)

var serverInstance *Server = nil

type ServerOpts func(*Server)

func New(serverOpts ...ServerOpts) *Server {
	s := &Server{
		listenAddr: ":3000",

		mu:    sync.RWMutex{},
		peers: make(map[helper.UserId]peer.Conn),
	}
	// save the server instance
	serverInstance = s
	for _, o := range serverOpts {
		o(s)
	}
	return s
}

func GetServer() (*Server, error) {
	if serverInstance != nil {
		return serverInstance, nil
	}
	return nil, fmt.Errorf("no server exists")
}

func WithListenAddr(listenAddr string) ServerOpts {
	return func(s *Server) {
		s.listenAddr = listenAddr
	}
}
