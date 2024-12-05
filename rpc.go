package nano

import (
	"errors"
	"sync"

	"github.com/ouyangzhongmin/nano/internal/runtime"
	"github.com/ouyangzhongmin/nano/session"
)

var (
	rpcSessions sync.Map
)

/**
 * 用于内部集群Rpc的调用 oyzm
 * 由于nano的rpc异步规则导致不能有返回数据
**/
func RPCWithAddr(route string, v interface{}, addr string) error {
	if addr == "" {
		return errors.New("addr is empty!")
	}
	ss, ok := rpcSessions.Load(addr)
	if !ok {
		var err error
		// 这里由于是主要用于rpc，不需要用到push到client的操作, 所以没用真实的gateAddr
		gateAddr := addr
		ss, err = runtime.CurrentNode.NewRpcSession(gateAddr)
		if err != nil {
			return err
		}
		rpcSessions.Store(addr, ss)
	}

	return ss.(*session.Session).RPCWithAddr(route, v, addr)
}

/**
 * 用于内部集群Rpc的调用 oyzm
 * 由于nano的rpc异步规则导致不能有返回数据
**/
func RPC(route string, v interface{}) error {
	addr := runtime.CurrentNode.ServiceAddr
	ss, ok := rpcSessions.Load(addr)
	if !ok {
		var err error
		// 这里由于是主要用于rpc，不需要用到push到client的操作, 所以没用真实的gateAddr
		gateAddr := addr
		ss, err = runtime.CurrentNode.NewRpcSession(gateAddr)
		if err != nil {
			return err
		}
		rpcSessions.Store(addr, ss)
	}

	return ss.(*session.Session).RPCWithAddr(route, v, "")
}
