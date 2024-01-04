[参考链接]
https://github.com/KatelynHaworth/go-tproxy

package main

import (
    "log"
    "net"
    "os"
    "syscall"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("Failed to start listener: %s", err)
    }

    file, err := listener.(*net.TCPListener).File()
    if err != nil {
        log.Fatalf("Failed to get file descriptor: %s", err)
    }
    defer file.Close()

    fd := int(file.Fd())

    // 设置透明传输选项
    err = syscall.SetsockoptInt(fd, syscall.SOL_IP, syscall.IP_TRANSPARENT, 1)
    if err != nil {
        log.Fatalf("Failed to set IP_TRANSPARENT option: %s", err)
    }

    // 设置非阻塞模式
    err = syscall.SetNonblock(fd, true)
    if err != nil {
        log.Fatalf("Failed to set non-blocking mode: %s", err)
    }

    // 创建文件描述符的net.Listener
    listener, err = net.FileListener(file)
    if err != nil {
        log.Fatalf("Failed to create listener from file descriptor: %s", err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            // 处理非阻塞错误
            if opErr, ok := err.(*net.OpError); ok && opErr.Temporary() {
                continue
            }

            log.Printf("Error accepting connection: %s", err)
            break
        }

        // 处理连接
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // 处理连接逻辑
}
