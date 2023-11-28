package main

import "fmt"

type porcAttr struct {
	Dir string
	Env string
}

type sysAttr struct {
	Dir string
	Env string
}

func main() {

	var pAttr = new(porcAttr)
	pAttr.Dir = "dir"
	pAttr.Env = "env"

	// 只是把 pAttr 的值给到 sAttr
	sAttr := &sysAttr{
		Dir: pAttr.Dir,
		Env: pAttr.Env,
	}

	fmt.Printf("parrt dir %p\n", &pAttr.Dir)
	fmt.Printf("parrt Env %p\n", &pAttr.Env)
	fmt.Printf("sAttr dir %p\n", &sAttr.Dir)
	fmt.Printf("sAttr Env %p\n", &sAttr.Env)

	fmt.Printf("\npAttr %p\n", &pAttr)
	fmt.Printf("sAttr %p\n", &sAttr)
}
