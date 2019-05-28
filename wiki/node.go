package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var newArr = [4]interface{}{1, 2, 'A', "H"}
	fmt.Printf("%p\n", &newArr)
	fmt.Println("************")
	for i := 0; i < len(newArr); i++ {
		fmt.Println(unsafe.Sizeof(newArr[i]), ":", &newArr[i],"-",newArr[i])
	}
	fmt.Println("----------")
	//var a int32 = 8
	//fmt.Println(unsafe.Sizeof(a))
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%p\n", &arr)
	fmt.Println("************")
	for i := 0; i < len(arr); i++ {
		fmt.Println(&arr[i])
	}
}
