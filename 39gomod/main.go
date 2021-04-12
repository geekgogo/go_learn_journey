package main

import (
	"39gomod/calc"
	"fmt"
)

// init会在main方法之前被系统调用，无需手动调用。

/*
当调用的包中也含有init方法时，它的调用顺序为，最先导入的包最后调用
例如：
这里调用了 39gomod/calc，它有一个init方法打印 calc init...
39gomod/calc中调用了39gomod/info，它也有一个init方法打印 info init ...
这里的包的调用顺序为：main.go --> calc.go --> info.go
因此他们的init方法执行顺序为：info.go --> calc.go --> main.go
*/
func init() {
	fmt.Println("main init ...")
}

func main() {
	fmt.Println("🍅 go mod的用法")
	sum := calc.Add(10, 20)
	fmt.Println(sum)
	calc.GetInfo()
}
