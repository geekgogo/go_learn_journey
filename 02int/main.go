package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 //定义a为int8整型
	a = 15
	fmt.Printf("%v, 类型是：%T\n", a, a)
	fmt.Println(unsafe.Sizeof(a)) //打印变量a在内存所占的字节数

	//ouput
	// 15, 类型是：int8
	// 1
	//int8整型占1个字节，也就是8位

	var b int32
	b = 12
	fmt.Printf("%v, 类型是：%T\n", b, b)
	fmt.Println(unsafe.Sizeof(b))
	//output
	// 12, 类型是：int32
	// 4 占4个字节

	//虽然b比a小，但是在内存占用上却多得多，编写代码时需要选择合适的数据类型

	//无符号型与有符号型所占用的空间一样
	var c uint8
	c = 254
	fmt.Printf("%v, type: %T\n", c, c)
	fmt.Println("占用空间：", unsafe.Sizeof(c))
	//output
	// 254, type: uint8
	// 占用空间： 1

	//数据类型的转换，需要注意的是高位向低位转换时的溢出问题
	var d int16
	d = 155
	fmt.Println(d)
	fmt.Println(int8(d)) //这里d的长度大于int8类型的长度，所以转换会无法得到预期的效果

	//不同类型的整型不能直接做数学运算
	var e int8
	var f int16
	e = 12
	f = 31
	// fmt.Println(e + f)
	//output
	//报错 mismatched types int8 and int16
	//使用类型转换相加
	fmt.Println(e + int8(f))

	//数字字面量语法
	var g int8
	g = 12

	fmt.Printf("%v\n", g) //原样输出
	fmt.Printf("%d\n", g) //十进制输出
	fmt.Printf("%b\n", g) //二进制输出
	fmt.Printf("%o\n", g) //八进制输出
	fmt.Printf("%x\n", g) //十六进制输出

	//小知识点
	//在64位操作系统中，int类型就是int64，占用8个字节，32位中int类型是int32，占用4个字节
	var h int
	fmt.Println("小知识点：")
	fmt.Printf("%v, %T\n", h, h)
	fmt.Println(unsafe.Sizeof(h))

}
