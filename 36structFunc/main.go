package main

import "fmt"

type Person struct {
	name   string
	gender string
	age    int
}

// 给Person结构体定义一个方法
func (p Person) PrintInfo() {
	fmt.Println(p.name, p.gender, p.age)
}

// 定义修改结构体的方法，结构体属于值类型，需要使用指针修改
func (p *Person) setInfo(name string, age int) {
	p.name = name
	p.age = age
}

//自定义的数据类型也可以定义方法
type Myint int

func (m Myint) PrintInfo(x int) {
	fmt.Println(x)
}

func main() {
	fmt.Println("🍍 给结构体定义方法")
	/*
		调用结构体的方法
		1.实例化结构体
		2.使用实例调用方法
	*/
	var p1 = Person{
		name:   "李白",
		gender: "男",
		age:    34,
	}
	var p2 = Person{
		name:   "杜牧",
		gender: "男",
		age:    34,
	}
	p1.PrintInfo()
	p1.setInfo("杜甫", 37)
	p1.PrintInfo()
	p2.PrintInfo()
	fmt.Println("🍇 给自定义类型定义方法")
	var m1 Myint
	m1.PrintInfo(3)

}
