package sort

import (
	"leetcode/util"
	"testing"
)

// go test -v -run TestBubbleSort -count=1
func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{2, 3, 1, 0},
			want: []int{0, 1, 2, 3},
		},
		{
			nums: []int{3, 1, 0},
			want: []int{0, 1, 3},
		},
	}

	for _, v := range tests {
		nums := make([]int, len(v.nums))
		copy(nums, v.nums)
		got := BubbleSort(v.nums)
		if !util.IsEqual(v.want, got) {
			t.Errorf("BubbleSort(%v) got: %v,wnat: %v", nums, got, v.want)
		}
	}
}

// go test -v -count 1 -run TestSelectSort
func TestSelectSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{5, 3, 2, 1, 7},
			want: []int{1, 2, 3, 5, 7},
		},
		{
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
	}

	for _, v := range tests {
		got := SelectSort(v.nums)
		if !util.IsEqual(got, v.want) {
			t.Errorf("selectSort(%v) got: %v,wnat: %v", v.nums, got, v.want)
		}
	}
}

// go test -v -count 1 -run TestInsertSort
func TestInsertSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{5, 3, 2, 1, 7},
			want: []int{1, 2, 3, 5, 7},
		},
		{
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
	}

	for _, v := range tests {
		got := InsertSort(v.nums)
		if !util.IsEqual(got, v.want) {
			t.Errorf("insertSort(%v) got: %v,wnat: %v", v.nums, got, v.want)
		}
	}
}

// go test -bench=BenchmarkSelectSort -count=3 -benchmem
func BenchmarkSelectSort(b *testing.B) {
	a := util.RandomIntArray(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selectSort(a)
	}
}

// go test -bench=BenchmarkInsertSort -count=3 -benchmem
func BenchmarkInsertSort(b *testing.B) {
	a := util.RandomIntArray(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		insertSort(a)
	}
}
