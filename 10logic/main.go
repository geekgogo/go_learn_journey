package main

import "fmt"

func test() bool {
	fmt.Println("test....")
	return true
}

func main() {
	fmt.Println("# 🚀1.逻辑运算，结果为true或者false")
	/*
		&& 与运算，当左和右的条件都为true，则结果为true，否则为false
		|| 或运算，当左和右的条件有一个为true，则结果为true，否则为false
		! 非运算，取反操作
	*/
	var a = 10
	var b = 3
	fmt.Printf("a > 10 并且 b < 2  -->  %v\n", a > 10 && b < 2)
	fmt.Printf("a >= 10 并且 b < 2  -->  %v\n", a >= 10 && b < 2)
	fmt.Printf("a >= 10 并且 b >= 2  -->  %v\n", a >= 10 && b >= 2)
	fmt.Printf("a > 10 或 b < 2  -->  %v\n", a > 10 || b < 2)
	fmt.Printf("a >= 10 或 b < 2  -->  %v\n", a >= 10 || b < 2)
	fmt.Printf("a >= 10 或 b >= 2  -->  %v\n", a >= 10 || b >= 2)
	fmt.Printf("取反：!true：--> %v\n", !true)
	fmt.Println("# 🚀2.逻辑运算中与运算和或运算的短路")
	/* 短路
	与运算：当前面的一个条件为false时，不会继续执行后面的判断。因为不管后面的条件真假如何都不会影响最终结果为false
	或运算：当前面的一个条件为true时，不会继续执行后面的判断。因为不管后面的条件真假如何都不会影响最终结果为true
	*/
	fmt.Println("# 🚀2.1与运算的短路，未执行到test函数")
	if a > 11 && test() {
		fmt.Println("a大于11")
	}
	fmt.Println("# 🚀2.2与运算的普通状态，会执行到test函数")
	if a < 11 && test() {
		fmt.Println("a小于11")
	}

	fmt.Println("# 🚀2.3或运算的短路，未执行到test函数")
	if a < 11 || test() {
		fmt.Println("a小于11")
	}
	fmt.Println("# 🚀2.4或运算的普通状态，会执行到test函数")
	if a > 11 || test() {
		fmt.Println("a小于11")
	}
}
