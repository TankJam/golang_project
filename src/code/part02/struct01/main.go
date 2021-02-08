package main

import "fmt"

/*
	结构体：Go语言中实现面向对象编程
		- 相当于python中的类
	- 结构体的定义:
		- type 和 struct关键字来定义结构体
			type 类型名 struct{
				字段名 字段类型
				字段名 字段类型
				...
			}

	- 类型名: 标识自定义结构体的名称，在同一个包内不能重复。
	- 字段名: 表示结构体字段名, 结构体中的字段名必须唯一。
	- 字段类型: 表示结构体字段的具体类型。
*/

// 定义结构体
type person struct {
	name, city string
	age        int8
}

// 结构体的实例化: 只有实例化才会开辟内存
func main() {
	var pObj person
	pObj.name = "tank"
	pObj.age = 18
	pObj.city = "上海"
	fmt.Println(pObj.name, pObj.age, pObj.city)

	// 匿名结构体
	var user struct{
		name string
		married bool
	}

	user.name = "tankjam"
	user.married = false

	fmt.Println(user, user.name, user.married)

}
