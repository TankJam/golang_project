package main

import "fmt"


// 切片 （slice）
// 切片其实底层就是数组，由于数组的个数与类型都是固定的
// 所以出现了切片类型，切片类型与python中的list类型都是引用传递类型
// 而数组是值传递类型
func main(){
	// 基于数组得到的切片
	a := [5]int{55, 66, 77, 88, 99}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// 切片再次切片
	c := b[0: len(b)]
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	
	// make函数构造切片
	// 构造一个长度为5，容量为10的切片
	d := make([] int,5,10)
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	// 通过len函数获取切片的长度
	fmt.Println(len(d))

	// 通过cap() 函数获取切片的容量
	fmt.Println(cap(d))

	// 切片: 由 “指针” + “长度” + “容量”
	// 指针指向切片中的第一个索引的值的空间

	// nil
	var a2 []int  // 声明int类型切
	var b2 = []int{}  // 声明并初始化
	c2 := make([]int, 0)
	if a2 == nil {  
		fmt.Println("a 是一个 nil")
	}

	fmt.Println(a2, len(a2), cap(a2))
	if b == nil {
		fmt.Println("b 是一个 nil")
	}
	fmt.Println(b2, len(b2), cap(b2))

	if c2 == nil{
		fmt.Println("c 是一个 nil")
	}
	fmt.Println(c2, len(c2), cap(c2))

	// 切片是赋值拷贝，浅拷贝，数组是深拷贝
	a3 := make([]int, 3)  // [0,0,0]
	b3 := a3
	// 修改了b3，且影响了b3
	b3[0] = 100
	fmt.Println(a3, b3)

	// 切片的遍历
	a4 := []int{1, 2, 3, 4,5}
	for i:=0; i< len(a4); i++ {
		fmt.Println(i, a[i])
	}

	for _, v4 := range a4{
		fmt.Println(v4)
	}
}