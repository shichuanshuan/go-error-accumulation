下面这段代码能否编译通过？如果可以，输出什么？

const (
	x = iota
	_
	y
	z = "zz"
	k 
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}
答案解析：
参考答案：编译通过，输出：0 2 zz zz 5