// Package myrpc 尝试自定义一种rpc协议
package msgpackrpc

import (
	"net"

	"github.com/vmihailenco/msgpack/v5"

	"github.com/vito-go/gostd/pkg/rpc"
)

type serverCodec struct {
	conn net.Conn // 用于读写数据，实际是一个网络连接
	arg  []byte   // param
}

func NewClientCodec(conn net.Conn) *clientCodec {
	return &clientCodec{conn: conn}
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) error {
	msg, err := read(c.conn)
	if err != nil {
		return err
	}
	r.Seq = msg.seq
	r.ServiceMethod = msg.serviceMethod
	c.arg = msg.content
	return nil
}

func (c *serverCodec) ReadRequestBody(arg interface{}) error {
	return msgpack.Unmarshal(c.arg, arg)

}
func (c *serverCodec) WriteResponse(r *rpc.Response, reply interface{}) error {
	if r.Error != "" {
		return write(c.conn, r.ServiceMethod, r.Seq, 1, []byte(r.Error))
	}
	contentB, err := msgpack.Marshal(reply)
	if err != nil {
		return err
	}
	return write(c.conn, r.ServiceMethod, r.Seq, 0, contentB)
}

func (c *serverCodec) Close() error {
	return c.conn.Close()
}
func (c *serverCodec) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}
