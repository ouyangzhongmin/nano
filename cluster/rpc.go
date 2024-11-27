package cluster

import (
	"github.com/ouyangzhongmin/nano/service"
	"github.com/ouyangzhongmin/nano/session"
)

func (n *Node) NewRpcSession(addr string) (*session.Session, error) {
	sid := service.Connections.SessionID()
	return n.findOrCreateSession(sid, addr)
}
