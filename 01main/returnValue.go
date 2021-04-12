package main

import "fmt"

func returnValue(x, y int, z string) (int, int, string) {
	return x, y, z
}

func main() {
	a, b, c := returnValue(1, 2, "hello world!")
	fmt.Println(a, b, c)
}
