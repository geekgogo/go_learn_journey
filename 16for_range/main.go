package main

import "fmt"

func main() {
	var a = "你好golang"
	fmt.Println("# 🚀1.循环字符串：将字符串以字符的形式逐个输出")
	for k, v := range a {
		fmt.Printf("k:%v,v:%c\n", k, v)
	}
	fmt.Println("# 🚀2.循环切片")
	var b = []string{"php", "golang", "python"}
	for _, v := range b {
		fmt.Println(v)
	}

}
