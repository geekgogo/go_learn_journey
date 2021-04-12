package main

import "fmt"

func main() {
	fmt.Println("# 🚀goto可以根据标签无条件的跳转到指定位置")
	fmt.Println("## 没有goto的情况下：")
	var n = 20
	if n > 18 {
		fmt.Println("成年人")
	}
	fmt.Println("aaa")
	fmt.Println("bbb")
	fmt.Println("ccc")
	fmt.Println("ddd")
	fmt.Println("## 使用goto跳转的情况下：")
	var n1 = 20
	if n1 > 18 {
		fmt.Println("成年人")
		goto label3
	}
	fmt.Println("aaa")
	fmt.Println("bbb")
label3:
	fmt.Println("ccc")
	fmt.Println("ddd")

}
