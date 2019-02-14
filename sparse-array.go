package main

import "fmt"

type Node struct {
	row int //行
	col int //列
	val int
}

func main() {

	//1、先创建一个原始数组
	//元素1代码黑棋，2代表白棋
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//2、输出查看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3、转换为稀疏数组
	//思路
	//a：遍历chessMap，如果发现一个元素的值不为0，创建一个Node结构体
	//b：将结构体放在对应的切片中
	var sparseArr []Node

	//4、标准的稀疏数组应该包含一个该数组规模的记录（行，列，值）
	initNode := Node{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, initNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个值节点
				valNode := Node{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println(sparseArr)
	for i, node := range sparseArr {
		fmt.Println("节点", i, node)
	}

	fmt.Println("*******************************")
	//[{11 11 0} {1 2 1} {2 3 2}]
	fmt.Println(sparseArr)

	//稀疏数组恢复
	var originChessMap [11][11]int
	//恢复每一个节点的数组，
	for i, valNode := range sparseArr {
		if i == 0 {
			continue
		}
		originChessMap[valNode.row][valNode.col] = valNode.val
	}
	fmt.Println(originChessMap)
}
