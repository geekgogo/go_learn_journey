package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("# 🚀1.golang的字符串默认为空，字符串必须用双引号")
	var a string = "this is a string"
	var b string
	// var c string = 'string' //单引号定义string类型报错
	fmt.Printf("a的值是：%v, a的类型是： %T\n", a, a)
	fmt.Printf("字符串的默认值是：%v\n", b)
	fmt.Println("# 🚀2. strings的转义符：\\")
	// \\表示反斜杠，\"表示引号，\r表示回车，\n表示换行等等
	fmt.Println("# 🚀3. 使用反引号定义多行字符串，并且输出会保持原样")
	var d string = `使用反引号定义多行字符串
							使用反引号定义多行字符串
	使用反引号定义多行字符串
	使用反引号定义多行字符串`
	fmt.Printf("d的值是：%v\n", d)
	fmt.Println("# 🚀4. 求字符串的长度：len()方法")
	var e string = "this is a str" //计算长度时空格也会加入
	f := "你好"                      //一个汉字占用3个字节
	fmt.Printf("e的长度是：%v，f的长度是：%v\n", len(e), len(f))
	fmt.Println("# 🚀5.拼接字符串：+ 或者fmt.Sprintf")
	var g string = "你好"
	var h string = "golang"
	i := g + h
	j := fmt.Sprintf("%v-*-%v", g, h) //fmt.Sprintf可以返回输出的结果
	fmt.Println(i, "分隔符", j)
	fmt.Println("直接拼接多行字符串")
	// + 号必须放在字符串的末尾
	k := "使用反引号定义多行字符串" +
		"使用反引号定义多行字符串" +
		"使用反引号定义多行字符串"
	fmt.Println(k) // 多行字符串拼接后为一行输出
	fmt.Println("# 🚀6. 分割字符串：strings.Split")
	var m string = "188-8888-8888"
	n := strings.Split(m, "-")
	fmt.Println(n) //输出切片
	fmt.Println("# 🚀7. 将切片结合成字符串：strings.Join")
	fmt.Println(strings.Join(n, "**"))
	fmt.Println("# 🚀8. 判断是否包含strings.Contains")
	var o string = "this is a is str is"
	var p string = "this"
	fmt.Printf("o是否包含p:%v\n", strings.Contains(o, p))
	fmt.Println("# 🚀9.判断字符串的前缀HasPrefix和后缀HasSuffix(从0开始)")
	fmt.Printf("o的前缀是否为p：%v\n", strings.HasPrefix(o, p))
	fmt.Printf("o的后缀是否为p：%v\n", strings.HasSuffix(o, p))
	fmt.Printf("o的后缀是否为str：%v\n", strings.HasSuffix(o, "str"))
	fmt.Println("# 🚀10.字符串的索引")
	fmt.Printf("p在o中的从前往后的索引位置是：%v\n", strings.Index(o, p))
	fmt.Printf("is在o中的从前往后索引位置是：%v\n", strings.Index(o, "is"))
	fmt.Printf("is在o中的从后往前索引位置是：%v\n", strings.LastIndex(o, "is"))
	fmt.Println("⚠️ LastIndex指的是从后往前算第一次出现字符串时，在整个字符串中的位置（位置的计算依然是从前往后以0开始）")

}
