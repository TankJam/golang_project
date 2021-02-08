package main

import (
	"fmt"
	"strings"
)

// // map(映射)
// func main() {
// 	// 元素类型为map的切片
// 	// 支持完成了切片的初始化
// 	var mapSlice = make([]map[string]int, 8, 8)
// 	// nil, nil, nil, nil ...
// 	// 初始化得到的mapSlice的值等同于 8 个 nil
// 	fmt.Println(mapSlice[0] == nil)
// 	// 切片内的map并没有初始化，所以会报错
// 	// mapSlice[0]["tankjam"] = 100
// 	// fmt.Println(mapSlice)

// 	// 完成map的初始化
// 	mapSlice[0] = make(map[string]int, 8)
// 	mapSlice[0]["tankjam"] = 200
// 	fmt.Println(mapSlice[0])
// 	fmt.Println(mapSlice)
// }

// map(映射)
// func main() {
// 	// 值为切片的map
// 	var sliceMap = make(map[string][]int, 8) // 只完成了map的初始化
// 	// v, ok 若key有值则为 值, true， 否则 [], false
// 	v, ok := sliceMap["中国"] // 获取key为中国的键值对
// 	// fmt.Println(v, ok, "???????")
// 	if ok {
// 		fmt.Println(v)
// 	} else {
// 		// sliceMap 中没有中国这个键
// 		sliceMap["中国"] = make([]int, 8) // 完成了对切片的初始化
// 		sliceMap["中国"][0] = 100
// 		sliceMap["中国"][1] = 200
// 		sliceMap["中国"][3] = 300
// 	}
// 	// fmt.Println(sliceMap)
// 	// 遍历sliceMap
// 	for k, v := range sliceMap{
// 		fmt.Println(k, v)
// 	}
// }

func main(){
	// 统计一个字符串中每个单词出现的次数
	str1 := "how do you do"
	// 1. 定义一个map[string]int
	var wordCount = make(map[string]int, 10)
	// 2.字符串中有哪些单词
	words := strings.Split(str1, " ")
	// 3.遍历单词做统计
	for _, word := range words{
		// 以单词作为key取值
		v, ok := wordCount[word]
		if ok{
			// 若单词存在，则技术+1
			wordCount[word] = v + 1  
		}else{
			wordCount[word] = 1
		}
	}

	for k, v := range wordCount{
		fmt.Println(k, v)
	}
}
