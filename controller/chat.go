package controller

import (
	"chat-app/pkg/peer"
	// "chat-app/pkg/server"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Chat(context context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "chat started")

	// s, err := server.GetServer()
	// if err != nil {
	// 	return err
	// }

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("something went wrong while upgrading the connection", err)
	}

	conn := peer.Conn{
		Conn: wsConn,
	}
	peer.NewWS(conn)

	// s.AddNewPeer(peer)

	return nil
}
