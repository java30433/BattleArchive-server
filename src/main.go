package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:99")
	if err != nil {
		fmt.Println("服务器启动失败：", err)
		return
	}
	go ClientsSync()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端连接失败：", err)
			continue
		}
		fmt.Println("新客户端连接")
		AddClient(conn)
	}
}
