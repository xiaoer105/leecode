package level_1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DiameterOfBinaryTree(root *TreeNode) int {
	return diameterOfBinaryTree(root)
}

func diameterOfBinaryTree(root *TreeNode) int {
	leftDepth := 0
	rightDepth := 0

	if nil != root.Left { // 得到左子树的深度
		leftDepth = getTreeDepth(root.Left, 0, 0) + 1
	}

	if nil != root.Right { // 得到左子树的深度
		rightDepth = getTreeDepth(root.Right, 0, 0) + 1
	}

	return leftDepth + rightDepth + 1
}

func getTreeDepth(tree *TreeNode, curr, left, right int) int {
	if nil != tree.Left {
		left++
		getTreeDepth(tree.Left, curr, left, right)
	}

	curr++

	if nil != tree.Right {
		right++
		getTreeDepth(tree.Right, curr, left, right)
	}

	if left == right {
		return left + curr
	} else if left > right {
		return left + curr
	} else {
		return right + curr
	}
}
