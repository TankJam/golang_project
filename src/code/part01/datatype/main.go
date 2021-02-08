package main

import (
	"fmt"
	"math"
)

// 基本数据类型

func main() {
	// 整型
	// 十进制打印为二进制
	n := 10
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)

	// 八进制
	m := 075
	fmt.Printf("%d\n", m)
	fmt.Printf("%o\n", m)

	// 十六进制
	f := 0xff
	fmt.Println(f)
	fmt.Printf("%x\n", f)

	// uint9
	var age uint8 // 0~255
	age = 255
	fmt.Println(age)

	// 浮点数
	/// go语言自带的float包
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	// 布尔值
	/// 无法进行运算与类型转换
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	///  cannot use 3333 (type untyped int) as type bool in assignment
	// a = 3333

	// 字符串
	s1 := "hello GD"
	s2 := "你好 广东"
	fmt.Println(s1, s2)
	/// 字符串转义符
	//// 打印window平台下的一个路径 c:\code\go.exe
	fmt.Println("c:\\code\\go.exe") // \\转义
	fmt.Println("\t制表符\n换行符")       // 转义符 \t tab, \n换行
	/// 打印多行字符串
	s3 := `
	从前现在 过去再不来
	红尘落叶永埋尘土内
	`
	fmt.Println(s3)
}
