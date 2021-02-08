package main

/*
两个goroutine，两个channel
	1.生成0~100的数字发送到ch1
	2.从ch1中取出数据计算它的平方，把结果发送到ch2中
*/

// 生成0~100的数字发送到ch1
// 通过单向通道限制函数的行为
func f1(ch chan int){
	for i:=0;i < 100;i ++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 chan int, ch2 chan int){
	// 循环取出ch1的值给ch2
	for {
		// 若ch1的值去完后ok返回的就是false
		tmp,ok := <- ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)

	go f1(ch1)
	go f2(ch1, ch2)
}
