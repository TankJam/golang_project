package main

import "fmt"

// 全局变量
var n = 100

// 变量
func main() {
	// 标准声明格式
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)

	// 批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)

	// 声明变量同时指定初始值
	var name1 string = "tank jam"
	var age1 int = 18
	fmt.Println(name1, age1)
	var name2, age2 = "坦克", 18
	fmt.Println(name2, age2)

	// 类型推导，让 “编译器” 根据变量的初始值推到出其他类型，“不需要指定类型”
	var name3 = "帅气坦克"
	var age3 = 17
	fmt.Println(name3, age3)

	// 短变量声明，只能函数内部声明，类似于局部名称空间的变量
	m := 10
	fmt.Println(m)
	fmt.Println(n)

	// 匿名变量， “_” 命名变量

	// 声明变量注意:
	/// 1.函数外的每个词语必须以关键字开始 (var, const, func 等...)
	/// 2. := 不能使用在函数外
	/// 3. _ 多用于占位，表示忽略值

}
