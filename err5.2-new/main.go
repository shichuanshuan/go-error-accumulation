package main

import "fmt"

// new 用于创建值类型（如结构体）的实例，并返回指向该实例的指针。
// 它分配了零值填充的内存，并返回一个指向这块内存的指针。new 的语法是 new(T)，其中 T 是要创建的类型
func main1() {
	// 创建一个整数指针，并初始化为零值
	numPtr := new(int)

	fmt.Println(*numPtr) // 打印指针指向的值（零值）
	*numPtr = 42         // 更新指针指向的值
	fmt.Println(*numPtr) // 打印更新后的值
}

// make 用于创建引用类型（如切片、映射和通道）的实例。它分配了内部数据结构，并进行了适当的初始化，
// 返回一个已准备好使用的实例。make 的语法是 make(T, args)，其中 T 是要创建的类型，args 是用于初始化该类型的参数。
func main2() {
	// 创建一个切片，并初始化为指定的长度和容量
	slice := make([]int, 5, 10)

	fmt.Println(len(slice)) // 打印切片的长度
	fmt.Println(cap(slice)) // 打印切片的容量
}

func main() {
	arr1 := make([]int, 5)
	arr2 := make([]int, 5)

	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	fmt.Printf("%p %p\n", arr1, arr2)
	arr2[1] = arr1[1]
	//arr2[0] = 7

	fmt.Println("arr1=", arr1, "arr2=", arr2)
	fmt.Printf("%p %p\n", arr1, arr2)
}
