package peer

import (
	"chat-app/helper"
	"net"
)

type Peer struct {
	Id   helper.UserId
	Conn net.Conn
}

func New(conn net.Conn) *Peer {
	return &Peer{
		Conn: conn,
	}
}

func (p *Peer) ReadLoop() {
	// keep reading from the peer
}
