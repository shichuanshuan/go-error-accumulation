[错误原因]
存在一个结构体 study; 定义结构体 study 为变量 one；
在调用 SetFunc 函数时，SetFunc中定义结构体 study 为变量 one2， 并将 one2 返回给 one；
这时 one 中赋值过的变量会被清除掉

[结论]
定义过的变量，当作为函数接收值时，原来的数据会被覆盖