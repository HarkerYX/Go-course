package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// 解析日志库的配置文件

// Config 是一个日志配置项, 字段名大写
type Config struct {
	FilePath string `conf:"file_path" db:"name"`
	FileName string `conf:"file_name"`
	MaxSize  int64  `conf:"max_size"`
}

// 从conf文件中读取内容赋值给结构体指针
func parseConf(fileName string, result interface{}) (err error) {
	// 0. 前提条件，result必须是一个ptr
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)

	if t.Kind() != reflect.Ptr {
		err = errors.New("result必须是一个指针")
		return
	}
	// result不是结构体也不行
	tElem := t.Elem()
	if tElem.Kind() != reflect.Struct {
		err = errors.New("result必须是一个结构体指针")
		return
	}

	// 1. 打开文件
	data, err := ioutil.ReadFile(fileName) // data:[]byte
	if err != nil {
		// err = errors.New("打开配置文件失败")
		err = fmt.Errorf("打开配置文件%s失败", fileName)
		return err
	}
	// 2. 将读取的文件数据按照行分割，得到一个行的切片
	lineSlice := strings.Split(string(data), "\r\n")
	fmt.Println(lineSlice)
	// 一行一行的解析
	for index, line := range lineSlice {
		line = strings.TrimSpace(line) //去除字符串首尾的空白
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			// 忽略空行和注释
			continue
		}
		// 解析是不是正经的配置项(判断是不是有等号)
		equalIndex := strings.Index(line, "=")
		if equalIndex == -1 {
			err = fmt.Errorf("第%d行语法错误", index+1)
			return
		}
		// 按照=分割每一行，左边是key 右边是value
		key := line[:equalIndex]
		value := line[equalIndex+1:]
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if len(key) == 0 {
			err = fmt.Errorf("第%d行语法错误", index+1)
			return
		}
		//利用反射 给 result 赋值
		// 遍历结构体的每一个字段和key比较，匹配上了就把value赋值
		for i := 0; i < tElem.NumField(); i++ {
			field := tElem.Field(i)
			tag := field.Tag.Get("conf")
			if key == tag {
				// 匹配上了就把value赋值
				// 拿到每个字段的类型
				fieldType := field.Type
				switch fieldType.Kind() {
				case reflect.String:
					fieldValue := v.Elem().FieldByName(field.Name) // 根据字段名找到对应的字段值
					fieldValue.SetString(value)                    // 将配置文件中读取的value设置给结构体
				case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
					value64, _ := strconv.ParseInt(value, 10, 64) // strconv包去做进制转换
					v.Elem().Field(i).SetInt(value64)
				}
			}
		}

	}
	return
}

func main() {

	// 2. 读取内容
	// 3. 一行一行读取内容 根据tag找结构体里面的对应字段
	// 4. 找到了要赋值

	var c = &Config{} // 用来存储读取的数据
	err := parseConf("xxx.conf", c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", c) // 解析之后的数据

}
