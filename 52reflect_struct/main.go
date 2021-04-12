package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

// 获取结构体的信息
func (s Student) GetInfo() {
	fmt.Printf("姓名：%v，年龄：%v，分数：%v\n", s.Name, s.Age, s.Score)
}

// 修改结构体的信息，
func (s *Student) SetInfo(name string, age, Score int) {
	s.Name = name
	s.Age = age
	s.Score = Score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法")
}

func (s Student) ReInfo() string {
	name := s.Name
	return name
}

//结构体反射的相关操作
func reflectSturct1(a interface{}) {

	t := reflect.TypeOf(a) // 类型变量
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的不是一个结构体")
		return
	}
	// v : = reflect.Value(a) // 值变量
	//1. 根据类型变量获取结构体的字段方法一，不建议使用
	field0 := t.Field(0)
	fmt.Printf("字段类型：%v，字段名称：%v，字段jsonTag：%v\n", field0.Type, field0.Name, field0.Tag)
	// fmt.Printf("%#v\n", field0) //reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x10ad140), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	//2. 根据类型变量获取结构体的字段方法二
	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Printf("字段名称：%v，类型：%v，tag：%v\n", field1.Name, field1.Type, field1.Tag.Get("json"))
	}
	//3. 根据 NumField 判断结构体字段的个数
	fieldCount := t.NumField()
	fmt.Printf("有%v个字段\n", fieldCount)
	// 4. 获取属性对应的值
	v := reflect.ValueOf(a)
	// 4.1根据索引获取，不建议使用
	value0 := v.Field(0)
	fmt.Println(value0)
	// 4.2 根据字段名获取
	name := v.FieldByName("Name")
	age := v.FieldByName("Age")
	score := v.FieldByName("Score")
	fmt.Printf("姓名：%v，年龄：%v，分数：%v\n", name, age, score)
}

// 获取结构中的方法
func getFn(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的不是一个结构体")
		return
	}
	// 1. 根据索引获取方法
	method0 := t.Method(0) // 0和结构体方法的定义顺序没有关系，和结构体方法名称的ASCII码有关系
	fmt.Println(method0.Name)
	//2. 判断结构体中有多少个方法
	methodCount := t.NumMethod()
	fmt.Printf("方法数量：%v\n", methodCount)
	//3.通过方法名称获取方法，s为指针结构体时才可以找到 SetInfo()
	method1, ok := t.MethodByName("SetInfo")
	if ok {
		fmt.Println(method1)
	} else {
		fmt.Println("没有找到该方法")
	}
	// 4. 调用结构体的方法，使用值变量类型
	v := reflect.ValueOf(s)
	v.MethodByName("Print").Call(nil) // nil 表示调用函数时不传入任何参数
	v.MethodByName("GetInfo").Call(nil)
	// 调用函数时传参
	var params []reflect.Value
	params = append(params, reflect.ValueOf("小红"))
	params = append(params, reflect.ValueOf(15))
	params = append(params, reflect.ValueOf(100))
	v.MethodByName("SetInfo").Call(params) // Call() 接收的是 reflect.Value 类型的切片
	v.MethodByName("GetInfo").Call(nil)

	// 5. 获取方法的返回值，返回的是 reflect.Value 类型的切片
	r := v.MethodByName("ReInfo").Call(nil)
	fmt.Println(r)
	// 6. 修改结构体属性的值
	field := v.Elem().FieldByName("Name") //根据字段名称获取属性的值
	fmt.Println(field)
	field.SetString("芥川龙之介")
	fmt.Println(field)
	v.MethodByName("GetInfo").Call(nil)
}
func main() {
	stu := Student{
		Name:  "小明",
		Age:   23,
		Score: 99,
	}
	fmt.Printf("%#v\n", stu)
	reflectSturct1(stu)
	fmt.Println("-----------😗 获取和执行结构体中的方法-------------")
	getFn(&stu)
}
