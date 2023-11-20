package main

import (
	"fmt"
	"net"
)

func main() {
	// 目标地址
	local := &net.UDPAddr{
		IP:   net.ParseIP("192.168.50.38"),
		Port: 8888,
	}

	address := &net.UDPAddr{
		IP:   net.ParseIP("192.168.50.38"),
		Port: 7777,
	}

	// 创建 UDP 连接
	conn, err := net.DialUDP("udp", local, address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("local ", conn.LocalAddr())
	fmt.Println("remote ", conn.RemoteAddr())

	// 发送数据
	message := "Hello, UDP Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("发送数据出错:", err)
		return
	}

	// 接收回复
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("接收回复出错:", err)
		return
	}

	fmt.Printf("收到服务器的回复：%s\n", string(buffer[:n]))
}
