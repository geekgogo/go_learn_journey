package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("-----1.创建单级目录-----")
	err1 := os.Mkdir("./a", 0666)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("-----2.创建多级目录-----")
	err2 := os.MkdirAll("./dir1/dir2/dir3", 0777)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("-----3.目录与文件的删除-----")
	err3 := os.Remove("./del1.txt")
	if err3 != nil {
		fmt.Println(err3)
	}
	err4 := os.Remove("./del2")
	if err4 != nil {
		fmt.Println(err4)
	}
	err5 := os.RemoveAll("./del3")
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println("-----4.目录与文件的重命名-----")
	err6 := os.Rename("./a", "./b")
	if err6 != nil {
		fmt.Println(err6)
	}

}
