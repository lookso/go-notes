package main

import (
	"fmt"
)

type LinkNode struct {
	no   int    //编号
	name string //姓名
	next *LinkNode
}

//直接在队尾插入
func InsertNode(head *LinkNode, newNode *LinkNode) {

	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	temp.next = newNode
}

//按编号no从小到大顺序插入
func InsertNode2(head *LinkNode, newNode *LinkNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		} else if temp.next.no >= newNode.no {
			break
		}
		temp = temp.next
	}
	newNode.next = temp.next
	temp.next = newNode
}

//删除队列
func DelNode(head *LinkNode, id int) {
	temp := head

	if temp.next == nil {
		fmt.Println("队列为空！")
		return
	}

	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			temp.next = temp.next.next
			break
		}
		temp = temp.next
	}
}

//显示队列
func ShowNode(head *LinkNode) {
	temp := head
	if temp.next == nil {
		return
	}
	for {
		fmt.Printf("%d:%s\t", temp.next.no, temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	head := &LinkNode{}
	newNode1 := &LinkNode{
		no:   1,
		name: "a",
	}
	newNode2 := &LinkNode{
		no:   2,
		name: "b",
	}
	newNode3 := &LinkNode{
		no:   3,
		name: "c",
	}
	InsertNode2(head, newNode2)
	InsertNode2(head, newNode1)
	InsertNode2(head, newNode3)
	ShowNode(head)
}
