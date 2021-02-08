/*
构造初始化方法 __init__
*/
package main

import "fmt"

// 结构体构造函数: 构造一个结构体实例的函数
// 结构体是值类型
type person struct{
	name, city string
	age int8
}

// 构造函数: 类似于python中的__init__(),调用该函数时实例化对象
func newPerson(name, city string, age int8) person{
	return person{
		name: name,
		city: city,
		age: age,  // 注意需要加逗号
	}
}

//func main(){
//	name := "tank"
//	city := "上海"
//	age := int8(17)
//	// p 时初始化的对象
//	p := newPerson(name, city, age)
//	fmt.Println(name, city, age)
//	fmt.Printf("type:%T  value:%v#\n", p, p)
//	fmt.Println(p.name, p.city, p.age)
//}

/*

- 结构体使用指针类型:
- 什么时候使用指针类型?
	- 1.需要修改接受者中的值
	- 2.接收者是拷贝代价比较大的对象。
	- 3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/
// 编写结构体
type person2 struct{
	name, city string
	age int8
}

// 构造函数
func newPerson2(name, city string, age int8) *person2{
	return &person2{
		name: name,
		city:city,
		age: age,
	}
}

//func main(){
//	p2 := newPerson2("tank2", "china", int8(7))
//	fmt.Printf("type:%T  value:%#v", p2, p2)
//}



/*
嵌套结构体    类似于python的组合
- 结构体字段冲突问题
*/
type Address struct{
	Province string // 省份
	City string // 城市
	UpdateTime string
}

type Email struct {
	Addr string
	UpdateTime string
}

type Person3 struct {
	Name string
	Gender string
	Age int
	Address  // 嵌套另一个结构体
	Email  // 嵌套另一个结构体
}

//func main(){
//	p3 := Person3{
//		Name: "tank",
//		Gender: "男",
//		Age: 18,
//		Address: Address{
//			Province: "广东",
//			City: "湛江",
//			UpdateTime: "2021-01-07",
//		},
//		Email: Email{
//			Addr: "tankjam@163.com",
//			UpdateTime: "2020-01-07",
//		},
//	}
//
//	fmt.Printf("type: %T   value: %v#\n", p3, p3)
//	fmt.Println(p3.Address.Province, p3.Address.City)
//	fmt.Println(p3.Email.Addr, p3.Email.UpdateTime)
//}

/*
结构体的继承
*/
// 父结构体
type Animal struct{
	Name string
}

// Animal结构体有 move 方法
func (a *Animal) move(){
	fmt.Printf("%s 会动~\n", a.Name)
}

// 子结构体
type Dog struct {
	Feet int8
	// Dog结构体集成了Animal结构体中的属性与方法
	*Animal  // 匿名嵌套，而且嵌套的是一个结构体指针，实现继承
}

// 给结构体添加wang方法
func (d *Dog) wang(){
	fmt.Printf("%s 会汪汪叫~\n", d.Name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			Name: "图图",
		},
	}
	d1.wang()
	d1.move()
}



