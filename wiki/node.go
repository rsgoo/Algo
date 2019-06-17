package main

import "fmt"

func jc(n int) int {
	if n <= 1 {
		return 1
	}
	return n * jc(n-1)
}

func main() {
	for i := 1; i < 16; i++ {
		fmt.Println(i,jc(i))
	}
}
