package main

import "fmt"

// 无参函数
func sayHello() {
	fmt.Println("sayHello")
}

// 定义接收string类型的有参函数
func sayHello2(name string) {
	fmt.Println("Hello", name)
}

// 定义接收多个参数的函数并且一个返回值
// 算两个数的和,()括号内的是接收的参数，括号后面是返回值的类型
func intSum(a int, b int) int {
	sumValue := a + b
	return sumValue
}

// 返回值简写方式
func intSum2(a int, b int) (res int) {
	res = a + b
	// 自动返回res
	return
}

// 接收可变参数，在参数名后面加... 表示可变参数
// 可变参数在函数体重是切片类型
func intSum3(a ...int) int {
	ret := 0
	for _, arg := range a {
		ret += arg
	}
	return ret
}

// 接收参数简写
func intSum5(a, b int) (ret int) {
	ret = a + b
	return ret
}

// 多返回值函数, 也接收类型简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

// 函数
func main() {
	// 函数调用
	// sayHello()

	// name := "坦克大宝贝"
	// sayHello2(name)

	// res := intSum(100, 200)
	// fmt.Println(res)

	// 调用函数
	// ret := intSum3(100, 200, 300, 400)
	// fmt.Println(ret)

	// ret2 := intSum3()
	// ret3 := intSum3(20)
	// ret4 := intSum3(10, 20)
	// ret5 := intSum3(30, 40, 50)
	// fmt.Println(ret2, ret3, ret4, ret5)

	x, y := calc(100, 200)
	fmt.Println(x, y)
}
