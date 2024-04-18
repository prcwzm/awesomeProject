package main

import (
	"../algorithm"
	"fmt"
)

func main() {
	combineTwoArrays()
}

func combineTwoArrays() {
	array := []int{1, 2, 3, 4, 10}
	a := make([]int, 10, 10)
	copy(a, array)
	b := []int{2, 4, 6, 8, 10}
	algorithm.CombineTwoArrays(a, b, 5, 5)
}

func binaryOne() {
	i := 1000
	fmt.Println(algorithm.BinaryOne(i))
}

func FindLessElement() {
	list := []int{3, 4, 5, 1, 2}
	fmt.Println(algorithm.FindLessElement(list))
}

func FrogDump() {
	fmt.Println(algorithm.FrogDump(2))
}

func CQueue() {
	q := algorithm.CQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop()) // 1
	fmt.Println(q.Pop()) // 2
	q.Push(4)
	fmt.Println(q.Pop()) // 3
	fmt.Println(q.Pop()) // 4
	fmt.Println(q.Pop()) // -1
}

func Fib() {
	algorithm.Fib(5)
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

func StockTime() {
	lit := []float64{0, 1, 2, 3, 4}
	fmt.Println(algorithm.StockProfit(lit))
}

func Jump() {
	list := []float64{2, 2, 1, 0, 1, 2, 3, 4}
	algorithm.Jump(list)
}

func JumpZero() {
	list := []float64{2, 2, 1, 0, 1, 2, 3, 4}
	algorithm.JumpZero(list)
}
