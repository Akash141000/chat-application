package peer

import (
	"chat-app/helper"
	"fmt"

	"github.com/gorilla/websocket"
)

type Peer interface {
	GetId() helper.UserId
	GetConn() Conn
}

type Conn struct {
	*websocket.Conn
}

type WSPeer struct {
	Id   helper.UserId
	Conn Conn
}

func NewWS(conn Conn) *WSPeer {
	return &WSPeer{
		Conn: conn,
	}
}

func (p *WSPeer) ReadLoop() {
	// keep reading from the peer
	fmt.Println("p", p)
}

func (p *WSPeer) GetId() helper.UserId {
	return p.Id
}

func (p *WSPeer) GetConn() Conn {
	return p.Conn
}
