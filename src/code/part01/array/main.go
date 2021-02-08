package main

import "fmt"

// 数组相关内容
func main(){
	var a [3] int
	var b [4] int
	// a=b
	fmt.Println(a)
	fmt.Println(b)

	// 数组初始化
	// 1.定义时使用初始值列表的方式初始化
	// 注意：数组的长度必须是常亮
	var cityArray = [4]string{"北京", "深圳", "上海", "广州"}
	fmt.Println(cityArray)

	// 2.编译器推到数组的长度
	var boolArray = [...]bool{true, false, true}
	fmt.Println(boolArray)

	// 3.使用索引值方式初始化
	var langArray = [...]string{1: "golang", 2: "Python"}
	fmt.Println(langArray)
	fmt.Printf("%T", langArray)

	// 数组遍历
	// 1.for循环遍历
	for i :=0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	// 2.for range 遍历
	// range 遍历出数组中的索引与值
	// go语言中 _ 就是废弃的值
	for _, value := range cityArray{
		fmt.Println(value)
	}

	// 二维数组
	// 注意: 多维数组只有外层才能简写数组的长度，内层允许使用 [...]
	cityArray2 := [4][2]string{
		{"北京1", "西安1"},
		{"北京2", "西安2"},
		{"北京3", "西安3"},
		{"北京4", "西安4"},
	}
	fmt.Println(cityArray2)

	// 二维数据遍历
	for _, v1 := range cityArray2{
		fmt.Println(v1)
		for _, v2 := range v1{
			fmt.Println(v2)
		}
	}

	// 数组是值类型 
	x := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	f1(x)
	// 数组是值传递类型，不是引用传递，与python的list类型不一样
	y := x 
	y[0][0] = 1000
	fmt.Println(x)
	f2()
	f3()
}

func f1(a [3][2]int){
	// fmt.Println(a)
	// 由于是值传递，所以调用函数重新赋值后原数组值不变
	a[0][0] = 300
	// fmt.Println(a)
}


// 练习题
// 1.球数组 [1, 3, 5, 7, 8] 所有元素的和
func f2(){
	arrary1 := [...]int {1, 3, 5, 7, 8}
	fmt.Println(arrary1)
	number := 0
	for _, v3 := range arrary1{
		number += v3
	}
	fmt.Println(number)
}

// 2.找出数组中和为指定值的两个元素下标，比如从
// 数组[1,3,5,7,8]中找出和为8的两个元素的下表分别为(0, 3)和(1,2)
// func f3(){
// 	arrary1 := [...]int {1, 3, 5, 7, 8}
// 	n1, n2 := 0, 0
// 	for index, v3 := range arrary1{
// 		if index == 0{
// 			n1 = v3
// 		}else{
// 			v2 = n1 + v3
// 		}
// 	}
	
// }

