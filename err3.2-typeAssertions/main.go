package main

import (
	"fmt"
	"net"
)

func main1() {
	// 创建一个 net.Addr 类型的变量并赋值为 UDPAddr 类型的值
	var addr net.Addr = &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}

	// 在进行类型断言之前，最好使用类型断言的形式 value, ok := expression.(Type) 来进行安全的类型检查。
	// 将 addr 断言为 UDPAddr 类型，并进行类型检查
	if udpAddr, ok := addr.(*net.UDPAddr); ok {
		// 类型断言成功，获得了一个 net.UDPAddr 类型的值
		fmt.Println("IP:", udpAddr.IP)
		fmt.Println("Port:", udpAddr.Port)
		fmt.Println(udpAddr.AddrPort())
	} else {
		// 类型断言失败
		fmt.Println("addr 不是 net.UDPAddr 类型")
	}
}

func main() {
	var addr net.Addr = &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 2323,
	}

	if udpAddr, ok := addr.(*net.UDPAddr); ok {
		fmt.Println(udpAddr.IP)
	}
}
