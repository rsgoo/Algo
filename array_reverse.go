package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	//生成随机数种子
	rand.Seed(time.Now().UnixNano())
	var randArr [20]int
	for i := 0; i < 20; i++ {
		randArr[i] = rand.Intn(10)
	}
	fmt.Println(randArr)

	//数组反转
	for i:=19; i>=0; i--{
		//19 因为数组第一个元素和最后一个元素索引的下表和是 19
		randArr[i] = randArr[19-i]
	}

	fmt.Println(randArr)
}

