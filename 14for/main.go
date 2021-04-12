package main

import "fmt"

func main() {
	/*
		for 循环
		for 初始条件;条件表达式;结束语句 {
			循环体
		}
		golang中没有while语句
		for 循环可以通过break,goto,return,panic语句强制退出循环
	*/

	fmt.Println("# 🍞打印1到10")
	fmt.Println("# 🚀1.for循环的第一种写法")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("# 🚀2.for循环的第二种写法")
	var i = 1
	for ; i <= 10; i++ { //初始条件后的封号（;）不能省略
		fmt.Println(i)
	}
	fmt.Println("# 🚀3.for循环的第三种写法")
	var a = 1
	for a <= 10 {
		fmt.Println(a)
		a++
	}
	fmt.Println("# 🚀4.for循环的第四种写法")
	var b = 1
	for {
		if b <= 10 {
			fmt.Println(b)
		} else {
			break
		}
		b++
	}

}
