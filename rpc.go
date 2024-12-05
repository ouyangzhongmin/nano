package nano

import (
	"sync"

	"github.com/ouyangzhongmin/nano/internal/runtime"
	"github.com/ouyangzhongmin/nano/session"
)

// func NewRpcSession(gateAddr string) (*session.Session, error) {
// 	n := runtime.CurrentNode
// 	return n.NewRpcSession(gateAddr)
// }

var (
	rpcSessions sync.Map
)

/**
 * 用于内部集群Rpc的调用 oyzm
 * 由于nano的rpc异步规则导致不能有返回数据
**/
func RPCWithAddr(route string, v interface{}, addr string) error {
	ss, ok := rpcSessions.Load(addr)
	if !ok {
		var err error
		// 这里由于是主要用于rpc，不需要用到push到client的操作, 所以没用真实的gateAddr
		ss, err = runtime.CurrentNode.NewRpcSession(addr)
		if err != nil {
			return err
		}
		rpcSessions.Store(addr, ss)
	}

	return ss.(*session.Session).RPCWithAddr(route, v, addr)
}
