package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	switch {
	case list1 == nil:
		return list2
	case list2 == nil:
		return list1
	}

	// Default ordering
	main := list1
	sub := list2

	// Switch to avoid reverse traversal
	if list1.Val > list2.Val {
		main = list2
		sub = list1
	}

	traverseMerge(main, sub)
	return main
}

func traverseMerge(main *ListNode, node *ListNode) {
	// if node is nil, end recursion
	if node == nil {
		return
	}

	// Merge the current node into the main list
	insertSort(main, node.Val)

	// Continue the traversal
	traverseMerge(main, node.Next)
}

func insertSort(node *ListNode, n int) {
	switch {
	// if the next node is nil, set the next value as the tail
	case node.Next == nil:
		node.Next = &ListNode{Val: n}
	// if the value is >= to the nodes val and less than the next nodes val
	// insert the new node between the current node and the next node
	case node.Val <= n && n < node.Next.Val:
		newNode := &ListNode{Val: n, Next: node.Next}
		node.Next = newNode
	// else, continue the traversal
	default:
		insertSort(node.Next, n)
	}
}
