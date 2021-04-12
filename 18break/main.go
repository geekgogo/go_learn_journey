package main

import "fmt"

func main() {
	//break，跳出循环
	fmt.Println("# 🚀1.当只有一层循环时，跳出循环")
	for i := 0; i < 10; i++ {
		if i > 3 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("# 🚀2.当有多层循环时，跳出当前层的循环")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j > 2 {
				break
			}
			fmt.Printf("i=%v,j=%v\n", i, j)
		}
	}
	fmt.Println("# 🚀3.当有多层循环时，可根据label标签跳出指定的循环")
label1: //label标签名称可以随意定义
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j > 2 {
				break label1 //直接跳出label1后面跟着的循环，这里是最外层的for循环
			}
			fmt.Printf("i=%v,j=%v\n", i, j)
		}
	}

}
