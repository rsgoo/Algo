package main

import (
	"fmt"
	"unsafe"
)

type Slice struct {
	ptr unsafe.Pointer // Array pointer
	len int            // slice length
	cap int            // slice capacity
}
//yiw
func main() {
	var a = [][]int{
		{0, 1, 2, 3},   /* 第一行索引为 0 */
		{4, 5, 6, 7},   /* 第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	fmt.Printf("%p\n",(&a))	//0xc000044400
	fmt.Println()
	fmt.Printf("%p---base\n",&a[0]) //0xc000084000---base[540672]
	fmt.Println(a[0]) //0xc000084000---base[540672]
	fmt.Println(unsafe.Sizeof(a[0]))
	fmt.Println(&a[0][0]) //0xc0000480c0
	fmt.Println(&a[0][1]) //0xc0000480c8
	fmt.Println(&a[0][2]) //0xc0000480d0
	fmt.Println(&a[0][3]) //0xc0000480d8
	fmt.Println()

	fmt.Printf("%p---base\n",&a[1]) //0xc000084018---base
	fmt.Println(unsafe.Sizeof(a[1]))
	fmt.Println(&a[1][0]) //0xc0000480e0
	fmt.Println(&a[1][1]) //0xc0000480e8
	fmt.Println(&a[1][2]) //0xc0000480f0
	fmt.Println(&a[1][3]) //0xc0000480f8
	fmt.Println()

	fmt.Printf("%p---base\n",&a[2]) //0xc000084030---base
	fmt.Println(unsafe.Sizeof(a[2]))
	fmt.Println(&a[2][0]) //0xc000048100
	fmt.Println(&a[2][1]) //0xc000048108
	fmt.Println(&a[2][2]) //0xc000048110
	fmt.Println(&a[2][3]) //0xc000048118
	fmt.Println()
}
