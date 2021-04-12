package main

import "fmt"

func getUserInfo() (string, int) {
	return "jack", 10
}

func main() {
	// _ 为匿名变量，不占用命名空间，不分配内存
	a, _ := getUserInfo()
	fmt.Println(a)
}
