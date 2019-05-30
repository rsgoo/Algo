package main

import "fmt"

//在一个切片中获取指定值得下标

func FindTarget(mySlice []int, target int) (int, int) {
	sliceLen := len(mySlice)
	for i := 0; i < sliceLen; i++ {
		if mySlice[i] == target {
			return i, mySlice[i]
		}
	}
	return -1, -1
}

func main() {
	var mySlice = []int{1, 0, -10, 67, 19, 111, 19}
	index, value := FindTarget(mySlice, 19)
	fmt.Println("index: ", index)
	fmt.Println("value: ", value)
}
