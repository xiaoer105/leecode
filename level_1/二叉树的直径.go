package level_1

/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。

路经长度最大值 =节点左子树 + 节点右子树，需要注意的是，结点可以是子节点，并非一定是根节点
遍历每个结点，得到所有节点的最长路径
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

// 每个结点最大值　；　左结点高度　+ 右结点高度
func diameterOfBinaryTree(root *TreeNode) int {

	max = 0

	if nil == root {
		return 0
	}

	getTreeDepth(root)

	return max
}

func getTreeDepth(node *TreeNode) int {
	if nil == node.Left && nil == node.Right {
		return 1
	}

	left := 0
	right := 0

	if nil != node.Left {
		left = getTreeDepth(node.Left)
	}

	if nil != node.Right {
		right = getTreeDepth(node.Right)
	}

	if left+right > max {
		max = left + right
	}

	if left < right {
		return right + 1
	}

	return left + 1
}
