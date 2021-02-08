package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 原子操作

var x int64
var l sync.Mutex
var wg sync.WaitGroup


// 普通版函数
func add(){
	x ++
	wg.Done()
}

// 互斥锁版
func mutexAdd(){
	l.Lock()
	x ++
	l.Unlock()
	wg.Done()
}

// 原子操作
func atomicAdd(){
	atomic.AddInt64(&x, 1)
	wg.Done()

}

func main() {
	start := time.Now()

	for i:=0; i < 10000; i++ {
		wg.Add(1)
		//go add()
		//go mutexAdd()  // 并发安全，牺牲性能
		go atomicAdd()  // 并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	println(end.Sub(start))

}
