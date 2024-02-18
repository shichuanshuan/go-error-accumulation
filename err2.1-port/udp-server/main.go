package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听地址
	address := "192.168.50.38:7777"

	// 创建 UDP 连接
	conn, err := net.ListenPacket("udp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Printf("UDP 服务器已启动，监听地址：%s\n", address)

	// 接收数据
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("接收数据出错:", err)
			continue
		}

		fmt.Printf("接收到来自 %s 的数据：%s\n", addr.String(), string(buffer[:n]))

		// 回复数据
		_, err = conn.WriteTo([]byte("已收到你的消息"), addr)
		if err != nil {
			fmt.Println("回复数据出错:", err)
		}
	}
}
