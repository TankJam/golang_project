package main

import "fmt"

func main() {

	var a *int
	// invalid memory address or nil pointer dereference
	// 不能对空指针进行赋值，需要使用new来做初始化
	a = new(int) //new函数用得比较少，返回的是指针
	// make返回的是具体的类型
	*a = 200
	fmt.Println(*a)

	// map，是引用类型，其实已经有了内存空间，可以修改属性
	var b map[string]int
	b["tank jam"] = 300
	fmt.Println(b)
}

/*
	- new与make的区别:
		1.二者都是用来做内存分配的。
		2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身。
		3.而new用于类型的内存分配，比如int、string...，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
