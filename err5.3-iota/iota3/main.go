package main

import "fmt"

// 细节6 定义在一行的情况
const (
	Apple, Banana     = iota + 1, iota + 3 // iota = 0
	Cherimoya, Durian                      // iota = 1
	Elderberry, Fig                        // iota = 2
)

// 细节7 中间插队
const (
	i = iota
	j = 3.14
	k = iota
	l
)

// 下面这段代码能否编译通过？如果可以，输出什么？
func main() {
	fmt.Printf("Apple[%v], Banana[%v], Cherimoya[%v], Durian[%v], Elderberry[%v], Fig[%v]\n", Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
	fmt.Println("==============================")
	fmt.Printf("i[%v] j[%v] k[%v] l[%v]", i, j, k, l)
}
