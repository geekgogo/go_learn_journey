package main

import (
	"fmt"
)

/*
单向管道：指的是仅能进行读操作或者写操作的管道
定义仅写操作的int型管道：var ch chan<- int
定义仅读操作的int型管道：var ch <-chan int
*/

func fn1(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func fn2(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("🙂 单向管道的定义")
	ch := make(chan<- int, 3) //定义写操作的单向管道
	ch <- 10
	ch <- 20
	//读操作报错
	// <-ch//invalid operation: <-ch (receive from send-only type chan<- int)
	// ch1 := make(<-chan int, 3) //定义读操作的单向管道
	//写操作报错
	// ch1 <- 1
	fmt.Println("🙃 单向管道的用处")
	//读写操作分离
	ch2 := make(chan int, 10)
	fn1(ch2)
	fn2(ch2)
	fmt.Println("😉 管道的多路复用：同时获取多个管道中的数据（不使用协程的情况下）")
	ch3 := make(chan int, 10)
	ch4 := make(chan string, 5)
	for i := 0; i < 10; i++ {
		ch3 <- i
	}
	array := [5]string{"赤", "橙", "黄", "绿", "蓝"}
	for _, v := range array {
		ch4 <- v
	}
	/*使用select实现多路复用。select类似switch的用法，它有一系列的分支和一个默认分支。每个case语句对一个管道的通信（接收或发送），select会一直
	等待，直到某个case语句的通信操作完成，最后执行default的语句，使用select不能关闭channel
	*/
	for {
		select {
		case x := <-ch3:
			fmt.Println(x)
		case x := <-ch4:
			fmt.Println(x)
		default:
			fmt.Println("数据使用完毕")
			return //跳出for循环
		}
	}

}
