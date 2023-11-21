package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

type Study struct {
	name string
}

func Name() (stu *Study) {
	// 错误示范
	/*
		stu.name = "shics"
	*/

	stu = &Study{
		name: "shics",
	}
	return
}

func main() {

	var stu1 *Study

	stu1 = Name()

	fmt.Println(stu1)

}