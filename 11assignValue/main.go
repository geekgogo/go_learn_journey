package main

import "fmt"

func main() {
	/*
		赋值运算符：
		= 赋值
		+= 相加后再赋值
		—= 相减后再赋值
		*= 相乘后再赋值
		/= 相除后再赋值
		%= 求余后再赋值
	*/
	fmt.Println("# 🚀1. 赋值")
	var a int = 10
	var b = a
	fmt.Printf("将a的值赋值给b，b=%v\n", b)
	fmt.Println("# 🚀2. 相加后再赋值")
	a += 10
	fmt.Printf("a += 10的值为：%v\n", a)
	fmt.Println("# 🚀3. 相减后再赋值")
	a -= 2
	fmt.Printf("a -= 2的值为：%v\n", a)
	fmt.Println("# 🚀4. 相乘后再赋值")
	a *= 2
	fmt.Printf("a *= 2的值为：%v\n", a)
	fmt.Println("# 🚀5. 相除后再赋值")
	a /= 2
	fmt.Printf("a /= 2的值为：%v\n", a)
	fmt.Println("# 🚀6. 求余后再赋值")
	a %= 2
	fmt.Printf("a对2求余后再赋值的值为：%v\n", a)

}
