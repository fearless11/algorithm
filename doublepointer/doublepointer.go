package doublepointer

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
	// 输入: nums = [1,2,3,4,5,6,7], k = 3
	// 输出: [5,6,7,1,2,3,4]
	// 解释:
	// 向右轮转 1 步: [7,1,2,3,4,5,6]
	// 向右轮转 2 步: [6,7,1,2,3,4,5]
	// 向右轮转 3 步: [5,6,7,1,2,3,4]
	rotate2(nums, k)
}

// T(n)=O(n)  S(n)=O(n)
func rotate1(nums []int, k int) {
	n := len(nums)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		idx := (i + k) % n
		temp[idx] = nums[i]
	}
	copy(nums, temp)
}

// T(n)=O(n) S(n)=O(1)
func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
