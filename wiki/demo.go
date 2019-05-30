package main

import "fmt"

func main() {
	var intArr = [...]int{11, 22, 33, 44, 55, 66}
	fmt.Printf("%p---base\n", &intArr)
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	fmt.Println(&intArr[3])
	fmt.Println(&intArr[4])

	var intSlice = intArr[1:3]
	fmt.Println(intSlice)
	fmt.Printf("%p---base\n", &intSlice)
	fmt.Println(&intSlice[0])
	fmt.Println(&intSlice[1])
	intSlice[0] = 22222
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))
	fmt.Println(intSlice)
	fmt.Println(intArr)
}
