package main

import "fmt"


// map（映射）  
func main(){
	// 仅声明map，但是没有初始化，a的初始值就是nil
	var a map[string]int
	fmt.Println(a == nil)
	// map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	// map中添加键值对  类似于python的dict
	a["tank jam"] = 100
	a["tank zhang"] = 200
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)

	// 声明map同时并初始化
	b := map[int]bool{
		1:true,
		2:false, 
	} // 注意，需要再末尾加, 否则花括号需要跟在值的最后面
	fmt.Println("b: %#v\n", b)
	fmt.Printf("type:%T\n", b)
	
	// 声明的变量必须要使用，否则会报错
	// var c map[int]int
	// c[100] = 200 // 报错。因为map没有初始化
	// fmt.Println(c)

	// 判断某个键是否存在
	// map[key_type]value_type
	var scoreMap = make(map[string]int, 8)
	scoreMap["tank jam"] = 100
	scoreMap["tank Z"] = 200

	// 判断 张二狗子 在不在scoreMap中
	value, ok := scoreMap["张二狗子"]
	fmt.Println(value, ok)
	if ok{
		fmt.Println("张二狗子在", value)
	}else{
		fmt.Println("查无此人")
	}
}