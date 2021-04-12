package main

import (
	"fmt"
	"sort"
)

//函数接收slice作为参数
//定义一个选择排序的升序函数

func sliceAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

//函数接收map类型作为参数
//定义一个函数，以map类型作为参数，返回一个字符串，将map的key和value使用 => 按照字符串升序连接
func toArrow(map1 map[string]string) string {
	var sliceB []string
	for k := range map1 {
		sliceB = append(sliceB, k)
	}
	sort.Strings(sliceB)
	var str string
	for _, v := range sliceB {
		str += fmt.Sprintf("%v=>%v", v, map1[v])
	}
	return str
}

var k = 10

func test1() {
	var j = 20
	fmt.Println(j) // 局部变量
	fmt.Println(k) //全局变量
}

func main() {
	var sliceA = []int{10, 23, 65, 7, 91}
	r := sliceAsc(sliceA)
	var map1 = map[string]string{
		"username": "jack",
		"age":      "23",
		"address":  "whitehouse",
		"sex":      "male",
	}
	r1 := toArrow(map1)
	fmt.Println(r)
	fmt.Println(r1)
	//变量的作用域
	//全局变量和局部变量
	//在函数外定义的变量是全局变量，在函数内定义的变量是局部变量
	fmt.Println(k)
	test1()
	//下面的i是块作用域（也是局部变量），只在for循环内能调用。
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

}
