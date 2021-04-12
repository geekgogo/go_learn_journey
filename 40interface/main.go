package main

import "fmt"

//定义一个接口:接口中可以包含方法且方法没有方法体，并且不能包含变量
type Usber interface {
	start()
	stop()
}

//Phone 接口的实现只能使用结构体或者自定义方法，一般使用结构体
type Phone struct{}

//使用结构体实现接口的要求是：必须实现接口中的所有方法
func (p Phone) start() {
	fmt.Println("开机")
}
func (p Phone) stop() {
	fmt.Println("关机")
}

// 接口 例2:

// Computer 结构体
type Computer struct{}

// Computer的函数接收 usber 接口的作为参数
func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

// Ssd 定义Ssd结构体用来实现使用usber接口连接Computer结构体
type Ssd struct {
	Name string
}

// 实现usber接口必须完成定义接口中的所有方法
func (s Ssd) start() {
	fmt.Printf("%v开始连接电脑\n", s.Name)
}
func (s Ssd) stop() {
	fmt.Printf("%v停止连接电脑\n", s.Name)
}

func main() {
	fmt.Println("🥑 接口表示约束和规范")
	fmt.Println("🍆 第一种使用接口的方法")
	var per Usber   //定义接口类型的变量
	var p = Phone{} //结构体和接口相互关联
	per = p
	per.start()
	fmt.Println("🍆 第二种使用接口的方法")
	var p2 Usber = Phone{}
	p2.stop()
	fmt.Println("🥦 接口例2：")
	var s Usber = Ssd{
		Name: "TOSHIBA硬盘",
	}
	s.start()
	s.stop()
}
