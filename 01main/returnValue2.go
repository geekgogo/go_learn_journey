package main

import "fmt"

func returnValue2(x, y int, z string) (int, int, string) {
	return x, y, z
}

func demo() {
	a, b, c := returnValue2(1, 2, "hello world!")
	fmt.Println(a, b, c)
}
