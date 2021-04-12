package main

import "fmt"

func main() {
	fmt.Println("# 🚀1.golang中字符型是用单引号定义，字符串用双引号定义。一个字母或者汉字是一个字符。字符的数据类型为unit8")
	var a byte = 'a'
	fmt.Printf("ASCII值：%v，值：%c 类型：%T\n", a, a, a) //字符型数据 %v 输出的是对应的ASCII码的值，%c才是原样输出
	fmt.Println("# ⚠️ 2.一个英文字符占用1个字节，一个汉字占用3个字节")
	fmt.Println("# 🚀3. 字符型分为两种，一种是byte类型或者unit8类型，代表一个ASCII码；一种是rune类型或int32类型，代表一个utf-8字符")
	var b byte = 'b'
	c := '你' //golang中汉字是用utf-8表示
	fmt.Printf("%v(%c) %T\n", b, b, b)
	fmt.Printf("%v(%c) %T\n", c, c, c)
	fmt.Println("# 🚀4. 输出字符串中的字符")
	var d string = "golang"
	fmt.Printf("%v(%c), %T\n", d[0], d[0], d[0])
	fmt.Println("# 🚀5. 通过循环打印字符串中的字符")
	e := "你好， golang"
	for i := 0; i < len(e); i++ { //普通for循环是通过byte类型打印，无法打印中文
		fmt.Printf("%v(%c)", e[i], e[i])
	}
	fmt.Printf("\n")
	for _, r := range e { //range打印是通过rune类型，可正常打印汉字
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Printf("\n")
	fmt.Println("# 🚀6. 修改字符串，golang中修改字符串必须先转换为字符数组([]byte或者[]rune)，然后转换为字符串")
	f := "big"
	g := []byte(f) //强制转换为byte型数组
	g[0] = 'p'
	fmt.Printf("把big改成pig：%v\n", string(g))
	h := "我的"
	i := []rune(h) //强制转换为rune型数组
	i[0] = '你'
	fmt.Printf("修改汉字的字符：%v\n", string(i))
}
