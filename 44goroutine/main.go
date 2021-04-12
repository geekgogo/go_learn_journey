package main

import (
	"fmt"
	"sync"
	"time"
)

/*
进程、线程、并行、并发
并行：多个cpu同时执行多个任务（任务数大于cpu时，并行和并发同时存在）
并发：单核cpu轮流执行多个任务
*/

// 定义一个test函数，每隔100毫秒打印 2
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("2")
		time.Sleep(time.Millisecond * 100)
	}
}
func test1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("协程test1执行%v次\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}

// 定义一个全局的sync变量
var wb sync.WaitGroup

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("协程test2执行%v次\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	wb.Done() //表示协程运行完成
}

func test3(num int) {
	defer wb.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("协程[%v]执行%v次\n", num, i)
	}

}
func main() {
	fmt.Println("goroutine多协程实现并行")
	// 顺序执行
	// test()
	// 主函数中每隔100毫秒打印 hello world
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("1")
	// 	time.Sleep(time.Millisecond * 100)
	// }

	// 使用goroutine多协程执行，结果为 1 和 2 交替打印。
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("1")
		time.Sleep(time.Millisecond * 100)
	}
	// 注意：使用协程时，若主线程运行时间比协程时间短，主线程结束时，不管协程有没有结束会一并退出
	fmt.Println("😁 协程实例2")
	// 这个例子可以发现协程 test1并没有执行完成
	go test1()
	for i := 0; i < 10; i++ {
		fmt.Printf("主线程2，执行第%v次\n", i)
		time.Sleep(time.Millisecond * 50)
	}

	fmt.Println("😆 避免主线程运行时间比协程时间短造成协程未运行完成就退出的解决方法")
	wb.Add(1) // 监听一个协程
	go test2()
	for i := 0; i < 10; i++ {
		fmt.Printf("主线程3，执行第%v次\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	wb.Wait() // 等待wb完成
	fmt.Println("😅 多个协程监听")
	for i := 0; i < 5; i++ {
		wb.Add(1)
		go test3(i)
	}
	wb.Wait()
	fmt.Println("主线程关闭")

}
