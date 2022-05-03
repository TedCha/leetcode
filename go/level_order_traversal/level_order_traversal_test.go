package level_order_traversal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Level_Order_Traversal_Case_1(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	r := levelOrder(tree)
	require.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, r)
}

func Test_Level_Order_Traversal_Case_2(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	r := levelOrder(tree)
	require.Equal(t, [][]int{{1}}, r)
}

func Test_Level_Order_Traversal_Case_3(t *testing.T) {
	var tree *TreeNode

	r := levelOrder(tree)
	require.Equal(t, [][]int{}, r)
}

func Test_Level_Order_Traversal_Case_4(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	r := levelOrder(tree)
	require.Equal(t, [][]int{{1}, {2}}, r)
}

func Test_Level_Order_Traversal_Case_5(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}

	r := levelOrder(tree)
	require.Equal(t, [][]int{{1}, {2, 3}, {4, 5}}, r)
}
