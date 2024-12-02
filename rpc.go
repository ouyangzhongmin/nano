package nano

import (
	"github.com/ouyangzhongmin/nano/internal/runtime"
	"github.com/ouyangzhongmin/nano/session"
)

/**
 * 创建一个用于Rpc的Session oyzm
**/
func NewRpcSession(gateAddr string) (*session.Session, error) {
	n := runtime.CurrentNode
	return n.NewRpcSession(gateAddr)
}
