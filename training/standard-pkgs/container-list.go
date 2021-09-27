package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("========== container/list ==============")
	mlist := list.New()
	// 尾插
	mlist.PushBack(5)
	mlist.PushBack(6)
	mlist.PushBack(7)
	// 头插
	mlist.PushFront(3)
	mlist.PushFront(2)
	mlist.PushFront(1)
	printList(mlist, "init") // 1 2 3 5 6 7
	// 返回链表第一个元素
	fmt.Println("Front:", mlist.Front().Value) // Front: 1
	// 返回链表最后一个元素
	fmt.Println("Back:", mlist.Back().Value) // Back: 7
	// Remove: 移除某个元素
	mlist.Remove(mlist.Front())
	mlist.Remove(mlist.Front())
	mlist.Remove(mlist.Front())
	printList(mlist, "Remove") // 5 6 7
	mlist.InsertAfter(111, mlist.Front()) // 5 111 6 7
	printList(mlist, "InsertAfter")
	mlist.MoveAfter(mlist.Front().Next(), mlist.Back())
	printList(mlist, "MoveAfter") // 5 6 7 111
	mlist.MoveToFront(mlist.Back())
	printList(mlist, "MoveToFront") // 111 5 6 7
	mlist.Remove(mlist.Front())
	mlist.InsertBefore(222, mlist.Back())
	printList(mlist, "InsertBefore") // 5 6 222 7
	mlist.MoveToBack(mlist.Back().Prev())
	printList(mlist, "MoveToBack") // 5 6 7 222
}

func printList(mlist *list.List, prefix string)  {
	for curhead := mlist.Front(); curhead != nil; curhead = curhead.Next() {
		fmt.Println(prefix, ":", curhead.Value)
	}
}