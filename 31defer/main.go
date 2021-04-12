package main

import "fmt"

/*
defer: 延迟执行，若定义了多个defer，则按定义的顺序逆序执行
*/
//defer在命名返回值函数和匿名函数中的区别
//匿名函数，结果为0
func f1() int {
	var a int      // 0
	defer func() { // 延迟调用
		a++
	}()
	return a // 先return a,再调用上面的defer语句
}

//命名返回值函数，结果为1
func f2() (a int) {
	defer func() {
		a++
	}()
	return a
}

/*
f1和f2的结果不同的原因在于golang中的return语句的操作顺序
函数中return语句的底层实现：
return x
	第一步： 返回值 = x
	第二步： ret指定
defer语句执行的时机：
return x
	第一步：返回值 = x
	第二步： 运行defer
	第三步：ret指令

*/
//案例3:

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	// fmt.Println("🐸 正常执行顺序")
	// fmt.Println("start...")
	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println("end")
	// fmt.Println("🐦 defer执行顺序")
	// fmt.Println("start...")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("end")
	fmt.Println("🦢 匿名函数的中使用defer")
	fmt.Println(f1())
	fmt.Println("🦢 命名返回值函数的中使用defer")
	fmt.Println(f2())
	fmt.Println("🦄️ defer函数的案例3：")
	/*执行顺序：defer先注册，后返回。注册时
	注册时：
	第一步：calc("AA", x, calc("A", x, y))
	第二步：calc("BB", x, calc("B", x, y))
	执行时：延迟执行
	第一步：calc("BB", x, calc("B", x, y))
	第二步：calc("AA", x, calc("A", x, y))
	所以结果：
	A 1 2 3
	B 10 2 12
	BB 10 12 22
	AA 1 3 4
	*/
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

}
