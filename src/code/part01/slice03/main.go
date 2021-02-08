package main

import (
	"fmt"
	"sort"
)


// 切片的copy
func main(){
	a := []int{1, 2, 3, 4,5}
	b := make([]int, 5, 5)
	c := b
	// copy(被拷贝元素, 拷贝元素)
	copy(b, a)
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 切片的元素删除
	a2 := []string{"guangzhou","shenzhen", "beijing", "shanghai"}
	// 删除第三个beijing
	// 先窃取a2中的第一个第二个值，然后将a2中的最后一个值添加到a2中
	a2 = append(a2[0:2], a2[3:]...)
	fmt.Println(a2)

	// append(a2[:index], a[index+1:]...)

	// 注意: 切片不能用来直接比较

	// 排序:
	var a3 = [...]int{3, 7, 8, 9, 1}
	// a3[:]得到的是一个切片，指向了底层的数组a3
	sort.Ints(a3[:])
	fmt.Println(a3)
}

