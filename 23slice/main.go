package main

import "fmt"

func main() {
	fmt.Println("# 🚀1. 切片的定义与初始化")
	var a []int //切片的底层是数组，定义时跟数组类似，不指定长度时就是切片。
	fmt.Println(a)
	fmt.Println(a == nil)  //定义时没有指定值，默认值为nil
	var b = []int{1, 2, 3} //定义切片并初始化
	fmt.Println(b)
	var c = []int{1: 2, 2: 3, 4: 6} //根据下标初始化切片
	fmt.Println(c)
	fmt.Println("# 🚀2. 切片的遍历")
	var d = []string{"python", "golang", "java", "c#"}
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}
	//for range用法
	for _, v := range d {
		fmt.Println(v)
	}
	fmt.Println("# 🚀3. 基于数组的切片")
	var e = [...]int{10, 13, 46, 93, 20, 78}
	f := e[3:]  //取下标为3（包括3）往后所有的元素组成切片
	g := e[:3]  //取下标为3（不包括3）往前所有的元素组成切片
	h := e[1:3] //取下标为1（包括1）至下标为3（不包括3）的元素组成切片
	/*
		切片的长度：切片中包含的元素的个数
		切片的容量：切片的第一个元素到底层数组的最后一个元素的个数，始终会算到数组的最后一个元素
		f的容量：从下标为3（包括3）开始，至底层元素的最后一个的个数，也就是3
		g的容量：从下标为0（包括0）开始到底层数组的最后一个的个数（并不是到第三个为止，而是算到末尾），也就是6
		h的容量：从下标为1开始到底层数组的最后一个的个数，也就是5
	*/
	fmt.Printf("切片的长度：%d，切片的容量：%d\n", len(f), cap(f))
	fmt.Printf("切片的长度：%d，切片的容量：%d\n", len(g), cap(g))
	fmt.Printf("切片的长度：%d，切片的容量：%d\n", len(h), cap(h))
	fmt.Println("# 🚀4. 基于切片的切片")
	var i = []int{10, 13, 46, 93, 20, 78}
	j := i[2:]  //容量：4
	k := i[:4]  //容量：6
	l := i[2:4] //容量：4
	fmt.Printf("切片的长度：%d，切片的容量：%d\n", len(j), cap(j))
	fmt.Printf("切片的长度：%d，切片的容量：%d\n", len(k), cap(k))
	fmt.Printf("切片的长度：%d，切片的容量：%d\n", len(l), cap(l))
	fmt.Println("# 🚀5. 使用make构造函数定义切片")
	var m = make([]int, 4, 4) //make的参数为：切片类型，长度，容量
	fmt.Printf("值：%v，长度：%d，容量：%d\n", m, len(m), cap(m))
	fmt.Println("# 🚀6. 切片扩容")
	//append函数
	var n []int
	n = append(n, 3, 2, 12, 3) //一次性添加多个元素
	fmt.Printf("值：%v，长度：%d，容量：%d\n", n, len(n), cap(n))
	fmt.Println("# 🚀7. append将两个切片合并")
	var o = []int{1, 2, 3}
	var p = []int{4, 5, 6}
	o = append(o, p...) //append第二个参数为切片时必须加上...三个点
	fmt.Printf("值：%v，长度：%d，容量：%d\n", o, len(o), cap(o))
	fmt.Println("# 🚀8. 切片删除元素")
	var q = []string{"python", "golang", "java", "nodejs", "c++", "c#", "c"}
	//golang中切片没有专门的函数来删除元素，可以用append函数实现
	q = append(q[:3], q[4:]...)
	fmt.Printf("值：%v，长度：%d，容量：%d\n", q, len(q), cap(q))
	fmt.Println("# 🚀9. 复制切片，切片为引用类型，修改副本的值会连带修改原数据，使用copy方法可避免这种情况")
	var s = []int{12, 34, 45}
	// var t = s
	// s[0] = 112
	// fmt.Println(s, t)
	t := make([]int, 3, 3) //使用make构造一个用来装复制后的切片，不能使用 var t []int这种定义方式，没有容量无法将原切片复制到此
	copy(t, s)             //copy的参数为(目标切片，源切片)
	fmt.Println(s, t)
	s[0] = 112 //修改原切片后，副本没被修改
	fmt.Println(s, t)
	fmt.Println("# 🚀10. 切片的扩容策略")
	var u []int
	for i := 0; i < 10; i++ {
		u = append(u, i)
		fmt.Printf("值：%v，长度：%d，容量：%d\n", u, len(u), cap(u))
	}

}
