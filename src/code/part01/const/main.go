package main

import "fmt"

// 常量
/// 常量定义后必须要赋值

// const pi = 3.1415
// const e = 2.7

// const (
// 	pi2 = 3.1314
// 	e2 = 2.8
// )

// const (
// 	n1 = 10
// 	n2 = 10
// 	n3 = 10
// )

// iota 主要用于枚举，每写一行都会累计 + 1
// const (
// 	n1 = iota // 0
// 	n2 = iota // 1
// 	n3 = iota // 2
// 	n4 = iota // 3
// )

// 中间插队，每新增一行都会 + 1
const (
	n1 = iota // 0
	n2 = iota // 1
	n3 = 100  // 100
	n4 = iota // 3
)

const (
	_  = iota             // 0
	KB = 1 << (10 * iota) // 1 << 10
	MB = 1 << (10 * iota) // 1 << 20
	GB = 1 << (10 * iota) // 1 << 30
	TB = 1 << (10 * iota) // 1 << 40
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 // 1, 2
	c, d = iota + 1, iota + 2 // 2, 3
)

func main() {
	// fmt.Println(n1, n2, n3, n4)
	// fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d)
}
