package nano

import (
	"github.com/lonng/nano/internal/runtime"
	"github.com/lonng/nano/session"
)

/**
 * 创建一个用于Rpc的Session oyzm
**/
func NewRpcSession(addr string) (*session.Session, error) {
	n := runtime.CurrentNode
	return n.NewRpcSession(addr)
}
