package main

import "fmt"

func main() {
	fmt.Println("# 🚀1. bool类型的默认值是false")
	var a bool
	fmt.Printf("a的值是：%v, a的类型是：%T\n", a, a)
	fmt.Println("# 🚀2. bool类型不能与其他数据类型进行转换，比如python语言中1为true，0为false")
	var b bool = true
	// var c = 0
	// 正确写法
	if b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	//错误写法
	// if c {
	// 	fmt.Println("true")
	// }
	fmt.Println("# 🚀3. go语言中不能将整型强制转换称bool型")
	//报错
	// var d = 0
	// var e = bool(d)

}
