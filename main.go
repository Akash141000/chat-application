package main

import (
	"chat-app/pkg/server"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("something went wrong %v", r)
		}
	}()

	s := server.New(server.WithListenAddr(":3000"))
	if err := s.Start(); err != nil {
		slog.Error("server failed to start", err)
		os.Exit(1)
	}

}
