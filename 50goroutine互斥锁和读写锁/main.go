package main

import (
	"fmt"
	"sync"
	"time"
)

/*
当有多个协程访问同一个资源时，由于是并行执行，会导致资源争夺的现象，互斥锁保证了资源的一致性，当某一个协程正在使用时，其他资源必须等待其完成方能访问

小知识点：有些程序使用go run main.go时不能发现是否会发生资源争夺的现象。此时使用 go build -race main.go生成可执行文件，运行时便一目了然
*/

var wg sync.WaitGroup
var mutex sync.Mutex //引入互斥锁
//定义一个全局变量
var m = make(map[int]int, 0)

func m1(num int) {
	mutex.Lock() //加锁，当一个协程进来时，把资源锁定
	var sum int = 0
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum
	fmt.Println(m)
	mutex.Unlock() //结束释放锁，其他进程进来时可以使用资源
	wg.Done()
}

/*
读写锁
作用是保证读操作并行执行，但对于写操作是完全互斥的。意思就是，当一个goroutine进行写操作时，其他goroutine既不能读也不能写
*/
var rwmutex sync.RWMutex //定义读写互斥锁
func write() {
	//给写操作加锁，使其互斥，顺序执行
	rwmutex.Lock()
	fmt.Println("正在进行写入操作")
	time.Sleep(time.Second * 2)
	rwmutex.Unlock()
	wg.Done()
}

func read() {
	//给读操作做加读写锁，并发执行
	rwmutex.RLock()
	fmt.Println("正在进行读取操作")
	time.Sleep(time.Second * 2)
	rwmutex.RUnlock()
	wg.Done()
}

func main() {
	fmt.Println("😍 goroutine互斥锁:sync.Mutex")
	//开启40个协程运行m1，由于协程是并行，而m是全局变量，会引发抢占资源的现象，程序无法正常运行
	// for i := 0; i < 40; i++ {
	// 	wg.Add(1)
	// 	go m1(10)
	// }
	//正确写法，引入互斥锁
	// for i := 0; i < 40; i++ {
	// 	wg.Add(1)
	// 	go m1(i)
	// }
	// wg.Wait()
	fmt.Println("🥰 goroutine读写互斥锁")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	// 可以发现读取操作是一起打印，而写入操作是隔2秒打印的
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()

}
