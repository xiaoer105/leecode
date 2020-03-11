package level_1

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   4,
				Left: nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	num := DiameterOfBinaryTree(root)

	t.Log("num:", num)
}
