package main

import "fmt"

// Person 匿名结构体
type Person struct {
	string
	int
}

/*结构体的嵌套
1. 定义两个结构体分别为 User和Address
2. 在User结构体中嵌套Address
*/

// User 定义一个用户结构体
type User struct {
	Username string
	Password string
	// Address  Address //这里表示定义一个Address的变量，类型是下面的Address结构体，也可以直接写成结构体名称，不需要变量名
	Address
	Email
}

// Address 定义一个地址结构体
type Address struct {
	City    string
	Phone   string
	Name    string
	AddTime string
}

// Email
type Email struct {
	AddTime string
}

//结构体的继承

// Animal 父结构体
type Animal struct {
	Name string
}

// 父结构体的方法
func (a Animal) run() {
	fmt.Printf("%v正在跑。。。\n", a.Name)
}

// Dog 子结构体
type Dog struct {
	Name string
	Animal
}

func main() {
	fmt.Println("🧀 struct嵌套和继承")
	fmt.Println("🍏 匿名结构体，不定义变量名称只写类型，但不能定义相同类型，实例化时赋予的变量必须顺序保持一致。匿名结构体默认使用类型名作为字段名")
	p := Person{
		"李白",
		34,
	}
	fmt.Printf("%#v\n", p)
	fmt.Println("🍐 结构体的嵌套")
	var u User //实例化结构体
	u.Username = "dufu"
	u.Password = "123456"
	u.Address.City = "山东" //给结构体中的结构体字段赋值
	u.Address.Phone = "12345678901"
	// 给嵌套结构体赋值时可以省略结构体的名称，golang会先从自身结构体寻找字段，找不到则会向嵌套的结构体寻找
	u.Name = "杜甫"
	fmt.Printf("%#v\n", u)
	fmt.Println("⚠️ 一个结构体可以嵌套多个结构体，若多个结构体里面的字段重复时，赋值必须指定结构体的名称")
	// u.AddTime = "2020-05-20" //报错：ambiguous selector u.AddTime
	u.Address.AddTime = "2020-05-20"
	u.Email.AddTime = "2020-05-21"
	fmt.Printf("%#v\n", u)
	fmt.Println("🍒 结构体的继承")
	var d Dog
	d.Animal.Name = "柴犬"
	d.run() //调用父结构体的方法，也就是子结构体继承了父结构体的方法

}
