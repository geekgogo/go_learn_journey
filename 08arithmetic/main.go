package main

import "fmt"

func main() {
	/*
		+ 加
		- 减
		* 乘
		/ 除
		% 取余的规则是：被除数 - (被除数 / 除数) * 除数

	*/
	var a = 10
	var b = 3
	fmt.Println("# 🚀1. 整数间的运算")
	fmt.Printf("10 + 3 = %v\n", a+b)
	fmt.Printf("10 - 3 = %v\n", a-b)
	fmt.Printf("10 * 3 = %v\n", a*b)
	fmt.Printf("10 / 3 = %v\n", a/b)
	fmt.Printf("10对3取余 = %v\n", a%b)
	fmt.Println("# 🚀2. 浮点间的运算")
	var c float64 = 10
	var d float64 = 3
	fmt.Printf("10 + 3 = %v\n", c+d)
	fmt.Printf("10 - 3 = %v\n", c-d)
	fmt.Printf("10 * 3 = %v\n", c*d)
	fmt.Printf("10 / 3 = %v\n", c/d) //当被除数或者除数有一个为浮点类型，则值为浮点类型
	fmt.Println("# 🚀3. 负数间的运算")
	var e = -10
	var f = 3
	fmt.Printf("-10 + 3 = %v\n", e+f)
	fmt.Printf("-10 - 3 = %v\n", e-f)
	fmt.Printf("-10 * 3 = %v\n", e*f)
	fmt.Printf("-10对3取余 = %v\n", e%f)   //遵循规则：余 =  被除数 - (被除数 / 除数) * 除数
	fmt.Printf("10对-3取余 = %v\n", 10%-3) //遵循规则：余 =  被除数 - (被除数 / 除数) * 除数

	/*
		i ++ 自增
		i -- 自减
	*/
	fmt.Println("# 🚀 4.golang中的自增自减，必须放在变量后面，而且只能作为单独的语句，不能进行赋值操作")
	var g int64 = 10
	var h int64 = 10
	g++
	h--
	fmt.Printf("10的自增结果为：%v\n", g)
	fmt.Printf("10的自减结果为：%v\n", h)
	// i := g++  //错误写法，不能进行赋值，只能作为单独的语句
	// ++g //错误写法，只能放在变量的后面
	// --h //错误写法，只能放在变量的后面
}
