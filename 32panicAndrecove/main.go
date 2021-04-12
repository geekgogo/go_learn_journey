package main

import (
	"errors"
	"fmt"
)

/*
golang(Go1.12)中没有异常机制：可以使用panic和recover模式来处理错误
panic可以在任何地方调用，recover只能在defer语句中调用
*/
func f1() {
	fmt.Println("f1")
}
func f2() {
	panic("我是panic手动造成的异常")
}

//使用revcover来接收异常
func f3() {
	defer func() {
		err := recover() //这里接收到了panic造成的异常，并打印
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("我是panic手动造成的异常")
}

//例1:
func f4(a, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	return a / b
}

//例2:
//定义一个读取文件的函数
func readFile(name string) error {
	if name == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func f5() {
	err := readFile("xx.go")
	defer func() {
		error := recover()
		if error != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	if err != nil {
		panic(err)
	}
}
func main() {
	f1()
	// f2()
	f3()
	fmt.Println(f4(10, 0))
	fmt.Println(f4(10, 2))
	fmt.Println("end")
	fmt.Println("📧 例2:")
	f5()
	fmt.Println("继续执行")
}
