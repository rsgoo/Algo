package main

import (
	"math"
	"fmt"
)

func main() {

	var num int
	fmt.Println("输入查找的数字")
	fmt.Scanln(&num)
	var arr = [...]int{1, 3, 5, 7, 9}
	var low int = 0;
	var top int = len(arr)

	for {
		mid := math.Ceil(float64((low + top) / 2))
		if num == arr[int(mid)] {
			fmt.Println("中点：", mid)
			break
		} else if (num > arr[int((mid))]) {
			low = int(mid) + 1
		} else {
			top = int(mid) - 1
		}
	}

}
