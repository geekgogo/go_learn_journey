package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

func main() {
	//定义一个map类型数据
	a := make(map[string]interface{})
	a["name"] = "李白"
	a["age"] = 34
	a["hobby"] = []string{"写诗", "喝酒"}
	a["address"] = Address{
		Name:  "杭州",
		Phone: 1587111111,
	}
	//我想获取a中hobby的元素  写诗
	//错误写法
	// fmt.Println(a["hobby"][0]) //invalid operation: a["hobby"][0] (type interface {} does not support indexing)
	//正确写法，使用断言
	//这条语句的意思是：判断a["hobby"]是否为切片类型，两个返回值：1.值 2.bool值
	value, _ := a["hobby"].([]string)
	fmt.Println(value[0])
	// 获取a中address的Name和Phone
	//错误写法
	// fmt.Println(a["address"].Name) //a["Address"].Name undefined (type interface {} is interface with no methods)
	//正确写法，同样使用断言
	v, _ := a["address"].(Address)
	fmt.Println(v.Name)
	fmt.Println(v.Phone)

}
