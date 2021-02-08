package main

import "fmt"

// defer: 延迟执行，函数中有多个defer相当于从上到下去注册，
// 当defer以为的代码执行完毕后，从下到上执行defer

func main() {
	fmt.Println("start ...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end ...")
}
