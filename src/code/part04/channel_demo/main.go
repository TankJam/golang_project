package main

// channel的使用，让不同的线程中实现数据共享

import "fmt"

func main() {
	// 引用类型，需要初始化之后才能使用
	// 相当于生成了一个通道
	// 无缓冲区通道，会报错
	//ch1 := make(chan int)，
	// 有缓冲区通道，不报错
	ch1 := make(chan int, 1)  // 1代表，可以存放值的数量
	// 往通道发送值
	ch1 <- 1
	// 从通道中获取值
	x:= <- ch1
	fmt.Println(x)

	// 获取通道的数量与容量
	fmt.Println(len(ch1))
	fmt.Println(cap(ch1))

	// 关闭通道
	close(ch1)
}
