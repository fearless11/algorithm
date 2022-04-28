package doublepointer

import "fmt"

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	// diff left and right, get max to tag, then move left and right and tag
	// T(n)=O(n)
	// S(n)=O(n)

	n := len(nums)
	result := make([]int, n)
	i := 0
	j := n - 1
	tag := n - 1

	for i <= j {
		left := nums[i] * nums[i]
		right := nums[j] * nums[j]
		if left > right {
			result[tag] = left
			i++
		} else {
			result[tag] = right
			j--
		}
		tag--
	}
	return result
}

// https://leetcode-cn.com/problems/rotate-array/
func rotate(nums []int, k int) {
	fmt.Println("todo")
}
