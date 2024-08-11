package main

import (
	"chat-app/pkg/server"

	"golang.org/x/exp/slog"
)

func main() {
	s := server.New(server.WithListenAddr(":3000"))
	s.Start()

	if r := recover(); r != nil {
		slog.Error("something went wrong %v", r)
	}
}
