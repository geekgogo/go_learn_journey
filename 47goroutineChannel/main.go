package main

import (
	"fmt"
	"sync"
	"time"
)

/* 使用 goroutine和channel打印1到120000的素数
解题思路：
定义三个channel，分别为存储1到120000数据的ch1，存数素数的ch2，存储ch2是否能关闭的状态ch3.
定义三个函数，存储数据的putNum，存储素数的primeNum, 打印素数的printNum,以及主程序中的判断ch3是否取完的匿名函数.这些函数同时执行
为什么要定义一个bool类型的ch3？
因为存储素数的函数中有一个将素数存入ch2的操作，打印素数printNum使用的for range方法，此方法必须要关闭ch2才能正常运行。所以所有
存数素数的协程完成时需要关闭ch2.因此引入一个ch3的管道，每执行一次primeNum，在ch3中放入一个bool值。然后在主程序中循环取出ch3的值，
取完则表示所有的primeNum执行完成（管道边存边取的操作，取会等待存。详情在上一个目录46Chanel中可以了解），继而关闭ch2.

*/
var wg sync.WaitGroup

//存数据
func putNum(ch chan int) {
	for i := 2; i < 200000; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

// 存素数
func primeNum(ch1 chan int, ch2 chan int, ch3 chan bool) {
	for v := range ch1 {
		var flag = true
		for i := 2; i < v; i++ {

			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			ch2 <- v
		}
	}
	ch3 <- true
	wg.Done()
}

// 打印素数
func printPrime(ch chan int) {

	for v := range ch {
		fmt.Printf("%v是素数\n", v)
	}
	wg.Done()
}
func main() {
	ch1 := make(chan int, 200000)
	ch2 := make(chan int, 200000)
	ch3 := make(chan bool, 32)
	start := time.Now().Unix()
	wg.Add(1)
	go putNum(ch1)
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go primeNum(ch1, ch2, ch3)
	}
	wg.Add(1)
	go func(closeCh chan bool, ch chan int) {
		for i := 0; i < 32; i++ {
			<-closeCh
		}
		close(ch)
		wg.Done()
	}(ch3, ch2)
	wg.Add(1)
	go printPrime(ch2)
	wg.Wait()
	end := time.Now().Unix()
	fmt.Printf("程序执行完毕，耗时%v秒\n", end-start)

}
