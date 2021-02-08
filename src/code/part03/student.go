package main

import "fmt"

type student struct {
	id int // 学员是唯一的
	name string
	class string
}

// newStudent 是student类型的构造函数
func newStudent(id int, name, class string) *student{
	return &student{
		id: id,
		name: name,
		class:class,
	}
}

// 学员管理的类型
type studentMgr struct{
	allStudents []*student
}

// 学员管理构造函数
func newStudentMgr() *studentMgr{
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 1.添加学生
func (s *studentMgr) addStudent(newStu *student){
	s.allStudents = append(s.allStudents, newStu)
}

// 2.编辑学生
func (s *studentMgr) modifyStudent(newStu *student){
	for i, v := range s.allStudents{
		// 当学号相同时，就表示找到了要修改的学生
		if newStu.id == v.id{
			// 根据切片的索引直接把心学生赋值进来
			s.allStudents[i] = newStu
			return
		}
	}
	fmt.Printf("输入的学生有误，系统中没有学号是: %d 的学生\n", newStu.id)
}

// 3.展示学生
func (s *studentMgr) showStudent(newStu *student){
	for _,v := range s.allStudents{
		fmt.Printf("学号: %d  姓名: %s  班级: %s", v.id, v.name, v.class)
	}
}