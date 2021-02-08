package main

import (
	"fmt"
	"sync"
)

// sync.map
// 内置map不是线程安全的，会错乱

//var (
//	wg sync.WaitGroup
//)
//
//var m = make(map[int]int)
//
//func get(key int) int {
//	return m[key]
//}
//
//func set(key int, value int){
//	m[key] = value
//}
//
//func main() {
//
//	// 报错，因为map不是线程安全的
//	for i :=0;i < 20; i++ {
//		wg.Add(1)
//		go func(i int) {
//			set(i, i + 100)  // map中设置键值对
//			fmt.Printf("key: %v value:%v", i, get(i))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//
//
//}

// 使用sync内的map来实现

var (
	wg sync.WaitGroup
)

// 可以传任意类型，因为是空间类型
var m2 = sync.Map{}

func main() {

	// 报错，因为map不是线程安全的
	for i :=0;i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+100)  // map中设置键值对
			value, _ := m2.Load(i)
			fmt.Printf("key: %v value:%v\n", i, value)
			wg.Done()
		}(i)
	}
	wg.Wait()


}
