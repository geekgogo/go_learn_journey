package main

import (
	"encoding/json"
	"fmt"
)

//Person 定义一个人类结构体
type Person struct {
	Name   string
	Age    int
	Gender string
}

// Animal 给结构体打上json化后字段名的标签
type Animal struct {
	Classify   string `json:"classify"`
	Legs       int    `json:"legs"`
	Eat        bool   `json:"iseat"`
	vertebrate bool   `json:"vertebrate"`
}

// Office 将json转换成字符串首先得定义一个符合要求的结构体
type Office struct {
	Desk     string
	Chair    string
	Computer string
}

// Class 班级包含一个切片类型的结构体
type Class struct {
	Title    string
	Students []Students // 切片类型的Students结构体
}

// Students 学生信息
type Students struct {
	Id   int
	Name string
}

func main() {
	fmt.Println("🍊 结构体和json相互转换")
	fmt.Println("🍋 结构体转json：json.Marshal()")
	//先实例化结构体
	var p = Person{
		Name:   "李白",
		Age:    34,
		Gender: "男",
	}
	jsonByte, _ := json.Marshal(p)
	jsonStr := string([]byte(jsonByte))
	fmt.Printf("值：%v\n", jsonStr)
	fmt.Println("🍌 有时候我们需要将结构体转换成json的字段名称改为小写或者其他名称，这个时候我们就需要用到结构体的标签")
	var a = Animal{
		Classify:   "water",
		Legs:       4,
		Eat:        true,
		vertebrate: true,
	}
	aniByte, _ := json.Marshal(a)
	aniStr := string([]byte(aniByte))
	fmt.Println(aniStr)
	fmt.Println("🍉 json不能读取结构体中的私有字段")
	fmt.Println("🥭 json转换成结构体")
	//定义一个json格式的字符串
	var office = `{"Desk": "ikea","Chair": "ikea","Computer": "apple"}`
	//实例化结构体
	var o Office
	// Unmarshal两个参数：1.字节类型的json字符串 2.结构体的地址
	err := json.Unmarshal([]byte(office), &o)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", o)
	}
	fmt.Println("🍍 嵌套结构体和json的转换")
	//嵌套结构体转json
	var sli = make([]Students, 0)
	for i := 0; i <= 6; i++ {
		s := Students{
			Id:   i,
			Name: fmt.Sprintf("stu_%v", i),
		}
		sli = append(sli, s)
	}
	c := Class{
		Title:    "英语班",
		Students: sli,
	}
	classByte, _ := json.Marshal(c)
	classJSON := string(classByte)
	fmt.Println(classJSON)
	//json转嵌套结构体
	//定义json格式的字符串
	var c2j = `{"Title": "数学班", "Students":[{"Id":1, "Name":"李白"},{"Id":2, "Name":"杜甫"}]}`
	c2 := &Class{} //以指针类型实例化结构体，这里就代表了结构体的地址，下面使用Unmarshal时就无需带 & 符号
	err1 := json.Unmarshal([]byte(c2j), c2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("%#v\n", c2)
		fmt.Printf("%v\n", c2.Title)

	}

}
