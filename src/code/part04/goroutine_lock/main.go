package main

import (
	"fmt"
	"sync"
)

// 多个goroutine并发操作全局变量x
var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
)

func add(){
	// 可能导致x的值不一致，所以需要 加锁，将并发转为并行，牺牲效率，保证数据安全
	for i:=0;i<5000;i++{
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	// 要添加异步任务
	wg.Add(2)
	go add()
	go add()

	// 等待所有
	wg.Wait()

	fmt.Println(x)
}


// 读写锁，当写多用写锁，当读多用读锁
