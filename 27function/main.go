package main

import "fmt"

//语法：
// func 函数名(参数)返回类型 {
// 	函数体
// }

func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

//当参数类型相同的可以省略前面的变量类型
func sumFn2(x, y int) int {
	sum := x - y
	return sum
}

//多个返回值的函数
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//接收可变参数，以 ... 表示动态参数（不确定参数个数的情况下）
func sumFn3(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

// 函数同时接收普通参数和可变参数，可变参数必须放在固定参数后面
func sumFn4(x int, y ...int) int {
	for _, v := range y {
		x += v
	}
	return x
}

// 直接定义返回值的名称，直接在定义函数时给返回值命名。同样跟相同类型的参数一样，可以将返回值的类型省略
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}
func main() {
	s := sumFn(3, 2)
	s1 := sumFn2(9, 2)
	s2, s3 := calc(10, 8)
	s4 := sumFn3(1, 2, 3, 4, 5, 5)
	s5 := sumFn4(1, 2, 3)
	s6Sum, s6Sub := calc2(7, 7)
	fmt.Printf("s: %v, s1:%v, s2:%v, s3:%v, s4:%v, s5:%v, s6Sum:%v, s6Sub:%v\n", s, s1, s2, s3, s4, s5, s6Sum, s6Sub)
}
