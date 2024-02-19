package main

import "fmt"

func main1() {
	s := make([]int, 5)
	fmt.Printf("%p\n", s)
	s[0] = 0
	s[1] = 1
	fmt.Printf("%p\n", s)
	fmt.Println(s)
}

func main() {
	s := make([]int, 0)
	fmt.Printf("%p\n", s)
	s = append(s, 1, 2, 3, 4)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
	fmt.Println("=============")

	s1 := make([]int, 5)
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Printf("%p\n", s1)
	fmt.Println(s1)
}
