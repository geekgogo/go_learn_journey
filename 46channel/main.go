package main

import (
	"fmt"
	"sync"
	"time"
)

//管道遵循先入先出的规则。管道类似于一个存储数据的容器。管道是一种特殊的数据类型，且是一种引用类型。管道的3个操作：写入，读取，关闭
var wg sync.WaitGroup

//小知识点：当管道写入数据比读取数据慢时，读取操作会等待写入操作，并不会引起异常
func fn1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("【写入数据】：%v\n", i)
		time.Sleep(time.Millisecond * 250)
	}
	close(ch)
	wg.Done()
}

func fn2(ch chan int) {
	for v := range ch {
		fmt.Printf("【读取数据】：%v\n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}
func main() {
	//创建管道
	// var ch chan int
	ch := make(chan int, 3) //创建一个容量为3的int类型管道
	//管道接收数据
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("管道取出数据")
	x := <-ch
	fmt.Println(x)
	<-ch //从管道里面取值不赋值给任何变量
	z := <-ch
	fmt.Println(z) // z的值为3

	//此时管道的数据已经全部取出，又可以重新存入数据
	ch <- 4
	ch <- 5
	fmt.Printf("值：%v，容量：%v，长度：%v\n", ch, cap(ch), len(ch))
	//管道阻塞
	//管道容量已满继续存入数据或者容量为空继续取出数据都会触发管道阻塞（死锁现象）
	ch1 := make(chan int, 2)
	ch1 <- 12
	ch1 <- 13
	// ch1 <- 14 //因为容量为2，存入第三个数据时触发阻塞 fatal error: all goroutines are asleep - deadlock!
	<-ch1
	<-ch1
	// <-ch1 //因为容量仅为2，已经取完两个数据，继续取会引发阻塞
	fmt.Println("管道的循环遍历和关闭")
	ch2 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch2 <- i
	}
	close(ch2) //使用range取值，必须要关闭管道。而使用for循环则不需要
	fmt.Println("☺️ 使用range取值，需要注意的是管道没有key值，且管道需要关闭")
	for v := range ch2 {
		fmt.Println(v)
	}
	ch3 := make(chan int, 6)
	for i := 0; i < 6; i++ {
		ch3 <- i
	}
	fmt.Println("🤣 使用for循环，管道是否关闭都可")
	for i := 0; i < 6; i++ {
		x := <-ch3
		fmt.Println(x)
	}
	fmt.Println(`
需求：定义两个方法，一个给管道写入数据，一个从管道读取数据。同时进行
1. 开启一个fn1协程向管道ch4中写入10条数据
2. 开启一个fn2协程从管道ch4读取数据
3. fn1和fn2操作一个管道
4. 主线程必须等协程完成才能退出
	`)
	ch4 := make(chan int, 100)
	wg.Add(1)
	go fn1(ch4)
	wg.Add(1)
	go fn2(ch4)
	wg.Wait()
	fmt.Println("退出。。。")
}
