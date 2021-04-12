package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() { //值类型方法
	fmt.Printf("%v 开机\n", p.Name)
}
func (p Phone) stop() { //值类型方法
	fmt.Printf("%v 关机\n", p.Name)
}

type Camera struct {
	Name string
}

func (c *Camera) start() { //引用类型方法
	fmt.Printf("%v 开机\n", c.Name)
}
func (c *Camera) stop() { //引用类型方法
	fmt.Printf("%v 开机\n", c.Name)
}

// 一个结构体实现多个接口
type Animalera interface {
	GetAge() int
}

type Animalerb interface {
	SetAge(int)
}

type Cat struct {
	Age int
}

func (c *Cat) SetAge(age int) {
	c.Age = age
}

func (c Cat) GetAge() int {
	return c.Age
}

// 接口嵌套
type Ainterface interface {
	SetName(string)
}
type Binterface interface {
	GetName() string
}
type Animaler interface {
	Ainterface
	Binterface
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name

}

func (d Dog) GetName() string {
	return d.Name
}
func main() {
	fmt.Println("😀 结构体的值类型接收函数与指针类型接收函数实现接口")
	// 用值类型方法定义的结构体，使用值类型和引用类型实例化结构体都可以实现接口
	var p1 = Phone{
		Name: "小米",
	}
	var u1 Usber = p1
	u1.start()
	u1.stop()
	var p2 = &Phone{
		Name: "苹果",
	}
	var u2 Usber = p2
	u2.start()
	u2.stop()
	// 用引用类型方法的结构体，只能使用引用类型实例化才能实现接口
	var c = &Camera{
		Name: "fujitsu",
	}
	var u3 Usber = c
	u3.start()
	u3.stop()
	/*错误写法
	var c1 = Camera{
		Name: "sony",
	}
	var u4 Usber = c1
	c1.start()
	c1.stop()
	*/
	fmt.Println("😃 一个结构体实现多个接口")
	var ca = &Cat{
		Age: 5,
	}
	var aa Animalera = ca //让Cat实现Animalera接口
	var ab Animalerb = ca //让Cat实现Animalerb接口
	fmt.Println(aa.GetAge())
	ab.SetAge(3)
	fmt.Println(aa.GetAge())
	fmt.Println("😄 接口嵌套")
	var d = &Dog{
		Name: "阿柴",
	}
	var a Animaler = d
	name1 := a.GetName()
	fmt.Println(name1)
	a.SetName("黑子")
	name2 := a.GetName()
	fmt.Println(name2)
	fmt.Println("由此可见，接口是允许嵌套使用的，并且使用方法与正常的接口并无区别")
}
