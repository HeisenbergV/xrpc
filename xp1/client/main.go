package main

import (
	"fmt"
	"log"
	"net"
	"xp1/codec"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8889")
	defer func() { _ = conn.Close() }()

	cc := codec.NewGobCodec(conn)
	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("xrpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
