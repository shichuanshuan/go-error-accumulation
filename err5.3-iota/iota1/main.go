package main

import "fmt"

// 细节1 iota 只能在常量的表达式中使用
//fmt.Println(iota)  错误

// 细节2 每次 const 出现时，都会让 iota 初始化为0
const (
	zero = iota
	zero1
	zero2 = iota
)
const one = iota

// 细节3 自增长常量经常包含一个自定义枚举类型，允许你依靠编译器完成自增设置。
type Stereotype int

const (
	s1 Stereotype = iota
	s2
	s3
)

func main() {
	fmt.Printf("zero[%v] zero1[%v] zero2[%v] one[%v]\n", zero, zero1, zero2, one)

	fmt.Println("============")
	fmt.Printf("s1[%v] s2[%v]\n", s1, s2)
}
