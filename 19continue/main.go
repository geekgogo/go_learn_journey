package main

import "fmt"

func main() {
	//continue语句的核心在于 **跳过**
	fmt.Println("# 🚀1.continue语句，跳过所有大于3的整数")
	for i := 0; i < 10; i++ {
		if i > 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("# 🚀2.continue语句，跳过所有等于3的整数")
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("# 🚀3.continue结合label使用，表示跳到label指定的标签然后继续循环")
label2:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue label2 //跳到label2继续执行，与break不同
			}
			fmt.Printf("i=%v, j=%v\n", i, j)
		}
	}

}
