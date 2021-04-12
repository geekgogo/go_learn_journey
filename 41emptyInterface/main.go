package main

import "fmt"

// A 定义一个空接口
type A interface{}

// 断言
func testType1(x interface{}) {
	if _, ok := x.(string); ok {
		fmt.Println("传入的参数是string类型")
	} else if _, ok := x.(int); ok {
		fmt.Println("传入的参数是int类型")

	} else {
		fmt.Println("传入的参数是其他类型")

	}
}

// x.(type):接口.(type)断言只能在switch语句中使用，不能当作单独的语句
func testType2(x interface{}) {
	switch x.(type) {
	case string:
		fmt.Println("传入的参数是string类型")
	case int:
		fmt.Println("传入的参数是int类型")
	default:
		fmt.Println("传入的参数是其他类型")

	}
}
func main() {
	fmt.Println("🥬 接口中若不约定任何方法，则为空接口。空接口可以表示任何类型")
	// 实现空接口
	var a A
	var str = "你好golang"
	a = str // 用str实现空接口
	var b A = 20
	fmt.Printf("值：%v，类型：%T\n", a, a)
	fmt.Printf("值：%v，类型：%T\n", b, b)
	fmt.Println("🧐 接口也是一种类型，直接定义空接口")
	var c interface{}
	c = "golang"
	fmt.Println(c)
	c = 20
	fmt.Println(c)
	c = false
	fmt.Println(c)
	fmt.Println("🥒 利用空接口使切片的元素为任意类型")
	//普通切片定义只能使元素为同一种类型
	d := []int{1, 2, 3, 4}
	// 使用空接口实现多种类型
	d1 := []interface{}{1, 2, "你好golang", true, 1.34}
	fmt.Printf("普通切片：%v，任意类型的切片：%v\n", d, d1)
	fmt.Println("🌶 利用空接口使map的值为任意类型")
	//普通map
	e := make(map[string]string)
	e["name"] = "Lena"
	e["age"] = "30"
	// 空接口实现的map
	e1 := make(map[string]interface{})
	e1["name"] = "Banana"
	e1["age"] = 30
	e1["married"] = false
	e1["height"] = 176.5
	fmt.Printf("普通map：%v，值为任意类型的map：%v\n", e, e1)
	fmt.Println("🌽 空接口的断言：用来判断空接口接收数据的类型。使用方法：见testType1和testType2函数")
	testType1("你好golang")
	testType2(120)

}
