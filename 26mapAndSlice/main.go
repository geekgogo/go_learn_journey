package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("# 1. 定义map类型的切片")
	var a = make([]map[string]string, 2, 2)
	b := make(map[string]string)
	b["username"] = "kaik"
	b["age"] = "24"
	c := map[string]string{
		"username": "song",
		"age":      "24",
	}
	a[0] = b
	a[1] = c
	fmt.Println(a)
	fmt.Println("# 2. 定义值为切片类型的map")
	var d = make(map[string][]string)
	d["username"] = []string{
		"kai",
		"song",
	}
	d["age"] = []string{
		"20",
		"30",
		"40",
	}
	fmt.Println(d)
	fmt.Println("# 3. 循环遍历上述两种的复合类型")
	fmt.Println("## 3.1 循环元素为map类型的切片")
	for _, v := range a {
		for key, val := range v {
			fmt.Printf("键：%v, 值：%v\n", key, val)
		}
	}
	fmt.Println("## 3.2 循环值为切片类型的map")

	for k, v := range d {
		for _, value := range v {
			fmt.Printf("键：%v, 值：%v\n", k, value)

		}
	}
	fmt.Println("# 4. map的排序")
	var e = make(map[int]int)
	e[11] = 23
	e[34] = 12
	e[8] = 9
	e[1] = 91
	e[97] = 2
	e[98] = 22
	e[4] = 12
	//下面直接循环打印会乱序输出
	// for k, v := range e {
	// 	fmt.Printf("key:%v, value:%v\n", k, v)
	// }
	//思路：现将map的key存放到切片中，对切片进行排序，然后按照升序的依次打印
	var keySlice []int
	for k := range e {
		keySlice = append(keySlice, k)
	}
	sort.Ints(keySlice)
	fmt.Println(keySlice)
	for _, v := range keySlice {
		fmt.Printf("key:%v, value:%v\n", v, e[v])
	}
	fmt.Println("# 5. 统计字符串how do you do中每个单词出现的次数")
	var f string = "how do you do"
	//将字符串以空格为标准切分到一个切片中
	var g []string
	g = strings.Split(f, " ")
	//循环这个切片，出现相同的单词则加一
	var h = make(map[string]int)
	for _, v := range g {
		h[v]++
	}
	fmt.Println(h)
}
