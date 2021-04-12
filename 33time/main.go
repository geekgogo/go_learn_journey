package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("打印当前时间")
	timeObj := time.Now()
	fmt.Println(timeObj)
	fmt.Println("转换为可读格式")
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	//%02d：表示2位数的宽度，不足两位按0补齐
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Println("转为可读格式的第二种方式：format函数")
	/*golang中时间格式转换与其他语言不同，它的格式为：
	2006-01-02 03:04:05来表示年-月-日 时:分:秒（12小时制）
	2006-01-02 15:04:05来表示年-月-日 时:分:秒（24小时制）
	*/
	now12 := timeObj.Format("2006-01-02 03:04:05")
	now24 := timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(now12)
	fmt.Println(now24)
	fmt.Println("🍢 获取时间戳（时间戳：从1970年1月1日到现在的秒数）")
	timez := timeObj.Unix()
	fmt.Printf("以秒为单位的时间戳：%d\n", timez)
	timezNano := timeObj.UnixNano()
	fmt.Printf("以纳秒为单位的时间戳：%d\n", timezNano)
	fmt.Println("🍲 将时间戳转换为可读性强的格式")
	//毫秒单位的时间戳转换
	timek := time.Unix(timez, 0)                 // 先转换成time时间对象。time.Unix()接收两个参数，第一个是毫秒，第二个是纳秒。转换毫秒时将纳秒设置为0，反之亦然
	stime := timek.Format("2006-01-02 15:04:05") //然后将time时间对象格式化
	fmt.Println(stime)
	timeNanok := time.Unix(0, timezNano)
	s2time := timeNanok.Format("2006-01-02 15:04:05")
	fmt.Println(s2time)
	fmt.Println("🍜 将格式化的时间转换为时间戳")
	var times string = "2017/08/17 08:30:00"
	tmp := "2006/01/02 15:04:04"
	rtimez, _ := time.ParseInLocation(tmp, times, time.Local)
	rtime := rtimez.Unix()
	fmt.Println(rtime)
	fmt.Println("🍉 time包中定义的时间间隔常量")
	fmt.Println("1纳秒,", time.Nanosecond)
	fmt.Println("1微秒,", time.Microsecond)
	fmt.Println("1毫秒,", time.Millisecond)
	fmt.Println("1秒,", time.Second)
	fmt.Println("1分,", time.Minute)
	fmt.Println("1小时,", time.Hour)
	fmt.Println("🥝 时间间隔")
	//现在时间的1小时后是什么时间
	after1hour := timeObj.Add(time.Hour)
	fmt.Println("现在时间是：", timeObj.Format("2006-01-02 15:04:05"))
	fmt.Println("一小时后是：", after1hour.Format("2006-01-02 15:04:05"))
	//时间差
	var s1 string = "2017/08/17"
	var s2 string = "2018/05/05"
	tmp2 := "2006/01/02"
	r1time, _ := time.ParseInLocation(tmp2, s1, time.Local)
	r2time, _ := time.ParseInLocation(tmp2, s2, time.Local)
	myLove := r2time.Sub(r1time)
	fmt.Printf("相识：%v 结束：%v 有缘：%v天\n", s1, s2, myLove.Hours()/float64(24))
	fmt.Println("🍔 golang实现定时器")
	fmt.Println("🍔 使用time.NewTicker实现")
	ticker := time.NewTicker(time.Second)
	var num int = 0
	for t := range ticker.C {
		if num < 5 {
			fmt.Println(t)
		} else {
			ticker.Stop() //将ticker从内存中销毁
			break
		}
		num++
	}
	fmt.Println("🍔 使用time.sleep实现")
	fmt.Println(1)
	fmt.Println(2)
	time.Sleep(time.Second * 3)
	fmt.Println(3)
}
