package main

import "fmt"

func main() {
	//switch 可对大量的值进行判断
	fmt.Println("# 🚀1. 例一：")
	var a = "xxx"
	switch a {
	case ".js":
		fmt.Println("这是JavaScript文件")
		break
	case ".html":
		fmt.Println("这是html文件")
		break
	default:
		fmt.Println("这是未知文件")
	}
	fmt.Println("# 🚀2. 例二：")
	switch b := ".html"; b { //变量的作用域不同
	case ".js":
		fmt.Println("这是JavaScript文件")
		break
	case ".html":
		fmt.Println("这是html文件")
		break
	default:
		fmt.Println("这是未知文件")
	}
	fmt.Println("# 🚀3. 例三：")
	/*
		当case语句后面跟条件判断时，switch后面不用接变量
	*/
	var score = 91
	switch {
	case score > 90:
		fmt.Println("优秀")
		break
	case score <= 90 && score > 60:
		fmt.Println("中等")
		break
	case score <= 60 && score >= 0:
		fmt.Println("差劲")
	default:
		fmt.Println("输入错误")
	}
	fmt.Println("# 🚀4. 例四：一个case可以有多个值，用逗号分隔")
	var num = 8
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("num是奇数")
		break
	case 2, 4, 6, 8, 10:
		fmt.Println("num是偶数")
		break
	}
	fmt.Println("# 🚀5. 例五：switch中的穿透问题：fallthrough")
	/*
		fallthrough可以执行满足条件的case语句的下一个case语句，但只能穿透一层
	*/

	var score2 = 91
	switch {
	case score2 > 90:
		fmt.Println("优秀")
		fallthrough
	case score2 <= 90 && score2 > 60:
		fmt.Println("中等")
		break
	case score2 <= 60 && score2 >= 0:
		fmt.Println("差劲")
	default:
		fmt.Println("输入错误")
	}

}
