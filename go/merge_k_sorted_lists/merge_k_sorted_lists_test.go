package merge_k_sorted_lists

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Merge_K_Sorted_Lists_Case_1(t *testing.T) {
	l1 := buildList(1, 4, 5)
	l2 := buildList(1, 3, 4)
	l3 := buildList(2, 6)

	lists := []*ListNode{l1, l2, l3}
	r := mergeKLists(lists)
	s := traverseString(r)

	require.Equal(t, "1,1,2,3,4,4,5,6", s)
}

func Test_Merge_K_Sorted_Lists_Case_2(t *testing.T) {
	lists := []*ListNode{}
	r := mergeKLists(lists)
	s := traverseString(r)

	require.Equal(t, "", s)
}

func Test_Merge_K_Sorted_Lists_Case_3(t *testing.T) {
	var emptyList *ListNode
	lists := []*ListNode{emptyList}
	r := mergeKLists(lists)
	s := traverseString(r)

	require.Equal(t, "", s)
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

	currVal := strconv.Itoa(node.Val) + ","
	nextVal := traverseString(node.Next)

	if nextVal == "" {
		currVal = strings.TrimSuffix(currVal, ",")
	}

	return currVal + nextVal
}
