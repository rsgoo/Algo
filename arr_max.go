package main

import "fmt"

func main() {

	var intArr = [...]int{1, -2, 23, 5, 89, 11}
	maxVal := intArr[0]
	maxIdx := 0
	for i := 1; i < len(intArr); i++ {
		if intArr[i] > maxVal {
			maxVal = intArr[i]
			maxIdx = i
		}
	}
	fmt.Println(maxVal)
	fmt.Println(maxIdx)
}
