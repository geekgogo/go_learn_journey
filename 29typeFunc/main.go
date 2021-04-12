package main

import (
	"fmt"
)

//使用type语句自定义一个int类型
type myInt int

//使用type语句自定义一个string类型
type myString string

// 使用type语句自定义一个函数类型
type calc func(int, int) int

func add(x, y int) int {
	return x + y
}

//函数接收函数作为参数
func sub(x, y int, f calc) int {
	return f(x, y)
}
func sub2(x, y int) int {
	return x - y
}

//函数的返回值为函数
//例1:
func sub3(x, y int) func(int, int) int {
	return func(a int, b int) int {
		return x + y + a + b
	}
}

//例2:
type calcType func(int, int) int

func do(o string) calcType {
	switch o {
	case "+":
		return add
	case "-":
		return sub2
	case "*":
		//匿名函数
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}
func main() {
	a := add
	var a1 calc
	a1 = add
	fmt.Println("🐹 使用type自定义一个函数类型")
	fmt.Printf("将函数赋值利用推导式赋值给变量，此时它的类型是：%T，将自定义函数赋值给使用var定义的变量，此时它的类型是：%T\n", a, a1)

	fmt.Println("🐫 使用type自定义的函数作为参数传入另一个参数")
	c := sub(3, 1, sub2)
	fmt.Println(c)
	f := sub(4, 5, func(x, y int) int { return x * y })
	fmt.Printf("😎 使用匿名函数当做参数：%v\n", f)

	fmt.Println("🐵 使用type自定义一个myInt类型")
	var d myInt = 10
	var d1 int = 11
	fmt.Printf("自定义的类型也可以用来转换：d + myInt(d1) = %v\n", d+myInt(d1))
	fmt.Printf("%v--%T\n", d, d)

	fmt.Println("🐍 使用type自定义一个myString类型")
	var e myString = "hello golang"
	fmt.Printf("%v--%T\n", e, e)

	fmt.Println("🐸 匿名函数的使用")
	//没有返回值和参数的自执行匿名函数
	func() {
		fmt.Println("这里是匿名函数")
	}()
	// 将匿名函数复制给变量
	l := func(x, y int) int {
		return x * y
	}
	fmt.Println(l(10, 10))
	//带参数和返回的自执行匿名函数
	fmt.Println(func(x, y int) int { return x / y }(9, 3))

	fmt.Println("🐶 将函数作为函数的返回值， 例1")
	g := sub3(3, 4)
	fmt.Printf("%v--%T\n", g(5, 6), g(5, 6))
	fmt.Println("🐶 将函数作为函数的返回值，例2")
	h := do("+")
	i := do("-")
	j := do("*")
	fmt.Println(h(3, 4), i(5, 4), j(5, 4))
}
