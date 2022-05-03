package level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	levelOrderList := [][]int{}
	if root == nil {
		return levelOrderList
	}

	// init queue with root
	queue := []*TreeNode{root}
	levelOrderList = append(levelOrderList, []int{root.Val})

	for {

		if len(queue) == 0 {
			break
		}

		// pop first value off queue
		currentNode := queue[0]
		queue = queue[1:]

		if currentNode != nil {

			if currentNode.Left != nil && currentNode.Right != nil {
				queue = append(queue, currentNode.Left)
				queue = append(queue, currentNode.Right)
				levelOrderList = append(levelOrderList, []int{currentNode.Left.Val, currentNode.Right.Val})
			} else if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
				levelOrderList = append(levelOrderList, []int{currentNode.Left.Val})
			} else if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
				levelOrderList = append(levelOrderList, []int{currentNode.Right.Val})
			}
		}
	}

	return levelOrderList
}
