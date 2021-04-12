package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func myIoutilCopy(srcFile, dstFile string) error {
	sFile, err1 := ioutil.ReadFile(srcFile)
	if err1 != nil {
		return err1
	}
	err2 := ioutil.WriteFile(dstFile, sFile, 0666)
	if err2 != nil {
		return err2
	}
	return nil
}

func myOsCopy(srcFile, dstFile string) error {
	sFile, err1 := os.Open(srcFile)
	dFile, err3 := os.OpenFile(dstFile, os.O_CREATE|os.O_RDWR, 0666)
	defer sFile.Close()
	defer dFile.Close()
	if err1 != nil {
		return err1
	}
	if err3 != nil {
		return err3
	}
	var tempSlice = make([]byte, 1280)
	for {
		n, err2 := sFile.Read(tempSlice)
		if err2 == io.EOF {
			break
		}
		_, err4 := dFile.Write(tempSlice[:n])
		if err4 != nil {
			return err4
		}
	}
	return nil
}

func main() {
	fmt.Println("-------1.使用ioutil中的方法复制文件--------")
	err := myIoutilCopy("./src.txt", "./ioutilDst.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("复制完成")
	fmt.Println("-------2.使用os中的Open方法复制文件--------")
	err2 := myOsCopy("./myOsCopy.txt", "./myOsCopyDst.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("复制完成")

}
