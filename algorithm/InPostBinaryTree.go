package algorithm

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func InPostBinaryTree() {
	inorder := []int{2, 3, 4, 5, 6, 8, 9}
	/*5, 3, 2, 4, 8, 6, 9*/
	/*2, 3, 4, 5, 6, 8, 9*/
	postorder := []int{5, 3, 2, 4, 8, 6, 9}
	binaryTree := tree(postorder, inorder)
	printTree(binaryTree)
}

func printTree(tree *BinaryTree) {
	if tree == nil {
		return
	}
	println(tree.Val)
	printTree(tree.Left)
	printTree(tree.Right)
}

func tree(postorder []int, inorder []int) (treeNode *BinaryTree) {
	if len(postorder) == 0 {
		return nil
	}
	treeNode = &BinaryTree{Val: postorder[0]}
	index := 0
	for i, val := range inorder {
		if val == postorder[0] {
			index = i
			break
		}
	}

	treeNode.Left = tree(postorder[1:index+1], inorder[:index])
	treeNode.Right = tree(postorder[index+1:], inorder[index+1:])
	return
}
