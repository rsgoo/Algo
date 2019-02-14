package main

// 24 69 80 57 13
/*外层第一轮*/

//第一次：24 和 69 比较
// 24 69 80 57 13

//第二次 69 和 80 比较
// 24 69 80 57 13

//第三次 80 和 57 比较，交换两者
// 24 69  ==57 80== 13

//第四次 80 和 13 比较，交换两者
// 24 69  57 ==13  80==

// 24 69  57  13  80
/*外层第二轮*/

//第一次: 24 和 69 比较
//24 69  57  13  80

//第二次：69 和 57 比较， 交换
//24  57  69  13  80

//第三次：69 和 13 比较， 交换
//24  57  13  69  80

// 24  57  13  69  80
/*外层第三轮*/

//第一次: 24 和 57 比较
//24 57 13 69 80

//第二次: 57 和 13 比较，交换
//24 13 57 69 80

// 24 13 57 69 80
/*外层第四轮*/

//第一次：24 和 13 比较
// 13 24 57 69 80

func bubble(arr *[5]int) {
	//fmt.Println("排序前：", *arr)
	//数组五个元素 两两比较需要4次
	for i := 0; i < len(arr)-1; i++ {
		//-i 是因为每一次排序后已经确定了一个最值
		for j := 0; j < len(arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j+1], (*arr)[j] = (*arr)[j], (*arr)[j+1]
			}
		}
	}
	//fmt.Println("第一轮比较完后：", *arr)
}

func main() {
	arr := [5]int{24, 69, 50, 80, 13}
	bubble(&arr)
}
