package main

import (
	"fmt"
	"net"
)

func main() {
	// ����һ�� net.Addr ���͵ı�������ֵΪ UDPAddr ���͵�ֵ
	var addr net.Addr = &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}

	// �ڽ������Ͷ���֮ǰ�����ʹ�����Ͷ��Ե���ʽ value, ok := expression.(Type) �����а�ȫ�����ͼ�顣
	// �� addr ����Ϊ UDPAddr ���ͣ����������ͼ��
	if udpAddr, ok := addr.(*net.UDPAddr); ok {
		// ���Ͷ��Գɹ��������һ�� net.UDPAddr ���͵�ֵ
		fmt.Println("IP:", udpAddr.IP)
		fmt.Println("Port:", udpAddr.Port)
	} else {
		// ���Ͷ���ʧ��
		fmt.Println("addr ���� net.UDPAddr ����")
	}
}