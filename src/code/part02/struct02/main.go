package main

import "fmt"

// 结构体指针
type person struct {
	name, city string
	age        int8
}

func main() {
	var p2 = new(person)
	fmt.Printf("%T\n", p2)

	// (*p2).name = "坦克大宝贝"
	// (*p2).city = "上海"
	// (*p2).age = 18
	p2.name = "坦克大宝贝"
	p2.city = "上海"
	p2.age = 18
	// 查看p2指针的类型
	fmt.Printf("%#v\n", p2)

	// 取结构体的地址进行实例化
	p3 := &person{}
	p3.name = "tank jam"
	p3.age = 17
	p3.city = "虹口"
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)
}
