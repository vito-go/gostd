package msgpackrpc

import (
	"net"

	"github.com/vmihailenco/msgpack/v5"

	"github.com/vito-go/gostd/pkg/rpc"
)

type clientCodec struct {
	conn  net.Conn
	reply []byte
}

func NewServerCodec(conn net.Conn) *serverCodec {
	return &serverCodec{conn: conn}
}

func (c *clientCodec) WriteRequest(r *rpc.Request, arg interface{}) error {
	contentB, err := msgpack.Marshal(arg)
	if err != nil {
		return err
	}
	return write(c.conn, r.ServiceMethod, r.Seq, 0, contentB)
}

func (c *clientCodec) ReadResponseHeader(r *rpc.Response) error {
	data, err := read(c.conn)
	if err != nil {
		return err
	}
	r.Seq = data.seq
	r.ServiceMethod = data.serviceMethod
	if data.contentT == contentTErr {
		r.Error = string(data.content)
		// r.Error 不为空的时候 代表server返回的error，不会走ReadResponseBody
	}
	c.reply = data.content
	return nil
}

func (c *clientCodec) ReadResponseBody(reply interface{}) error {
	// reply is nil when response has error
	if reply == nil {
		return nil
	}
	return msgpack.Unmarshal(c.reply, reply)
}

func (c *clientCodec) Close() error {
	return c.conn.Close()
}
