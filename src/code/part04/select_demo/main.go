package main

import "fmt"

// select多路复用
func main() {
	ch := make(chan int, 1)
	for i:=0;i<10;i++ {
		select {
		case x:= <-ch:
			fmt.Println(x)
			case ch <- i:
		default:
			fmt.Println("啥也不干!")
		}
	}
}
