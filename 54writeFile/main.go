package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//写入文件的三种方法
	fmt.Println("-------1. os.WriteString------")
	//os.O_APPEND：追加
	file, err := os.OpenFile("./oswrite.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		file.WriteString("hello world1233123\r\n") // \r\n 表示在文件中换行
	}

	fmt.Println("-------2. bufio中的方法------")
	// os.O_TRUNC：清空
	file2, err2 := os.OpenFile("./bufiowrite.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file2.Close()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	write := bufio.NewWriter(file2)
	write.WriteString("这是使用bufio写入文件的方法") //写入缓存
	write.Flush()                         // 刷新缓存，写入文件
	fmt.Println("-------3. ioutil中的方法，ioutil中没有追加模式------")
	ioutil.WriteFile("./ioutil.txt", []byte("这是ioutil中的方法"), 0666)
}
