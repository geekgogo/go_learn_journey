package main

/*
反射：reflect用于判断变量的类型，在空接口函数中用得较多
reflect.TypeOf():判断类型
reflect.ValueOf():查看原始值
reflect.Kind():判断原始类型（较多用于判断自定义类型或者空接口类型的原始类型）
reflect.ValueOf().SetInt():设置int类型的值
reflect.ValueOf().SetString():设置string类型的值
等等
reflect.ValueOf().Elem().SetInt(): 加上Elem()表示接收的参数为地址时设置值
*/

import (
	"fmt"
	"reflect"
)

func test(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	n := t.Name()
	k := t.Kind()
	fmt.Printf("原始类型是：%v，类型是：%v，类型名称是：%v，值是：%v\n", k, t, n, v)
	switch k {
	case reflect.Int:
		fmt.Printf("接收的是Int型：%v\n", v.Int())
	case reflect.String:
		fmt.Printf("接收的是string型：%v\n", v.String())
	case reflect.Bool:
		fmt.Printf("接收的是bool型：%v\n", v.Bool())
	case reflect.Float64:
		fmt.Printf("接收的是float64型：%v\n", v.Float())
	default:
		fmt.Println("无法判断类型")
	}
}

//通过reflect设置值，接收的参数需要是地址，否则会引发panic
func set(a interface{}) {
	v := reflect.ValueOf(a)
	//当传入的变量为指针类型时，可以用 Elem().Kind()获取类型，使用 Elem().SetInt()等等改变值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(100)
	}
}

func main() {
	fmt.Println("😘 反射概览")
	var a int = 10
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.Kind(a))
	reflect.ValueOf(&a).Elem().SetInt(23)
	fmt.Println(a)
	fmt.Println(reflect.ValueOf(a).String())
	fmt.Println("😘 反射的例子")
	var b string = "hello world"
	var c bool = false
	var d float64 = 10.34
	test(b)
	test(c)
	test(d)
	test(&b)
	fmt.Println("😘 反射的设置值")
	var g int64 = 1234
	set(&g)
	fmt.Println(g)

}
