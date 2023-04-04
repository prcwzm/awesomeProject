package algorithm

func InBackBinaryTree(inOrder []int, backOrder []int) *BinaryTree {
	if len(backOrder) == 0 {
		return nil
	}
	treeNode := &BinaryTree{Val: backOrder[len(backOrder)-1]}
	index := 0
	for ; index < len(inOrder); index++ {
		if inOrder[index] == treeNode.Val {
			break
		}
	}

	treeNode.Left = InBackBinaryTree(inOrder[:index], backOrder[:index])
	treeNode.Right = InBackBinaryTree(inOrder[index+1:], backOrder[index:len(backOrder)-1])
	return treeNode
}
