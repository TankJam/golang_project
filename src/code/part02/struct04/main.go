/*
结构体字段可见性和json
	- 重点，Go语言数据转为json
 */
package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Id int
	Name string
}

// student类构造函数
func newStudent(id int, name string)student{
	return student{
		Id: id,
		Name: name,
	}
}

type class struct {
	Title string   `json:"title"`
	// 数组中存对象
	Students []student  `json:"student_list" db:"student" xml:"ss"`
}



func main(){
	// 创建一个班级变量c1
	c1 := class{
		Title: "火箭101",
		// 创建 空间为20的空数组
		Students: make([]student, 0, 20),
	}

	// 往班级c1中添加学生
	for i:=0; i < 10; i++ {
		// 构造10个学生
		//fmt.Sprintf("stu%02d", i) 格式化输出得到的字符串
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)
	// 序列化
	// 注意: json序列化对象时，对象的属性必须是首字母大写，否则无法识别
	// 解决: 定制化json模块
	// json序列化： go语言中的数据 -> Json格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed, err: ", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)
	jsonStr := `{"Title":"火箭101","Students":[{"Id":0,"Name":"stu00"},{"Id":1,"Name":"stu01"},{"Id":2,"Name":"stu02"},{"Id":3,"Name":"stu03"},{"Id":4,"Name":"stu04"},{"Id":5,"Name":"stu05"},{"Id":6,"Name":"stu06"},{"Id":7,"Name":"stu07"},{"Id":8,"Name":"stu08"},{"Id":9,"Name":"stu09"}]}`


	// 反序列化
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2) // 修改要用指针
	if err != nil{
		println("json unmarshal failed err:", err)
		return

	}
	fmt.Printf("%#v\n", c2)
}
