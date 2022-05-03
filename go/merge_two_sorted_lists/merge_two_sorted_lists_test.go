package merge_two_sorted_lists

import (
	"strconv"
	"testing"
)

func Test_Merge_Two_Sorted_Lists_Case1(t *testing.T) {
	l1 := buildList(1, 2, 4)
	l2 := buildList(1, 3, 4)

	r := mergeTwoLists(l1, l2)
	s := traverseString(r)
	if s != "112344" {
		t.FailNow()
	}
}

func Test_Merge_Two_Sorted_Lists_Case2(t *testing.T) {
	l1 := buildList()
	l2 := buildList(0)

	r := mergeTwoLists(l1, l2)
	s := traverseString(r)
	if s != "0" {
		t.FailNow()
	}
}

func Test_Merge_Two_Sorted_Lists_Case3(t *testing.T) {
	l1 := buildList(2)
	l2 := buildList(1)

	r := mergeTwoLists(l1, l2)
	s := traverseString(r)
	if s != "12" {
		t.FailNow()
	}
}

func Test_Merge_Two_Sorted_Lists_Case4(t *testing.T) {
	l1 := buildList(-2, 5)
	l2 := buildList(-9, -6, -3, -1, 1, 6)

	r := mergeTwoLists(l1, l2)
	s := traverseString(r)
	if s != "-9-6-3-2-1156" {
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
