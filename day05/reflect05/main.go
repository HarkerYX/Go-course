package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name" hj:"mz"`
	Score int    `json:"score" hj:"fs"`
}

func main() {
	stu1 := Student{
		Name:  "豪杰",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)

	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)
	// 	fmt.Println(field.Tag.Get("hj"))
	// 	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	// }

	// // 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
