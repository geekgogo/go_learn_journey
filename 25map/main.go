package main

import "fmt"

func main() {
	//map：无序的复合数据类型，以键值对的方式存在。与切片一样属于引用类型
	fmt.Println("# 1.定义与初始化map")
	//中括号里面是键的类型，后面是值的类型，两者可以不一致
	//第一种方式：使用make方法。make(map[string]string [cap]) cap为可选参数，map类型的容量
	var a = make(map[string]string) //定义map
	a["username"] = "gogo"
	a["age"] = "24"
	fmt.Println(a)
	//第二种方式：在声明是填充内容
	var b = map[string]string{
		"username": "geek",
		"gender":   "male", //逗号不能少
	}
	fmt.Println(b)
	//第三种方式
	c := map[int]string{
		0: "name",
		1: "gender", //逗号不能少
	}
	fmt.Println(c)
	fmt.Println("# 2.map的获取")
	username := b["username"]
	fmt.Printf("map数据类型b中username的值是：%v，gender是：%v\n", username, b["gender"])
	fmt.Println("# 3.map的查询，判断指定的键是否存在于map中")
	// map的查询返回两个值，存在则返回值和true，否则返回定义时键值对中值的默认值和false
	v, ok := b["username"]
	v1, ok1 := b["sex"]
	fmt.Printf("username是否在b中：%v，它的值是：%v\n", ok, v) //true 和 值
	fmt.Printf("sex是否在b中：%v，它的值是：%v\n", ok1, v1)    //false 和 空值（string的默认值为空）
	fmt.Println("# 4.map的修改")
	b["username"] = "geekgogo"
	fmt.Println(b)
	fmt.Println("# 5.map的删除，语法：delete(map对象, 要删除的key)")
	delete(b, "username")
	fmt.Printf("b删除了username键值对后：%v\n", b)
	fmt.Println("# 6.map的增加")
	b["height"] = "180"
	fmt.Printf("b增加了height键值对后：%v\n", b)
	fmt.Println("# 7.map的遍历")
	var d = map[string]string{
		"name":          "golang",
		"studyDuration": "25天",
		"today":         "3hrs31mins",
	}
	for k, v := range d {
		fmt.Printf("键：%v，值：%v\n", k, v)
	}

}
