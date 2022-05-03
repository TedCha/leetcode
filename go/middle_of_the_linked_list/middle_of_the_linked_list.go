package middle_of_the_linked_list

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	depth := depth(head)

	middle := int(math.Ceil((float64(depth) / 2) + 0.5))

	return traverseToDepth(head, middle)
}

func depth(head *ListNode) (depth int) {
	node := head
	for node != nil {
		node = node.Next
		depth++
	}

	return depth
}

func traverseToDepth(head *ListNode, depth int) *ListNode {
	node := head
	index := 1
	for node != nil {
		if index == depth {
			return node
		}
		node = node.Next
		index++
	}

	return nil
}
