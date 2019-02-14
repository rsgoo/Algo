package main

import "fmt"

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	length := len(*q)
	popVal := (*q)[length-1]
	*q = (*q)[:(length - 1)]
	return popVal
}

func main() {
	q := Queue{1}
	q.Push(2)
	//1 2
	fmt.Println(q)
	q.Push(3)
	q.Push(4)
	//1 2 3 4
	fmt.Println(q)
	q.Pop()
	// 1 2 3
	fmt.Println(q)
	q.Pop()
	// 1 2
	fmt.Println(q)

}
