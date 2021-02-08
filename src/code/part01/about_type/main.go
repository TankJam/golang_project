package main

import "fmt"

// 自定义类型和类型的别名示例

// 1.自定义类型
// MyInt 基于int类型的自定义类型

type Myint int

// 2.类型别名
// NewInt int类型别名
type NewInt = int

func main() {
	var a Myint
	var b NewInt
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
