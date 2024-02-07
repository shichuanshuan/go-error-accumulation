package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	// map中的值是通过值复制的方式进行存储的。当你将一个结构体类型的值存储在 map 中时，实际上是将该结构体的副本存储在了map中，而不是原始结构体本身。
	// 所以，当你使用list["student"]获取map中的值时，你实际上获得的是student结构体的一个副本，并不是原始结构体本身。
	// 因此，尝试通过list["student"].Name去修改副本的字段值是无效的，不会对原始结构体产生影响。
	list["student"] = student
	list["student"].Name = "LDB"

	fmt.Println(list["student"])
}
