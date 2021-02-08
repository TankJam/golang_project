package main

import "fmt"

// 函数变量的作用域

// 定义全局变量
// var number1 = 10

// 定义函数
func testGlobal() {
	number1 := 100
	fmt.Println(number1)
}

// 函数对象
func test(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	testGlobal()

	// fmt.Println(number1)

	// 变量i此时只是在for循环语句中生效
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	abc := testGlobal
	fmt.Printf("%T\n", abc)

	res := test(100, 200, add)
	fmt.Println(res)

	res2 := test(100, 200, sub)
	fmt.Println(res2)

}
