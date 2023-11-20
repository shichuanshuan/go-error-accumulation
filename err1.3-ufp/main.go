go ���� udp ����� close �󣬼�����ַ��� 0.0.0.0 ��ô����

��Go�����У�������UDP�������е���Close()�����ر����Ӻ󣬷���˼����ĵ�ַ����0.0.0.0��������Ϊ��Close()����������ʱ���ײ���������ӱ��رգ��������������ĵ�ַ���ͷš�

�����ϣ���ڹر�UDP�������󣬼�����ַ�Ա���ԭ����������޸Ĵ��룬�ȹر����ݰ�����ѭ����Ȼ���ٹر����ӡ��������������ڹر����Ӻ󣬼�����ַ�Ͳ�����0.0.0.0��

������ʾ�����룺

go
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// ������ַ
	addr := "localhost:12345"

	// ����UDP����
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// ��׽�ж��źţ�ֹͣ������
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// ��������
	buffer := make([]byte, 1024)
	for {
		select {
		case <-stop:
			fmt.Println("ֹͣUDP������")
			return
		default:
			n, src, err := conn.ReadFrom(buffer)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Received from %s: %s\n", src.String(), string(buffer[:n]))
		}
	}
}
�������ʾ���У�����ʹ��net.ListenPacket()��������UDP���ӣ�������ʹ��net.ListenUDP()��Ȼ������ѭ����ʹ��conn.ReadFrom()�����������ݰ��������յ��ж��ź�ʱ����ӡһ����Ϣ�����أ��Ӷ�ֹͣ��������

����������ر�UDP�������󣬼�����ַ�Ͳ�����0.0.0.0�����Ǳ���ԭ����