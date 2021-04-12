package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{24, 23, 12, 65, 3, 81}
	fmt.Println("# 🚀1. 将给定的整型切片从小到大排序，使用选择排序算法")
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a)-1; j++ {
			// fmt.Printf("前：%v\n", a)
			if a[i] > a[j+1] {
				a[i], a[j+1] = a[j+1], a[i]
				// fmt.Printf("后：%v\n", a)
			}
		}
	}
	fmt.Println(a)
	var b = []int{23, 35, 54, 9, 12, 3, 89, 32, 73}
	fmt.Println("# 🚀2. 将给定的整型切片从小到大排序，使用冒泡排序算法")
	for x := 0; x < len(b); x++ {
		for i := 0; i < len(b)-1; i++ {
			if b[i] > b[i+1] {
				b[i], b[i+1] = b[i+1], b[i]
			}
		}
	}
	fmt.Println(b)
	fmt.Println("# 🚀3. 将给定的整型切片从小到大排序，使用sort()，默认是升序")
	var c = []int{10, 8, 23, 14, 51}
	var d = []float64{10, 10.2, 58.1, 1.23}
	var e = []string{"a", "c", "b", "f", "e"}
	sort.Ints(c)
	sort.Float64s(d)
	sort.Strings(e)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("# 🚀3. 使用sort()降序")
	var f = []int{10, 8, 23, 14, 51}
	var g = []float64{10, 10.2, 58.1, 1.23}
	var h = []string{"a", "c", "b", "f", "e"}
	sort.Sort(sort.Reverse(sort.IntSlice(f))) //sort方法对整型切片的降序排列，需要区别与升序中的sort.Ints，这里是sort.IntSlice
	sort.Sort(sort.Reverse(sort.Float64Slice(g)))
	sort.Sort(sort.Reverse(sort.StringSlice(h)))
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

}
