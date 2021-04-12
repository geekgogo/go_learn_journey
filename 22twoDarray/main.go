package main

import "fmt"

func main() {
	//数值类型和引用类型
	fmt.Println("# 🚀1.数组和基本数据类型都是数值类型，定义a的值为10，定义b等于a，相当于将a的值复制了一份给b，改变a的值并不会影响到b的值，这就是数值类型")
	var a = 10
	b := a
	a = 20
	fmt.Println(a, b)
	var c = [...]int{10, 20, 30}
	d := c
	c[0] = 11
	fmt.Println(d, c)
	fmt.Println("# 🚀2.切片属于引用类型：定义一个切片e，定义一个切片f等于e，引用类型会将e的内存地址之乡f，当e的内容改变时，f的内容也会随之改变")
	var e = []int{23, 34, 56}
	f := e
	e[0] = 100
	fmt.Println(e, f)
	fmt.Println("# 🚀3.二维数组定义的第一种方法")
	var g = [3][2]int{ //3行2列
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(g)
	fmt.Println("# 🚀3.二维数组定义的第二种方法")
	var h = [...][2]int{ //行可以使用编译器推导，列不能
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println(h)
	//错误的写法
	// var i = [3][...]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// 	{7, 8},
	// }
	fmt.Println("# 🚀4.二维数组取值")
	res := h[0][0] //第一个下边为行，第二个下标为列
	res2 := h[0][1]
	fmt.Println(res)
	fmt.Println(res2)
	fmt.Println("# 🚀5.二维数组的遍历for range")
	for _, v := range h {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
	fmt.Println("# 🚀5.二维数组的遍历for")
	for x := 0; x < len(h); x++ {
		for j := 0; j < len(h[x]); j++ {
			fmt.Println(h[x][j])
		}
	}

}
