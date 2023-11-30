package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {

		// 无论 select 最终选择了哪个 case，getVal() 都会按照源码顺序执行：getVal(1) 和 getVal(2)，也就是它们必然先输出：
		select {
		case ch <- getVal(1):
			fmt.Println("in first case")
		case ch <- getVal(2):
			fmt.Println("in second case")
		default:
			fmt.Println("default")
		}
	}()

	fmt.Println("The val:", <-ch)
}

func getVal(i int) int {
	fmt.Println("getVal, i=", i)
	return i
}
