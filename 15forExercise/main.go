package main

import "fmt"

func main() {
	fmt.Println("# ð1.æå°0-50ææçå¶æ°")
	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("# ð2.æå°1+2+3+..+100çå")
	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("# ð3.æå°1-100ä¹é´ææ9çåæ°çåä»¥åä¸ªæ°")
	var sum2 = 0
	var count = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			sum2 += i
			count++
		}
	}
	fmt.Printf("åï¼%vï¼ä¸ªæ°ï¼%v\n", sum2, count)
	fmt.Println("# ð4.è®¡ç®5çé¶ä¹")
	var sum3 = 1
	for i := 1; i <= 5; i++ {
		sum3 *= i
	}
	fmt.Println(sum3)
	fmt.Println("# ð5.æå°ä¸ä¸ªç©å½¢")
	/*
	******
	******
	******
	 */
	var row = 3
	var column = 6
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Print("*") //ä½¿ç¨Printé¿åæ¢è¡
		}
		fmt.Println("") //ä½¿ç¨Printlnæå°ç©ºå­ç¬¦ä¸²ï¼ä½¿å¶å¾ªç¯å®ä¸æ¬¡rowæ¢è¡
	}
	fmt.Println("# ð6.æå°ä¸ä¸ªä¸è§å½¢")
	for i := 0; i < 6; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("# ð7.æå°ä¸ä¸ª99ä¹æ³è¡¨")
	var row2 = 9
	for i := 1; i <= row2; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%vx%v=%v\t", j, i, j*i)
		}
		fmt.Println("")
	}

}
