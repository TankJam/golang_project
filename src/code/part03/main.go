package main

import (
	"fmt"
	"os"
)

// 学院信息管理系统
func showMenu(){
	fmt.Println("欢迎来到学员管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")
}

// 获取用户输入的学员信息
func getInput() *student{
	var (
		id int
		name string
		class string
	)
	fmt.Println("请按要求输入学员信息")
	fmt.Println("请输入学员的学号: ")
	fmt.Scanf("%d\n", id)
	fmt.Println("请输入学员的姓名: ")
	fmt.Scanf("%s\n", name)
	fmt.Println("请输入学员的班级: ")
	fmt.Scanf("%s\n", class)
	newStu := newStudent(id, name, class)
	return newStu
}

func main(){
	sm := newStudentMgr()
	// 1.打印系统菜单
	showMenu()
	// 2.等待用户选择要执行的选项
	var input int
	fmt.Scanf("%d\n", &input)
	fmt.Println("用户输入的是: ", input)
	// 3.执行用户选择的动作
	switch input{
	case 1:
		// 添加学员
		sm.addStudent()
	case 2:
		// 编辑学员
		sm.modifyStudent()
	case 3:
		// 展示所有学员
		sm.showStudent()
	case 4:
		// 退出
		os.Exit(0)
	}
}

