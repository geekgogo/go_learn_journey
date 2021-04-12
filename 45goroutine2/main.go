package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ordinary(num int) {
	for i := 2; i < num; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%v是素数\n", i)
		}
	}
}

func xc(num int) {
	for i := (num-1)*50000 + 1; i < num*50000; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%v是素数\n", i)
		}
	}
	wg.Done()
}
func main() {
	fmt.Println("😂 输出1～100000中的素数")
	fmt.Println("普通方法")
	// start := time.Now().Unix()
	// ordinary(200000)
	// end := time.Now().Unix()
	// fmt.Printf("耗时：%v秒\n", end-start) // 耗时：19秒
	fmt.Println("goroutine协程：开启四个协程，分别计算1～50000，50001～100000，100001～150000，150001～200000")
	start1 := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go xc(i)
	}
	wg.Wait()
	end1 := time.Now().Unix()
	fmt.Printf("goroutine耗时：%v秒\n", end1-start1) // goroutine耗时：13秒

}
