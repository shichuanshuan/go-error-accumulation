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
	ss := new(studyInfo)
	ss.name = "shics"
	fmt.Printf("ss ptr[%v] ss value[%v] ss[%v]\n", &ss, *ss, ss)
	ss = setStudyName()
	fmt.Printf("ss ptr[%v] ss value[%v] ss[%v]", &ss, *ss, ss)
}
