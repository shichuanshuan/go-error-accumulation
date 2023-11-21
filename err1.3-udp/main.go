package main

import (
	"fmt"
	"log"
	"net"
)

/*
go 语言 udp 服务端 close 后，监听地址变成 0.0.0.0 怎么回事

在Go语言中，当你在UDP服务器中调用Close()方法关闭连接后，服务端监听的地址会变成0.0.0.0。这是因为在Close()方法被调用时，底层的网络连接被关闭，而服务器监听的地址被释放。

如果你希望在关闭UDP服务器后，监听地址仍保持原样，你可以修改代码，先关闭数据包接收循环，然后再关闭连接。这样，服务器在关闭连接后，监听地址就不会变成0.0.0.0。

以下是示例代码：
*/

func main() {
	// 监听地址
	addr := "localhost:12345"

	// 错误方法创建
	// 创建 UDP 地址结构
	udpAddress, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// 创建 UDP 连接
	conn, err := net.ListenUDP("udp", udpAddress)

	// 创建UDP连接
	//conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 接收数据
	buffer := make([]byte, 1024)
	for {
		fmt.Printf("start read\n")
		n, src, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received from %s: %s\n", src.String(), string(buffer[:n]))
	}
}

/*
在上面的示例中，我们使用net.ListenPacket()函数创建UDP连接，而不是使用net.ListenUDP()。然后，在主循环中使用conn.ReadFrom()方法接收数据包。当接收到中断信号时，打印一条消息并返回，从而停止服务器。

这样，当你关闭UDP服务器后，监听地址就不会变成0.0.0.0，而是保持原样。
*/
