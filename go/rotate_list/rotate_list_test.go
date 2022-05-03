package rotate_list

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Rotate_List_Case_1(t *testing.T) {
	l := buildList(1, 2, 3, 4, 5)
	r := rotateRight(l, 2)
	s := traverseString(r)

	require.Equal(t, "4,5,1,2,3", s)
}

func Test_Rotate_List_Case_2(t *testing.T) {
	l := buildList(0, 1, 2)
	r := rotateRight(l, 4)
	s := traverseString(r)

	require.Equal(t, "2,0,1", s)
}
func Test_Rotate_List_Case_3(t *testing.T) {
	l := buildList()
	r := rotateRight(l, 0)
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
