package main

import "fmt"

// 指针  & 值 --> 地址     * 地址 ---> 值
// func main() {
// 	// a的值10在内存中
// 	a := 10 // int类型
// 	// 通过指针指向10的内存地址并赋值给b
// 	b := &a // &a  int类型(指针)
// 	fmt.Println(b)

// 	c := *b
// 	fmt.Println(c)
// }

// 在一个函数中修改另一个函数中变量的值
func modify(x int) {
	x = 100
}

// 接收指针
func modify2(y *int) {
	*y = 100
}

func main() {
	a := 1
	modify(a)
	fmt.Println(a) // 1，未能修改a的值

	// 修改一个函数中的变量，通过指针来实现
	modify2(&a)
	fmt.Println(a)
}
