package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{24, 23, 12, 65, 3, 81}
	fmt.Println("# ð1. å°ç»å®çæ´ååçä»å°å°å¤§æåºï¼ä½¿ç¨éæ©æåºç®æ³")
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a)-1; j++ {
			// fmt.Printf("åï¼%v\n", a)
			if a[i] > a[j+1] {
				a[i], a[j+1] = a[j+1], a[i]
				// fmt.Printf("åï¼%v\n", a)
			}
		}
	}
	fmt.Println(a)
	var b = []int{23, 35, 54, 9, 12, 3, 89, 32, 73}
	fmt.Println("# ð2. å°ç»å®çæ´ååçä»å°å°å¤§æåºï¼ä½¿ç¨åæ³¡æåºç®æ³")
	for x := 0; x < len(b); x++ {
		for i := 0; i < len(b)-1; i++ {
			if b[i] > b[i+1] {
				b[i], b[i+1] = b[i+1], b[i]
			}
		}
	}
	fmt.Println(b)
	fmt.Println("# ð3. å°ç»å®çæ´ååçä»å°å°å¤§æåºï¼ä½¿ç¨sort()ï¼é»è®¤æ¯ååº")
	var c = []int{10, 8, 23, 14, 51}
	var d = []float64{10, 10.2, 58.1, 1.23}
	var e = []string{"a", "c", "b", "f", "e"}
	sort.Ints(c)
	sort.Float64s(d)
	sort.Strings(e)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("# ð3. ä½¿ç¨sort()éåº")
	var f = []int{10, 8, 23, 14, 51}
	var g = []float64{10, 10.2, 58.1, 1.23}
	var h = []string{"a", "c", "b", "f", "e"}
	sort.Sort(sort.Reverse(sort.IntSlice(f))) //sortæ¹æ³å¯¹æ´ååççéåºæåï¼éè¦åºå«ä¸ååºä¸­çsort.Intsï¼è¿éæ¯sort.IntSlice
	sort.Sort(sort.Reverse(sort.Float64Slice(g)))
	sort.Sort(sort.Reverse(sort.StringSlice(h)))
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

}
