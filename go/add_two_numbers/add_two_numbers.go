package add_two_numbers

import (
	"fmt"
	"math/big"
	"strconv"
	"unicode/utf8"
)

// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// traverse first list and convert to bigint
	r1 := reverseTraverse(l1)
	n1 := new(big.Int)
	n1, _ = n1.SetString(r1, 10)

	// traverse second list and convert to int
	r2 := reverseTraverse(l2)
	n2 := new(big.Int)
	n2, _ = n2.SetString(r2, 10)

	// sum, convert back to string, and reverse
	s := n1.Add(n1, n2)
	rstr := reverse(s.String())

	l := new(ListNode)
	tail := l
	for i, c := range rstr {
		next := new(ListNode)
		tail.Val, _ = strconv.Atoi(fmt.Sprintf("%c", c))

		if i+1 != len(rstr) {
			tail.Next = next
			tail = next
		}
	}

	return l
}

func reverseTraverse(node *ListNode) string {
	if node == nil {
		return ""
	}

	currVal := strconv.Itoa(node.Val)
	nextVal := reverseTraverse(node.Next)

	return nextVal + currVal
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for i := 0; i < size; {
		r, n := utf8.DecodeRuneInString(s[i:])
		i += n
		utf8.EncodeRune(buf[size-i:], r)
	}
	return string(buf)
}
