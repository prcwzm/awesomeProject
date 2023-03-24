package main

import "../algorithm"

func main() {
	reserve()
}

func inPostBinaryTree() {
	algorithm.InPostBinaryTree()
}

func reserve() {
	list := &algorithm.ListNode{Val: 1}
	list1 := &algorithm.ListNode{Val: 2}
	list2 := &algorithm.ListNode{Val: 3}
	list.Next = list1
	list1.Next = list2
	node := algorithm.Reserve(list)
	println(node)
}
