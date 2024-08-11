package server

import (
	"chat-app/helper"
	"net"
	"sync"
)

type ServerOpts func(*Server)

func New(serverOpts ...ServerOpts) *Server {
	s := &Server{
		listenAddr: ":3000",

		mu:    sync.RWMutex{},
		peers: make(map[helper.UserId]*net.Conn),
	}
	for _, o := range serverOpts {
		o(s)
	}
	return s
}

func WithListenAddr(listenAddr string) ServerOpts {
	return func(s *Server) {
		s.listenAddr = listenAddr
	}
}
