package msgpackrpc

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/vmihailenco/msgpack/v5"

	"github.com/vito-go/gostd/pkg/rpc/msgpackrpc/rpc_protobuf/pb"
)

//go:generate go test -bench=. -benchmem

func BenchmarkJsonMarshal(b *testing.B) {
	a := pb.SquareResponse1{
		Num: 111,
		Ans: 222,
		Aaa: []string{"111", "bbb", "ccc"},
		Mmm: map[string]string{"mmm": "aaa", "nnn": "ddd"},
	}
	for i := 0; i < b.N; i++ {
		bb, err := json.Marshal(a)
		if err != nil {
			panic(err)
		}
		_ = bb
	}
}
func BenchmarkJsonUnmarshal(b *testing.B) {
	a := pb.SquareResponse1{
		Num: 111,
		Ans: 222,
		Aaa: []string{"111", "bbb", "ccc"},
		Mmm: map[string]string{"mmm": "aaa", "nnn": "ddd"},
	}
	bb, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	_ = bb
	for i := 0; i < b.N; i++ {
		var m pb.SquareResponse
		err = json.Unmarshal(bb, &m)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMsgPMarshal(b *testing.B) {
	a := pb.SquareResponse1{
		Num: 111,
		Ans: 222,
		Aaa: []string{"111", "bbb", "ccc"},
		Mmm: map[string]string{"mmm": "aaa", "nnn": "ddd"},
	}
	for i := 0; i < b.N; i++ {
		bb, err := msgpack.Marshal(&a)
		if err != nil {
			panic(err)
		}
		_ = bb
	}
}
func BenchmarkMsgPUnMarshal(b *testing.B) {
	a := pb.SquareResponse1{
		Num: 111,
		Ans: 222,
		Aaa: []string{"111", "bbb", "ccc"},
		Mmm: map[string]string{"mmm": "aaa", "nnn": "ddd"},
	}
	bb, err := msgpack.Marshal(&a)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		var m pb.SquareResponse1
		err = msgpack.Unmarshal(bb, &m)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkProtoMarshal(b *testing.B) {
	a := pb.SquareResponse{
		Num: 111,
		Ans: 222,
		Aaa: []string{"111", "bbb", "ccc"},
		Mmm: map[string]string{"mmm": "aaa", "nnn": "ddd"},
	}
	for i := 0; i < b.N; i++ {
		bb, err := proto.Marshal(&a)
		if err != nil {
			panic(err)
		}
		_ = bb
	}
}
func BenchmarkProtoUnmarshal(b *testing.B) {
	a := pb.SquareResponse{
		Num: 111,
		Ans: 222,
		Aaa: []string{"111", "bbb", "ccc"},
		Mmm: map[string]string{"mmm": "aaa", "nnn": "ddd"},
	}
	bb, err := proto.Marshal(&a)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		var m pb.SquareResponse
		err = proto.Unmarshal(bb, &m)
		if err != nil {
			panic(err)
		}
	}
}
