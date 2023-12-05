package main

import "fmt"

type Person struct {
	name string
	age  int
}

func MySelf() (pe *Person) {
	// 错误示范
	//pe.name = "shi"
	//pe.age = 3

	pe = &Person{
		name: "shi",
		age:  25,
	}

	return
}

func main() {
	p := MySelf()
	fmt.Println(*p)
}
