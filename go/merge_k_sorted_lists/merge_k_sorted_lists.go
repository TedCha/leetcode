package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
iteration 1: 1X, 1X, 2  [1, 1, ], min = 1
iteration 2: 4 , 3 , 2X [1, 1, 2, ], min = 2
iteration 3: 4 , 3X, 6  [1, 1, 2, 3, ] min = 3
iteration 4: 4X, 4X, 6  [1, 1, 2, 3, 4, 4, ] min = 4
iteration 5: 5X, O , 6  [1, 1, 2, 3, 4, 4, 5, ] min = 5
iteration 6: O , O , 6X [1, 1, 2, 3, 4, 4, 5, 6] min = 6
iteration 7: O , O , O  [1, 1, 2, 3, 4, 4, 5, 6] min = 6
*/
func mergeKLists(lists []*ListNode) *ListNode {

	// Create our sorted list
	var sortedList *ListNode

	// Make sure that at least one of our lists are populated
	isPopulatedListExist := false
	for i := range lists {
		if lists[i] != nil {
			isPopulatedListExist = true
		}
	}

	// Only init the sortedList if we have a list to process and one of the lists was populated
	if len(lists) > 0 && isPopulatedListExist {
		sortedList = &ListNode{}
	}

	currentNode := sortedList

	// Iterate until all nodes of each list are exhausted
	for {

		// Find the current min for the iteration's nodes
		currentMin := 10001
		for i := range lists {
			if lists[i] == nil {
				continue
			}

			if lists[i].Val < currentMin {
				currentMin = lists[i].Val
			}
		}

		// process the current node of each list
		processedNodeCount := 0
		for i := range lists {
			if lists[i] != nil && lists[i].Val <= currentMin && lists[i].Val != 10001 {

				// update the sorted list
				currentNode.Val = lists[i].Val
				currentNode.Next = new(ListNode)
				currentNode = currentNode.Next

				if lists[i].Next != nil {
					lists[i].Val = lists[i].Next.Val
					lists[i].Next = lists[i].Next.Next
				} else {
					// the next node is nil, means the list is exhausted
					// set to the higher bound to skip
					lists[i].Val = 10001
				}
				processedNodeCount++
			}
		}

		// if no nodes were processed, all lists have been exhausted
		if processedNodeCount == 0 {

			// Traverse the sorted list to set the last value nil
			currSortedNode := sortedList
			for {
				if currSortedNode != nil {
					if currSortedNode.Next != nil && currSortedNode.Next.Next != nil {
						currSortedNode = currSortedNode.Next
					} else {
						currSortedNode.Next = nil
						break
					}
				} else {
					break
				}
			}

			break
		}
	}

	return sortedList
}
