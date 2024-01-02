[答案]
map[string]Student 的 value 是一个 Student 结构值，所以当list["student"] = student,是一个值拷贝过程。而list["student"]则是一个值引用。那么值引用的特点是只读。所以对list["student"].Name = "LDB"的修改是不允许的。

方法一：

package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	//list["student"].Name = "LDB"

    /*
        方法1:
    */
    tmpStudent := list["student"]
    tmpStudent.Name = "LDB"
    list["student"] = tmpStudent

	fmt.Println(list["student"])
}
其中

/**
方法1:
*/
tmpStudent := list["student"]
tmpStudent.Name = "LDB"
list["student"] = tmpStudent
是先做一次值拷贝，做出一个tmpStudent副本,然后修改该副本，然后再次发生一次值拷贝复制回去，list["student"] = tmpStudent,但是这种会在整体过程中发生 2 次结构体值拷贝，性能很差。

方法二：

package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]*Student

func main() {

	list = make(map[string]*Student)

	student := Student{"Aceld"}

	list["student"] = &student
	list["student"].Name = "LDB"

	fmt.Println(list["student"])
}
我们将 map 的类型的 value 由 Student 值，改成 Student 指针。

var list map[string]*Student
这样，我们实际上每次修改的都是指针所指向的 Student 空间，指针本身是常指针，不能修改，只读属性，但是指向的 Student 是可以随便修改的，而且这里并不需要值拷贝。只是一个指针的赋值。
