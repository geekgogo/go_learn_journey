package main

import (
	"fmt"
	"unsafe"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("# 🚀1.golang语言中浮点类型有float32和float64")
	var a float32
	var b float64
	a = 3.1314
	b = 56.243432
	fmt.Printf("%v, type: %T\n", a, a)
	fmt.Printf("%v, type: %T\n", b, b)
	fmt.Printf("a的占用空间:%v个字节, b的占用空间：%v个字节\n", unsafe.Sizeof(a), unsafe.Sizeof(b))

	fmt.Println("# 🚀2. 浮点的运算会导致精度的丢失")
	var c float64
	var d float64
	c = 8.3
	d = 3.8
	fmt.Printf("8.3 - 3.8 = %v\n", c-d)
	fmt.Printf("8.3 + 3.8 = %v\n", c+d)
	fmt.Println("## 2.1 需要引入第三方包解决精度丢失的问题：github.com/shopspring/decimal")
	fmt.Printf("8.3 - 3.8 = %v\n", decimal.NewFromFloat(c).Sub(decimal.NewFromFloat(d)))
	fmt.Printf("8.3 + 3.8 = %v\n", decimal.NewFromFloat(c).Add(decimal.NewFromFloat(d)))

	fmt.Println("# 🚀3. float类型相互转换以及float类型和int类型的相互转换")
	var e float32 = 32.4
	var f = float64(e)
	fmt.Println("## 3.1 float32转float64")
	fmt.Printf("f的值是：%v, f的类型是：%T\n", f, f)
	fmt.Println("## 3.2 float64转float32")
	var g = float32(f)
	fmt.Printf("g的值是：%v， g的类型是：%T\n", g, g)
	fmt.Println("## 3.3 float与int互转")
	var h int8 = 4
	var i float64 = 38.885
	fmt.Printf("h的值是：%v， h初始的类型是：%T，h转float的值是：%v， h转float的类型是：%T\n", h, h, float64(h), float64(h))
	fmt.Printf("i的值是：%v， i初始的类型是：%T，i转int的值是：%v， i转int的类型是：%T\n", i, i, int(i), int(i))

}
