package main
// 读写锁，当写多用写锁，当读多用读锁

import (
	"fmt"
	"sync"
	"time"
)

// 多个goroutine并发操作全局变量x
var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	//rwLock sync.RWMutex
)

func read(){
	lock.Lock()
	//rwLock.Lock()
	time.Sleep(time.Microsecond)
	lock.Unlock()
	//rwLock.Unlock()
	wg.Done()
}

func write(){
	lock.Lock()
	//rwLock.Lock()
	x = x + 1
	lock.Unlock()
	//rwLock.Unlock()
	time.Sleep(time.Microsecond * 10)
	wg.Done()
}

func main() {
	start := time.Now()

	for i:=0;i<1000;i++{
		wg.Add(1)
		go read()
	}

	for i:=0;i<10;i++{
		wg.Add(1)
		go write()
	}

	wg.Wait()
	pay_time := time.Now().Sub(start)
	fmt.Println(pay_time)
}

// sync.One在确保需要某一个函数任务只执行一次close关闭时，可以使用

// Go语言中，内置的map不是线程安全的，不会保证数据的安全


