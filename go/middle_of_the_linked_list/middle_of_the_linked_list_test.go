package middle_of_the_linked_list

import (
	"strconv"
	"testing"
)

func Test_Middle_Of_The_Linked_List_Case_1(t *testing.T) {
	l1 := buildList(1, 2, 3, 4, 5)

	r := middleNode(l1)
	s := traverseString(r)
	if s != "345" {
		t.FailNow()
	}
}

func Test_Middle_Of_The_Linked_List_Case_2(t *testing.T) {
	l1 := buildList(1, 2, 3, 4, 5, 6)

	r := middleNode(l1)
	s := traverseString(r)
	if s != "456" {
		t.FailNow()
	}
}

func Test_Middle_Of_The_Linked_List_Case_3(t *testing.T) {
	l1 := buildList()

	r := middleNode(l1)
	s := traverseString(r)
	if s != "" {
		t.FailNow()
	}
}

func buildList(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := new(ListNode)
	tail := l
	for i, d := range nums {
		next := new(ListNode)
		tail.Val = d

		if i+1 != len(nums) {
			tail.Next = next
			tail = next
		}
	}

	return l
}

func traverseString(node *ListNode) string {
	if node == nil {
		return ""
	}

	currVal := strconv.Itoa(node.Val)
	nextVal := traverseString(node.Next)

	return currVal + nextVal
}
