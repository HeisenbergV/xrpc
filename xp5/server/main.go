package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
	"xrpc"
)

type People struct {
}

func (p *People) Hello(msg string, reply *string) error {
	*reply = msg + "," + "123"
	fmt.Println("client send: " + msg)
	time.Sleep(2 * time.Second)
	return nil
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	ser := xrpc.NewServer(l)
	ser.Register(&People{})
	ser.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	ser.Close()
}
