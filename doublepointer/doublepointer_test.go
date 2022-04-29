package doublepointer

import (
	"testing"
)

func IsEqual(a, b []int) bool {
	lena := len(a)
	lenb := len(b)
	if lena != lenb {
		return false
	}

	for i := 0; i < lena; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// go test -v -run TestSortedSquares
func TestSortedSquares(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{nums: []int{-7, -3, 2, 3, 11}, want: []int{4, 9, 9, 49, 121}},
		{nums: []int{-7, -6, -5, -4}, want: []int{16, 25, 36, 49}},
	}

	for _, v := range tests {
		got := sortedSquares(v.nums)
		if !IsEqual(got, v.want) {
			t.Errorf("sortedSquares(%v) got: %v,wnat: %v", v.nums, got, v.want)
		}
	}
}

// go test -v -run TestRotate
func TestRotate(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want []int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, want: []int{5, 6, 7, 1, 2, 3, 4}},
	}

	for _, v := range tests {
		rotate(v.nums, v.k)
		got := v.nums
		if !IsEqual(got, v.want) {
			t.Errorf("rotate(%v,%v) got: %v,wnat: %v", v.nums, v.k, got, v.want)
		}
	}
}
