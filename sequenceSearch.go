package main

import "fmt"

//顺序查找
func main() {

	var num int
	var flag bool = false
	fmt.Println("输入查找的数字")
	fmt.Scanln(&num)
	arr := [5]int{24, 69, 50, 80, 13}
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			fmt.Println("下标:", i, arr[i])
			flag = true
			break
		}
	}

	if !flag {
		fmt.Println("查找元素", num, "失败")
	}
}
