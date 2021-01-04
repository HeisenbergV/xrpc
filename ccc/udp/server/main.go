package main

import (
	"fmt"
	"net"
)

func main() {
	//建立一个UDP的监听，这里使用的是ListenUDP，并且地址是一个结构体
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}

	fmt.Println("listen udp Start...:")

	for {
		var data [1024]byte
		//读取UDP数据
		count, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read udp failed, err:%v\n", err)
			continue
		}

		fmt.Printf("data:%s addr:%v count:%d\n", string(data[0:count]), addr, count)
		//返回数据
		_, err = listen.WriteToUDP([]byte("hello client"), addr)
		if err != nil {
			fmt.Printf("write udp failed, err:%v\n", err)
			continue
		}
	}
}
