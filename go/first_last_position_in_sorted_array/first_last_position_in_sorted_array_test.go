package first_last_position_in_sorted_array

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_First_Last_Position_In_Sorted_Array_Case_1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	r := searchRange(nums, 8)
	s := intSliceToString(r)
	if s != "3,4" {
		t.FailNow()
	}
}

func Test_First_Last_Position_In_Sorted_Array_Case_2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	r := searchRange(nums, 6)
	s := intSliceToString(r)
	if s != "-1,-1" {
		t.FailNow()
	}
}

func Test_First_Last_Position_In_Sorted_Array_Case_3(t *testing.T) {
	nums := []int{}
	r := searchRange(nums, 6)
	s := intSliceToString(r)
	if s != "-1,-1" {
		t.FailNow()
	}
}

func Test_First_Last_Position_In_Sorted_Array_Case_4(t *testing.T) {
	nums := []int{1}
	r := searchRange(nums, 1)
	s := intSliceToString(r)
	if s != "0,0" {
		t.FailNow()
	}
}

func Test_First_Last_Position_In_Sorted_Array_Case_5(t *testing.T) {
	nums := []int{2, 2}
	r := searchRange(nums, 2)
	s := intSliceToString(r)
	fmt.Println(s)
	if s != "0,1" {
		t.FailNow()
	}
}

func intSliceToString(x []int) string {
	s := make([]string, len(x))
	for i, v := range x {
		s[i] = strconv.Itoa(v)
	}

	return strings.Join(s, ",")
}
