package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// Create list to store rotated list
	rotatedList := new(ListNode)

	// traverse the list to append all of the of values in list order
	var rotationVals []int
	currentNode := head
	for i := 0; ; i++ {
		rotationVals = append(rotationVals, currentNode.Val)
		if currentNode.Next != nil {
			currentNode = currentNode.Next
		} else {
			break
		}
	}

	// create a rotation map; [rotated index]:{value}
	rotationMap := make(map[int]int)
	for i := range rotationVals {
		// current index + transformation % shift = rotated position
		ridx := (i + k) % len(rotationVals)
		rotationMap[ridx] = rotationVals[i]
	}

	currentNode = rotatedList
	for i := range rotationVals {
		// access the map in order to build the rotated list
		currentNode.Val = rotationMap[i]

		// last element doesn't have a next node
		if i != len(rotationVals)-1 {
			currentNode.Next = new(ListNode)
			currentNode = currentNode.Next
		}
	}

	return rotatedList
}
