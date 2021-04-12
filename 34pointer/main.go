package main

import "fmt"

func main() {
	//指针：存放一段内存地址，属于指针类型
	var a int = 10
	p := &a //定义一个指针类型的变量p存储a的内存地址
	v := *p // 定义一个变量 v存储p指针对应变量的值
	fmt.Printf("变量a的值：%v，a的类型：%v，a的地址：%p，根据a的内存地址取出值：%v\n", a, a, p, v)
	fmt.Printf("指针p的类型：%T\n", p)
	*p = 30 //利用指针改变a的值
	fmt.Printf("通过 *p=30 改变a的值后：a=%v\n", a)
	fmt.Println("🍈 定一个空指针")
	var p1 *int
	fmt.Printf("%v\n", p1)
	fmt.Println("🍑 指针必须设置内存地址才能使用，因为指针和切片、map类型数据一样属于引用类型数据。指针用new函数分配内存（实际项目中很少用到）")
	p2 := new(int) //定义一个int型指针类型并分配空间
	var p3 *bool
	p3 = new(bool) //定义一个bool型的指针类型并分配空间
	*p2 = 11
	*p3 = true
	fmt.Printf("p2的类型：%T，p2的值：%v，p3的类型：%T，p3的值：%v\n", p2, *p2, p3, *p3)
}
