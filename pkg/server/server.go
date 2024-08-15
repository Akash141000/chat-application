package server

import (
	"chat-app/helper"
	"chat-app/pkg/peer"
	"chat-app/router"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/exp/slog"
)

type Server struct {
	listenAddr string

	addPeerCh chan peer.Peer
	mu        sync.RWMutex
	peers     map[helper.UserId]peer.Conn
}

func (s *Server) Start() error {
	slog.Info("server", "starting the server", s.listenAddr)

	//register the routes
	mux := http.NewServeMux()
	router.New(mux)

	//start http server
	if err := http.ListenAndServe(s.listenAddr, mux); err != nil {
		slog.Error("server", "error starting the server", err)
		return err
	}

	//start accepting new peers
	if err := s.AcceptPeers(); err != nil {
		return err
	}

	return nil
}

// verify this
func (s *Server) AcceptPeers() error {
	for p := range s.addPeerCh {
		s.AddNewPeer(p)
	}
	return nil
}

// add new peer to server
func (s *Server) AddNewPeer(peer peer.Peer) error {
	id := peer.GetId()
	conn := peer.GetConn()
	slog.Info("server", "add new peer", id)
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.peers[id]; ok {
		return fmt.Errorf("peer already exists, %v", id)
	}
	s.peers[id] = conn

	return nil
}
