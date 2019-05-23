package main

import (
	"fmt"
	"time"
)

//打印1到N
func PrintN(n int) {
	if n > 0 {
		PrintN(n - 1)
		fmt.Println(n)
	}
}

func PrintM(m int) {
	for i := 1; i <= m; i++ {
		fmt.Println(i)
	}
}

func main() {
	begin := time.Now().Unix()
	PrintN(1000000)      //52
	PrintM(1000000)		//42
	finish := time.Now().Unix()
	fmt.Println("递归用时：", finish-begin)

}
