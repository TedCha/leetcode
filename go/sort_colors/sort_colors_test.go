package sort_colors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Sort_Colors_Case_1(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	require.Equal(t, []int{0, 0, 1, 1, 2, 2}, nums)
}

func Test_Sort_Colors_Case_2(t *testing.T) {
	nums := []int{2, 0, 1}
	sortColors(nums)
	require.Equal(t, []int{0, 1, 2}, nums)
}

func Test_Sort_Colors_Case_3(t *testing.T) {
	nums := []int{0, 1, 2}
	sortColors(nums)
	require.Equal(t, []int{0, 1, 2}, nums)
}
