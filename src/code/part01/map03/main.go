package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// 按照某个固定顺序遍历map
	var scoreMap = make(map[string]int, 100)

	// 创建一个map数据
	for i := 0; i < 50; i++ {
		// 注意: 格式化打印 Sprintf，  %02d，表示替换数字
		key := fmt.Sprintf("stu%02d", i)
		// 0~99的随机整数
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	// 遍历map
	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }

	// 按照key从小到大的顺序去遍历scoreMap
	// 1.先取出所有的key存到切片中
	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	// fmt.Println(scoreMap)
	// fmt.Println(keys)

	// 2.对key做排序
	// sort.Strings 对string类型的切片进行排序
	sort.Strings(keys) // keys目前是一个有序的切片

	// 3.按照排序后的key对scoreMap排序
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}
