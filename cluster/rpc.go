package cluster

import (
	"github.com/lonng/nano/service"
	"github.com/lonng/nano/session"
)

func (n *Node) NewRpcSession(addr string) (*session.Session, error) {
	sid := service.Connections.SessionID()
	return n.findOrCreateSession(sid, addr)
}
