package main

import (
	"fmt"
	"time"
)

//打印1到N
func Recursive(n int) {
	if n > 0 {
		Recursive(n - 1)
		fmt.Println(n)
	}
}

func PrintM(m int) {
	for i := 1; i <= m; i++ {
		fmt.Println(i)
	}
}

func main() {
	begin := time.Now()
	/*
	Recursive(10000) //191.759859ms
	finish := time.Since(begin)
	fmt.Println("递归用时：", finish)
	*/
	///*
	PrintM(10000)		//213.088676ms
	finish := time.Since(begin)
	fmt.Println("循环用时：", finish)
	//*/
}
