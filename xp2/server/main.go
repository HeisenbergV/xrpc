package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"xrpc"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	ser := xrpc.NewServer(l)
	ser.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	ser.Close()
}
