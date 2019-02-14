package main

import "fmt"

func main() {
	arr := []int{4, -3, 5, -2, -1, 2, 6, -2}
	ThisSum := 0
	MaxSum := 0
	length := len(arr)

	for i := 0; i < length; i++ {
		ThisSum += arr[i]
		if ThisSum > MaxSum {
			MaxSum = ThisSum
		}
		if ThisSum < 0 {
			ThisSum = 0
		}
	}
	fmt.Println(MaxSum)
}
