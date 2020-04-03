package main

import "fmt"

// Student 学生结构体
type Student struct {
	name  string
	age   int
	id    int64
	class string
}

// NewStudent 是一个创造新学生对象构造函数
func NewStudent(name string, age int, id int64, class string) *Student {
	return &Student{
		name:  name,
		age:   age,
		id:    id,
		class: class,
	}
}

// StudetMgr 定义一个学生信息管理的结构体
type StudetMgr struct {
	AllStudents []*Student
}

// NewStudentMgr 创建新的学生信息管理结构体对象
func NewStudentMgr() *StudetMgr {
	return &StudetMgr{
		AllStudents: make([]*Student, 0, 100),
	}
}

// AddStudent 添加学生的方法
func (s *StudetMgr) AddStudent(stu *Student) {
	for _, v := range s.AllStudents {
		if v.name == stu.name {
			fmt.Printf("姓名为%s的学生已存在\n", stu.name)
			return
		}
	}
	s.AllStudents = append(s.AllStudents, stu)
}

// EditStudent 修改学生的方法
func (s *StudetMgr) EditStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if v.name == stu.name {
			s.AllStudents[index] = stu
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在\n", stu.name)
}

// DeleteStudent 删除学生的方法
func (s *StudetMgr) DeleteStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if v.name == stu.name {
			// 从切片中按照索引删除指定的元素
			s.AllStudents = append(s.AllStudents[:index], s.AllStudents[index+1:]...)
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在\n", stu.name)
}

// ShowStudent 展示学生的方法
func (s *StudetMgr) ShowStudent() {
	for _, v := range s.AllStudents {
		fmt.Printf("name: %s age:%d id:%d class：%s\n", v.name, v.age, v.id, v.class)
	}
}
