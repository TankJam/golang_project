package main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包
// 定义一个函数它的返回值是一个函数
// 把函数作为返回值
func a(name string) func() {
	// name := "坦克大宝贝"
	return func() {
		fmt.Println("hello tank jam~~")
		fmt.Println(name)
	}
}

// 类似于python的装饰器
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		// 判断name结尾是否是
		if !strings.HasSuffix(name, suffix) {
			// 是则返回 name + 后缀
			return name + suffix
		}
		return name
	}
}


func main() {
	// 闭包函数 = 内层函数 + 外层变量的引用
	// r := a("坦克") // r此时就是一个闭包函数
	// r()          // 相当于执行了a函数内部的匿名函数

	// 闭包检测文件后缀名
	r := makeSuffixFunc(".txt")
	ret := r("坦克")
	fmt.Println(ret)
}
