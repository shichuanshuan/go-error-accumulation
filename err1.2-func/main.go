package main

import "fmt"

// 错误示例
// 创建的方法应该在 int 类型的包内
/*
func (i int) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i int = 1
	i.PrintInt()
}
*/

// 因为基于 Myint 创建的方法，和 Myint 定义在同一个包内，所有可以使用
type Myint int

func (i Myint) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i Myint = 1
	i.PrintInt()
}
