package main

import (
	"fmt"
	"log"
	"xrpc"
)

func main() {

	client, _ := xrpc.Dial("tcp", "127.0.0.1:8889")

	for i := 0; i < 5; i++ {
		args := fmt.Sprintf("sync req %d", i)
		var reply string
		if err := client.Call("Foo.Sum", args, &reply); err != nil {
			log.Fatal("call Foo.Sum error:", err)
		}
		log.Println("reply:", reply)
	}

	//异步
	// args := fmt.Sprintf("async req %d", 66)
	// var reply string
	// call := client.Go("Foo.Sum", args, &reply, nil)
	// c := <-call.Done
	// log.Println("reply:", c.Reply)

}
