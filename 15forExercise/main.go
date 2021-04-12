package main

import "fmt"

func main() {
	fmt.Println("# 🚀1.打印0-50所有的偶数")
	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("# 🚀2.打印1+2+3+..+100的和")
	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("# 🚀3.打印1-100之间所有9的倍数的和以及个数")
	var sum2 = 0
	var count = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			sum2 += i
			count++
		}
	}
	fmt.Printf("和：%v，个数：%v\n", sum2, count)
	fmt.Println("# 🚀4.计算5的阶乘")
	var sum3 = 1
	for i := 1; i <= 5; i++ {
		sum3 *= i
	}
	fmt.Println(sum3)
	fmt.Println("# 🚀5.打印一个矩形")
	/*
	******
	******
	******
	 */
	var row = 3
	var column = 6
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Print("*") //使用Print避免换行
		}
		fmt.Println("") //使用Println打印空字符串，使其循环完一次row换行
	}
	fmt.Println("# 🚀6.打印一个三角形")
	for i := 0; i < 6; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("# 🚀7.打印一个99乘法表")
	var row2 = 9
	for i := 1; i <= row2; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%vx%v=%v\t", j, i, j*i)
		}
		fmt.Println("")
	}

}
