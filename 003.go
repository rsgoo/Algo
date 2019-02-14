package main

import "fmt"

//金字塔打印
//  *
// ***
//*****

func main() {
	var num = 5
	//层数
	for i := 1; i <= num; i++ {
		//前边空格数
		for j := 1; j <= num-i; j++ {
			fmt.Print(" ")
		}

		//*数
		for k := 1; k <= 2*i-1; k++ {
			//第一个和最后一个打印*号，最后一行全部打印
			if k == 1 || k == 2*i-1 || i == num {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
