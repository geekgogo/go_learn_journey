package main

import "fmt"

//type的用法
//定义自定义类型
type myInt int

//定义类型的别名
type myFloat = float64

// Person 定义结构体，结构体名称可以大写（表示公有结构体），也可以小写（表示私有结构体）
type Person struct {
	name   string
	gender string
	age    int
}

func main() {
	fmt.Println("🍌 type自定义类型myInt，float64的别名myFloat")
	var a myInt = 64
	var b myFloat = 13.4
	fmt.Printf("myInt类型的值：%v，类型：%T，myFloat的值：%v，类型：%T\n", a, a, b, b)
	fmt.Println("🍎 struct结构体，类似于其他面向对象语言的类，下面给出七种实例化结构体的方法")
	//实例化结构体的第一种方法（类型实例）
	var p1 Person
	p1.name = "李白"
	p1.gender = "男"
	p1.age = 34
	fmt.Printf("值：%#v，类型：%T\n", p1, p1)
	//实例化结构体的第二种方法（指针类型）
	var p2 = new(Person)
	p2.name = "杜甫"
	p2.gender = "男"
	p2.age = 35
	fmt.Printf("值：%#v，类型：%T\n", p2, p2)
	//实例结构化的第三种方法（指针类型）
	p3 := &Person{}
	p3.name = "李商隐"
	p3.gender = "男"
	p3.age = 36
	fmt.Printf("值：%#v，类型：%T\n", p3, p3)
	//实例结构化的第四种方法（指针类型）
	p4 := &Person{
		name:   "杜牧",
		gender: "男",
		age:    37, //最后一个逗号必须填写
	}
	fmt.Printf("值：%#v，类型：%T\n", p4, p4)
	//实例结构化的第五种方法（指针类型）
	p5 := &Person{
		//需要与定义结构体时顺序保持一致
		"李清照",
		"女",
		28, //最后一个逗号必须填写
	}
	fmt.Printf("值：%#v，类型：%T\n", p5, p5)
	//实例结构化的第六种方法（类型实例）
	p6 := Person{
		"王勃",
		"男",
		16, //最后一个逗号必须填写
	}
	fmt.Printf("值：%#v，类型：%T\n", p6, p6)
	//实例结构化的第七种方法（类型实例）
	p7 := Person{
		name:   "贺知章",
		gender: "男",
		age:    39, //最后一个逗号必须填写
	}
	fmt.Printf("值：%#v，类型：%T\n", p7, p7)
}
