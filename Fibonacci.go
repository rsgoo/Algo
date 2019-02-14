package main

import "fmt"

func Fibonacci(n int) ([]uint64) {
	var result = make([]uint64, n) //声明一个切片
	result[0] = 0
	result[1] = 1
	for i := 2; i < n; i++ {
		result[i] = result[i-2] + result[i-1]
	}
	return result
}

//1	 1  2  3  5  8
//   A  B
//      A  B
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	fb10 := Fibonacci(10)
	fmt.Println(fb10)

	f := fibonacci()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	for i := 1; i < 30; i++ {
		defer fmt.Println(f())
	}
}
