package main

import "fmt"

//高效打印26个字母

func MyByteArr() [26]byte {
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'a' + byte(i)
	}
	return myChars
}

func main() {

	for _, value := range MyByteArr() {
		v := fmt.Sprintf("%c", value)
		fmt.Println(v)
	}
}
