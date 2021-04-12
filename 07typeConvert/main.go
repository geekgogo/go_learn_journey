package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println("# 🚀1. int类型相互转换，建议从低位转向高位，可避免精度丢失")
	var a int32 = 32
	var b int64 = 64
	// fmt.Println(a + b) // 直接相加会报错，因为是不同类型的整型
	fmt.Printf("%v--%T\n", int64(a)+b, int64(a)+b)
	fmt.Println("# 🚀2. float类型的相互转换，建议从低位转向高位，可避免精度丢失")
	var c float32 = 32
	var d float64 = 64
	fmt.Printf("%v--%T\n", float64(c)+d, float64(c)+d)
	fmt.Println("# 🚀3. int类型和float类型的相互转换")
	var e int = 32
	var f float64 = 65
	fmt.Printf("%v--%T\n", float64(e)+f, float64(e)+f)
	fmt.Printf("%v--%T--%v\n", e+int(f), e+int(f), unsafe.Sizeof(e+int(f)))
	fmt.Println("# 🚀4. 其他类型转换成字符串型")
	fmt.Println("# 🚀4.1 使用fmt.Sprintf方法")
	var g int = 64
	var h float64 = 64.64444
	var i bool
	var j byte = 'j'
	str1 := fmt.Sprintf("%d", g) // %d 接收整型
	str2 := fmt.Sprintf("%f", h) // %f 接收浮点型
	str3 := fmt.Sprintf("%t", i) // %t 接收布尔型
	str4 := fmt.Sprintf("%c", j) // %c 接收字符型
	fmt.Printf("%v, %v, %v, %v\n", str1, str2, str3, str4)
	fmt.Println("# 🚀4.2 使用strconv方法")
	/* FormatInt接收两个参数
	1. 必须是int64的整型
	2. 进制
	*/
	str5 := strconv.FormatInt(int64(g), 10)
	fmt.Printf("int转换为字符串%v\n", str5)
	/* FormatFloat接收4个参数
	1. 需要转换的值，必须是float64类型
	2. 转换后的格式，'f'表示小数点
		'b' (-ddddp±ddd, a binary exponent),
		'e' (-d.dddde±dd, a decimal exponent),
		'E' (-d.ddddE±dd, a decimal exponent),
		'f' (-ddd.dddd, no exponent),
		'g' ('e' for large exponents, 'f' otherwise),
		'G' ('E' for large exponents, 'f' otherwise),
		'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
		'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
	3. 保留的小数点的个数
	4. 位数（32或64）
	*/
	str6 := strconv.FormatFloat(float64(h), 'f', 2, 64)
	fmt.Printf("float转换为字符串%v\n", str6)
	//bool类型转换为string类型意义不大
	str7 := strconv.FormatBool(i)
	fmt.Printf("bool类型转换为string类型：%v\n", str7)
	/* byte型转换为string型 接收两个参数
	1. 必须是将字符转换为unint64
	2. 进制
	*/
	str8 := strconv.FormatUint(uint64(j), 10)
	fmt.Printf("byte类型转换为string类型：%v\n", str8)
	fmt.Println("# 🚀5. 字符串转换为其他类型")
	var k string = "123456"
	/* string转int接收三个参数
	1. 代转换的值
	2. 进制
	3. 位数
	并且返回两个值
	1. 结果
	2. 报错
	*/
	str9, _ := strconv.ParseInt(k, 10, 64)
	fmt.Printf("string转整型，strconv.ParseInt：%v，类型：%T\n", str9, str9)
	/* string转float接收两个参数
	1. 代转换的值
	2. 位数
	*/
	str10, _ := strconv.ParseFloat(k, 64)
	fmt.Printf("string转float，strconv.ParseFloat：%v，类型：%T\n", str10, str10)
	str11, _ := strconv.ParseBool(k)
	fmt.Printf("string转bool，strconv.ParseBool：%v，类型：%T\n", str11, str11)
	str12, _ := strconv.ParseUint(k, 10, 64)
	fmt.Printf("string转byte，strconv.ParseUint：%v，类型：%T\n", str12, str12)

}
