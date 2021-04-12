package main

import "fmt"

//递归：函数调用自身
//例1：定义一个函数，参数为整型，打印出1到这个整型之间所有的数字
func f1(n int) int {
	if n > 1 {
		fmt.Println(n)
		n--
		return f1(n)
	} else {
		return 1
	}
}

//例2：定义一个函数，参数为整型，算出这个整型的阶乘
func f2(n int) int {
	if n > 1 {
		return n * f2(n-1)
	} else {
		return 1
	}
}

//例3:递归实现1到100的和
func f3(n int) int {
	if n > 1 {
		return n + f3(n-1)
	}
	return 1
}

//闭包：在函数A内定义另一个函数B，B可以调用函数A的变量，不建议经常使用
/*
全局变量的特点：常驻内存，污染全局
局部变量的特点：不常驻内存，不污染全局
闭包的特点：可以让一个变量常驻内存，且不污染全局
*/

func A() func(x int) int {
	var b int = 10
	return func(x int) int {
		b += x
		return b
	}
}
func main() {
	fmt.Println(f1(10))
	fmt.Println(f2(5))
	fmt.Printf("递归实现1到100的和：%v\n", f3(100))
	g := A()           //将函数定义给变量g
	fmt.Println(g(10)) //返回10
	fmt.Println(g(10)) //返回20，由于上一个操作修改了b的值，这里再次调用闭包时b已经被修改。
}
