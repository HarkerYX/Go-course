package main

import "fmt"

// 链式操作

type student struct {
	name string
}

func (s student) learn() student {
	fmt.Printf("%s热爱学习！\n", s.name)
	return s
}

func (s student) doHomework() student {
	fmt.Printf("%s热爱交作业！\n", s.name)
	return s
}

func main() {
	haojie := student{"豪杰"}
	haojie.learn().doHomework()

	// ret := haojie.learn()
	// ret.doHomework()
}
