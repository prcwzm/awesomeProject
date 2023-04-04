package main

import (
	"../algorithm"
	"fmt"
)

func main() {
	InBackBinaryTree()
}

func InBackBinaryTree() {
	inOrder := []int{4, 2, 1, 5, 3, 6}
	backOrder := []int{4, 2, 5, 6, 3, 1}
	root := algorithm.InBackBinaryTree(inOrder, backOrder)
	fmt.Println("中序遍历结果为：")
	InOrder(root)
	fmt.Println("\n后序遍历结果为：")
	BackOrder(root)
}

func InOrder(root *algorithm.BinaryTree) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Print(root.Val, " ")
	InOrder(root.Right)
}

func BackOrder(root *algorithm.BinaryTree) {
	if root == nil {
		return
	}
	BackOrder(root.Left)
	BackOrder(root.Right)
	fmt.Print(root.Val, " ")
}

func BlankReplace() {
	fmt.Println(algorithm.BlankReplace("sda s s da asd fa s"))

}

func inPostBinaryTree() {
	algorithm.InPostBinaryTree()
}

func repeatCharacter() {
	var value = []int{1, 2, 3, 5, 5}
	fmt.Println(algorithm.RepeatCharacter(value))
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

func TestTwoDimFIndNumber() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	target := 5
	x, y := algorithm.TwoDimFIndNumber(matrix, target)
	if x != 1 || y != 1 {
		fmt.Printf("Expected (1, 1), but got (%d, %d)", x, y)
	}

	fmt.Println(x, y)
	target = 1
	x, y = algorithm.TwoDimFIndNumber(matrix, target)

	fmt.Println(x, y)

	matrix = [][]int{}
	target = 1
	x, y = algorithm.TwoDimFIndNumber(matrix, target)
	if x != -1 || y != -1 {
		fmt.Printf("Expected (-1, -1), but got (%d, %d)", x, y)
	}

	fmt.Println(x, y)

	matrix = [][]int{{}}
	target = 1
	x, y = algorithm.TwoDimFIndNumber(matrix, target)
	if x != -1 || y != -1 {
		fmt.Printf("Expected (-1, -1), but got (%d, %d)", x, y)
	}

	fmt.Println(x, y)
}

func ReverseOutput() {
	testList := []int{1, 2, 3, 4}
	algorithm.ReverseOutput(testList)
}
