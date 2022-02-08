package msgpackrpc

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type msgPackData struct {
	seq           uint64
	serviceMethod string
	content       []byte
	contentT      contentT
}

// 协议
// _total4 seq8 kenLen1 key[:] contentT_1 param/arg

// contentT 代表着数据类型，arg reply error. 当为error的时候 contentT=1，后面的error不需要进行编码/解码
type contentT byte

const contentTErr = 1

// write request or response according to the content,which is arg or reply.
func write(conn net.Conn, serviceMethod string, seq uint64, t contentT, contentB []byte) error {
	// content为错误的时候进行透传

	// todo校验contentB的长度
	total := 4 + 8 + 1 + len(serviceMethod) + 1 + len(contentB)
	buf := make([]byte, total)
	binary.BigEndian.PutUint32(buf[:4], uint32(total-4))
	binary.BigEndian.PutUint64(buf[4:12], seq)
	buf[12] = byte(len(serviceMethod))
	copy(buf[13:], serviceMethod)
	buf[13+len(serviceMethod)] = byte(t)
	copy(buf[13+1+len(serviceMethod):], contentB)
	n, err := conn.Write(buf)
	if err != nil {
		return err
	}
	if n != len(buf) {
		return fmt.Errorf("written len：%d is not equal buf len: %d", n, len(buf))
	}
	return nil
}

// read request or response to msgPackData.content
func read(conn net.Conn) (*msgPackData, error) {
	totalHeader := make([]byte, 4)
	_, err := io.ReadFull(conn, totalHeader)
	if err != nil {
		return nil, err
	}
	totalLen := binary.BigEndian.Uint32(totalHeader) // 去除头部的长度

	totalBuf := make([]byte, totalLen)
	_, err = io.ReadFull(conn, totalBuf)
	if err != nil {
		return nil, err
	}
	content := totalBuf[9+totalBuf[8]+1:]
	return &msgPackData{
		seq:           binary.BigEndian.Uint64(totalBuf[:8]),
		serviceMethod: string(totalBuf[9 : 9+totalBuf[8]]),
		content:       content,
		contentT:      contentT(totalBuf[9+totalBuf[8]]),
	}, nil
}

// 核心为使用是msgpack包的 marshal 和unmarshal方法
// 初测已经通过，尚未通过高级别的测试

// ithink@thinkpad-w520:~/go/src/vitogo.tpddns.cn/liushihao/gostd/pkg/rpc/msgpackrpc → dev$ go test -bench=.
// goos: linux
// goarch: amd64
// pkg: openblog/pkg/rpc/msgpackrpc
// cpu: Intel(R) Core(TM) i7-2760QM CPU @ 2.40GHz
// BenchmarkJsonMarshal-8            427910              2556 ns/op
// BenchmarkJsonUnmarshal-8          223660              6229 ns/op
// BenchmarkMsgPMarshal-8            817030              1471 ns/op
// BenchmarkMsgPUnMarshal-8          521344              2515 ns/op
// BenchmarkProtoMarshal-8           410091              2480 ns/op
// BenchmarkProtoUnmarshal-8         398979              2983 ns/op
// PASS
// ok      openblog/pkg/rpc/msgpackrpc     8.425s
