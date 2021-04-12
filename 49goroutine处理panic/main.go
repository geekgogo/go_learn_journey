package main

import (
	"fmt"
	"time"
)

// goroutine 处理错误
func say() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}
}

// 使hello方法出现错误
func hello() {
	// 解决方法，捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a map[int]string
	a[0] = "golang"
}
func main() {
	fmt.Println("😌 goroutine中的panic处理：多个协程中运行时，若有一个出现错误，则会影响其他协程的运行。")
	go say()
	go hello()
	// 防止主进程退出使用time.Sleep
	time.Sleep(time.Second)
}
