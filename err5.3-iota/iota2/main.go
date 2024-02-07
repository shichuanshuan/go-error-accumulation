package main

import "fmt"

// 细节4 我们可以使用下划线跳过不想要的值。
const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

// 细节5 位掩码表达式
type Allergen int

const (
	loEggs         Allergen = 1 << iota // 1 << 0
	lgChocolate                         // 1 << 1
	lgNuts                              // 1 << 2
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

func main() {
	fmt.Printf("x[%v] y[%v] z[%v] k[%v] p[%v]\n", x, y, z, k, p)
}
