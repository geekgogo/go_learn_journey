package info

import (
	"fmt"
)

func Info() {
	fmt.Println("这是info包中Info方法")
}

func init() {
	fmt.Println("info init ...")
}
