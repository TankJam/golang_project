package main

import "fmt"

// 编写结构体
type person struct {
	name, city string
	age        int
}

func main() {
	// 1.键值对初始化
	p4 := person{
		name: "tankjam",
		age:  17,
	}
	fmt.Printf("%T\n", p4)
}
