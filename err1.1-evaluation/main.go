package main

import "fmt"

type studyInfo struct {
	name string
	age  int
}

func setStudyName() (ss *studyInfo) {
	ss = new(studyInfo)
	ss.age = 26

	return ss
}

func main() {
	// 定义一个结构体变量，对结构体赋值
	ss := new(studyInfo)
	ss.name = "shics"
	fmt.Printf("ss ptr[%v] ss value[%v] ss[%v]\n", &ss, *ss, ss)

	// result
	// 函数中定义结构体变量，并设置变量的值，将结构体返回，若返回接收的变量是刚才定义的，则会覆盖刚才定义的变量
	ss = setStudyName()
	fmt.Printf("ss ptr[%v] ss value[%v] ss[%v]", &ss, *ss, ss)
}
