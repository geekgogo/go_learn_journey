package main

import "fmt"

func main() {
	// 数组长度无法改变
	fmt.Println("# 🚀1.数组的定义与初始化：第一种方法")
	var n = [3]int{} //定义长度为3的整型数组
	//初始化数组
	n[0] = 1
	n[1] = 2
	n[2] = 10
	// n[3] = 20  错误写法
	fmt.Println(n)
	var str = [2]string{} //定义长度为2的字符串数组
	//初始化数组
	str[0] = "python"
	str[1] = "golang"
	fmt.Println(str)
	fmt.Println("# 🚀2.数组的定义与初始化：第二种方法")
	var a = [3]int{1, 2, 11}
	fmt.Println(a)
	b := [2]string{"php", "java"}
	fmt.Println(b)
	fmt.Println("# 🚀3.数组的定义与初始化：第三种方法，让编译器自动获取数组的长度")
	var c = [...]int{10, 20, 30, 40, 50, 60, 70, 80}
	fmt.Printf("数组：%v,长度：%v,类型：%T\n", c, len(c), c)
	fmt.Println("# 🚀4.数组的定义与初始化：第四种方法，通过指定索引的值来初始化")
	var d = [...]int{0: 1, 1: 2, 5: 50} //没有指定索引和值，int类型数组默认为0
	fmt.Println(d)
	fmt.Println("# 🚀5.遍历数组：for和for range")
	var e = [...]int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(e); i++ {
		fmt.Println(e[i])
	}
	for k, v := range e {
		fmt.Printf("key:%v, value:%v\n", k, v)
	}
	fmt.Println("数组练习")
	fmt.Println("第一题：请求出一组数组里面元素的和与这些元素的平均值，分别用for 和for range实现")
	var f = [...]int{10, 21, 48, 3, 81}
	//使用for实现
	var sum = 0
	for i := 0; i < len(f); i++ {
		sum += f[i]
	}
	var avg = float64(sum) / float64(len(f))
	fmt.Printf("和：%v,平均值：%.2f\n", sum, avg)
	//for range实现
	var sum2 = 0
	for _, v := range f {
		sum2 += v
	}
	fmt.Printf("和：%v,平均值：%.2f\n", sum2, float64(sum2)/float64(len(f)))
	fmt.Println("第二题：求出一个数组的最大值，并取得下标")
	var g = [...]int{10, 21, 48, 3, 81, 6, 101, 8}
	var bid int
	for i := 1; i < len(g); i++ {
		if g[i] > g[0] {
			g[0] = g[i]
			bid = i
		}
	}
	fmt.Printf("最大的值为：%v，下标为：%v\n", g[0], bid)
	fmt.Println("第三题：从数组[1,3,5,7,8]中找出和为8的两个 元素的下标分别为(0,3)和(1,2)")
	var h = [5]int{1, 3, 5, 7, 8}
	for i := 0; i < len(h); i++ {
		for j := i + 1; j < len(h); j++ {
			if h[i]+h[j] == 8 {
				fmt.Printf("(%v,%v)\n", i, j)
			}
		}
	}
}
