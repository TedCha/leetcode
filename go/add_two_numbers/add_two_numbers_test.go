package add_two_numbers

import (
	"strconv"
	"testing"
)

func Test_Add_Two_Numbers(t *testing.T) {
	l1 := buildList(1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1)
	l2 := buildList(5, 6, 4)

	r := addTwoNumbers(l1, l2)
	s := traverse(r)
	if s != "6640000000000000000000000000001" {
		t.FailNow()
	}

}

func traverse(node *ListNode) string {
	if node == nil {
		return ""
	}

	currVal := strconv.Itoa(node.Val)
	nextVal := traverse(node.Next)

	return currVal + nextVal
}

func buildList(nums ...int) *ListNode {
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
