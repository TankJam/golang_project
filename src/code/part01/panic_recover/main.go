package main

import "fmt"

func a() {
	fmt.Println("func a")
}

func b() {
	// defer可以在异常触发后做某些操作
	defer func() {
		// 从当前的异常中获取错误信息，获取panic中的错误信息
		err := recover()
		if err != nil {
			fmt.Println("func b error")
			fmt.Println(err)
			fmt.Printf("%T\n", err)
		}
	}()
	// panic让程序异常退出，并自定义输出的错误信息
	panic("panic in b")
}

func main() {
	b()
}
