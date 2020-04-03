package main

import (
	"fmt"
	"os"
)

// 学生信息管理
// 命令行菜单  fmt.Scanln()
// 添加学生
// 修改学生
// 删除学生
// 展示学生

func showMenu() {
	fmt.Println("学生信息管理系统！")
	fmt.Println("1. 添加学生")
	fmt.Println("2. 修改学生")
	fmt.Println("3. 删除学生")
	fmt.Println("4. 展示学生")
	fmt.Println("5. 退出")
}

func userInput() *Student {
	var (
		name  string
		age   int
		id    int64
		class string
	)
	fmt.Println("请按提示录入信息")
	fmt.Printf("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Printf("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Printf("请输入学号：")
	fmt.Scanln(&id)
	fmt.Printf("请输入班级：")
	fmt.Scanln(&class)
	newStu := NewStudent(name, age, id, class)
	return newStu
}

func main() {
	stuMgr := NewStudentMgr()
	for {
		showMenu()
		// 获取用户输入的指令
		var i int
		fmt.Print("请输入指令：")
		fmt.Scanln(&i)
		fmt.Println("输入的选项是：", i)
		switch i {
		case 1:
			// 添加学生
			newStu := userInput()
			stuMgr.AddStudent(newStu)
		case 2:
			// 修改学生
			newStu := userInput()
			stuMgr.EditStudent(newStu)
		case 3:
			// 删除学生
			newStu := userInput()
			stuMgr.DeleteStudent(newStu)
		case 4:
			// 展示学生
			stuMgr.ShowStudent()
		case 5:
			// 退出
			os.Exit(0)
		default:
			fmt.Println("输入无效")
		}
	}
}
