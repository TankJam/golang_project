package main

import "fmt"

func main(){
	// 遍历map
	var scoreMap = make(map[string]int, 8)
	scoreMap["tankjam"] = 100
	scoreMap["tankz"] = 200

	// for range
	// map是无序的，键值对添加的顺序无关
	for k, v := range scoreMap{
		fmt.Println(k, v)
	}

	// 只遍历map中的key
	for k:=range scoreMap{
		fmt.Println(k)
	}

	// 只遍历map中的value
	for _, v:= range scoreMap{
		fmt.Println(v)
	}

	// 删除 delete(map变量, key)
	fmt.Println(scoreMap)
	delete(scoreMap, "tankjam")
	fmt.Println(scoreMap)

}