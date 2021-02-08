package main

import (
	"fmt"
	"strings"
)

// 字符串常用操作
func main() {
	// 求字符串的长度
	s := "hello tank"
	fmt.Println(s)
	s2 := "hello 小豆儿"
	fmt.Println(len(s2))

	// 拼接字符串
	fmt.Println(s + s2)
	s3 := fmt.Sprintf("%s - %s", s, s2)
	fmt.Println(s3)

	// 字符串分割
	s4 := "how do you do"
	// 字符串内置模块
	fmt.Println(strings.Split(s4, " "))

	// 判断是否包含
	fmt.Println(strings.Contains(s4, "do")) // true

	// 判断前缀
	fmt.Println(strings.HasPrefix(s4, "how")) // true

	// 判断后缀
	fmt.Println(strings.HasSuffix(s4, "do")) // true

	// 判断字符串的位置
	fmt.Println(strings.Index(s4, "do"))

	/// 最后字符串出现的位置
	fmt.Println(strings.LastIndex(s4, "do"))

	// join
	/// 给数组一次性添加多个值
	s5 := []string{"how", "do", "you", "do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5, "+"))

}
