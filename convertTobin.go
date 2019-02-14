package main

import (
	"strconv"
	"fmt"
)

func convertToBin(n int) string {
	result := "";
	for ; n > 0; n = n / 2 {
		lowBin := n % 2
		result = strconv.Itoa(lowBin) + result
	}
	return result
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)
}
