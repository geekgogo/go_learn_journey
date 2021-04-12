package main

import "fmt"

func main() {
	fmt.Println("# 🚀1. 条件判断，结果为true或这false")
	/*
		>
		<
		>=
		<=
		!=
	*/
	var a = 10
	var b = 3
	fmt.Printf("a大于b：%v\n", a > b)
	fmt.Printf("a小于b：%v\n", a < b)
	fmt.Printf("a大于等于b：%v\n", a >= b)
	fmt.Printf("a小于等于b：%v\n", a <= b)
	fmt.Printf("a不等于b：%v\n", a != b)

}
