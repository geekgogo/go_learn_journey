package main

import "fmt"

func main() {
	/*
		if语句
		if 条件判断 {
			执行的语句
		}
		tips：1.条件判断可以用（）括起来，但不建议使用
		2. {}的第一个{ 必须放在条件判断的后面，不能换行，否则会报错
	*/
	var flag = true
	if flag {
		fmt.Println("flag=true")
	}
	/*
		if 语句的第一种写法
	*/
	var score = 92
	if score > 90 { // 当前区域内的全局变量
		fmt.Println("优秀")
	} else if score > 80 {
		fmt.Println("中等")
	} else {
		fmt.Println("差劲")
	}
	fmt.Println(score)
	/*
		if 语句的第二种写法
	*/
	if a := 32; a > 90 { // 此时的a是if语句中的局部变量，无法在if语句外获取
		fmt.Println("优秀")
	} else {
		fmt.Println("差劲")
	}
	// 尝试在if语句外获取，报错
	// fmt.Println(a)
}
