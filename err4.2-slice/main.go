package main

import "fmt"

func change(s ...int) {
	fmt.Printf("change befor slice len[%v] cap[%v] addr[%p]\n", len(s), cap(s), s)
	s = append(s, 3)
	fmt.Printf("change after slice len[%v] cap[%v] addr[%p]\n", len(s), cap(s), s)
	fmt.Printf("change value s[%v]\n", s)
}

func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	fmt.Printf("1 befor slice len[%v] cap[%v] addr[%p]\n", len(slice), cap(slice), slice)
	// 当切片长度不够，需要扩容时，会产生一个新地址；
	// 因此新地址是 1、2、0、0、0、3, 而原地址还是 1、2、0、0、0
	change(slice...)
	fmt.Printf("1 after slice len[%v] cap[%v] addr[%p]\n", len(slice), cap(slice), slice)
	// 此时，输出的是原地址
	fmt.Println(slice)
	fmt.Println()

	// 因为不需要扩容，所以，原地址输出的是 1、2、3、0、0
	change(slice[0:2]...)
	fmt.Println(slice)
}
