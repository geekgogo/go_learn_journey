package calc

import (
	"39gomod/info"
	"fmt"
)

// 大写字母开头表示公有方法，可以在其他包中调用
func Add(x, y int) int {
	return x + y
}

func GetInfo() {
	info.Info()
}

func init() {
	fmt.Println("calc init ...")
}
