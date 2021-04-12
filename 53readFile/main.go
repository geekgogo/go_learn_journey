package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 读取文件的三种方法
	fmt.Println("-------------1. 使用 os中的Read方法读取文件（以流的方式读取文件）-------------")
	file, err := os.Open("./test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var tempSlice = make([]byte, 128) // 128 指定的是每次读取的字节数
	var fileSlice []byte
	// 循环读取所有文件内容
	for {
		n, err := file.Read(tempSlice) //Read()接收一个byte类型的切片用于存放读取的数据
		//判断文件读取完毕
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("字节数：%v\n", n)
		fileSlice = append(fileSlice, tempSlice...)
	}
	fmt.Println(string(fileSlice))
	fmt.Println("-------------2. 使用 bufio 读取文件（以流的方式读取文件）-------------")
	file2, err2 := os.Open("bufiotxt.txt")
	defer file2.Close()
	if err2 != nil {
		fmt.Println(err)
		return
	}
	read := bufio.NewReader(file2)    // 创建一个Reader对象
	str, err := read.ReadString('\n') // '\n' 表示读取一行
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	//读取所有文本则使用for循环
	file3, err3 := os.Open("./bufiotxt2.txt")
	defer file3.Close()
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	read2 := bufio.NewReader(file3)
	var str3 string
	for {
		str2, err4 := read2.ReadString('\n')
		if err4 == io.EOF {
			str3 += str2
			break
		}
		if err4 != nil {
			fmt.Println(err4)
			return
		}
		str3 += str2
	}
	fmt.Println(str3)
	fmt.Println("-------------3. 使用 ioutil 读取文件，读取小文件建议使用（100到200M可行）-------------")
	file5, err5 := ioutil.ReadFile("ioutil.txt")
	if err5 != nil {
		fmt.Println(err5)
		return
	}
	fmt.Println(string(file5))
}
