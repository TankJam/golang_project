package main

import "fmt"


// 切片的扩容，append方法  append(数组, 添加的数值)
// 注意: 切片需要初始化后才能使用
func main(){
	var a []int  // 此时并没有申请内存
	// a[0] = 100
	// fmt.Println(a)  # 报错
	// append的返回值需要重新赋值给一个新的变量
	a = append(a, 10)
	fmt.Println(a)

	var b []int
	// 添加10次数值
	for i :=0; i < 10; i++ {
		// 切片会自动扩容， 从1开始，当容量不够时则自动 * 2的容量
		b = append(b, i)
		// ptr:内存地址
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", b, len(b), cap(b), b)
	}
}


// 切片扩容的策略
// 原来切片的长度 小于1024，就会自动扩容到2048
// 因为切片会自动扩容，所以以后我们定义切片时要
// 先想清楚切片最大的容量避免程序在运行时自动扩容导致的报错!
