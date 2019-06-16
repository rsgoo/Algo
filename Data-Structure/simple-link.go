package main

import "fmt"

type HeroNode struct {
	no       int //编号
	name     string
	nickname string
	next     *HeroNode //指向下一个节点的指针
}

//给链表插入一个节点
//编写一个插入方法，在单链表最后插入
func InsertHeroNode(head, newHeroNode *HeroNode) {
	//思路
	//1：先找到该链表最后的节点
	//2：创建一个辅助节点
	tempNode := head
	for {
		//表示找到最后的节点
		if tempNode.next == nil {
			break
		}
		//让辅助节点不断指向下一个节点
		tempNode = tempNode.next
	}
	//3：将newHeroNode加入到链表的最后
	tempNode.next = newHeroNode
}

//显示列表的所有节点信息
func ListHeroNode(head *HeroNode) {
	//创建辅助节点
	tempNode := head

	//先判断该链表是否一个空链表
	if tempNode.next == nil {
		return
	}
	//遍历节点
	for {
		fmt.Printf("[%d, %s, %s]==> ", tempNode.next.no, tempNode.next.name, tempNode.next.nickname)
		tempNode = tempNode.next
		//判断是否链表
		if tempNode.next == nil {
			break
		}
	}
}

func main() {
	//创建头节点
	head := &HeroNode{}

	//创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	//加入节点
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)

	//显示节点
	ListHeroNode(head)
}
