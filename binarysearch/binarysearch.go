package binarysearch

import (
	"math/rand"
	"time"
)

// https://leetcode-cn.com/problems/binary-search
func search(nums []int, target int) int {
	// 1. set middle pointer
	// 2. move left or right pointer
	// 3. judge other condition

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == right {
		if nums[left] == target {
			return left
		}
	}

	return -1
}

// https://leetcode-cn.com/problems/first-bad-version/
// time complexity O(logn)
// space complexity O(log1)
func firstBadVersion(n int) int {
	// 1. when exit
	// 2. move left or right
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(n int) bool {
	// min, max := 1, n+1
	num := n + 10
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(num) + 1
	if r == n {
		return true
	}
	return false
}

// https://leetcode-cn.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	tag := left

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			tag = mid
			break
		}

		if nums[mid] < target {
			left = mid + 1
			tag = left
		} else {
			right = mid
		}
	}

	if target > nums[right] {
		tag = right + 1
	}

	return tag
}

// https://leetcode-cn.com/problems/search-a-2d-matrix/submissions
func searchMatrix(matrix [][]int, target int) bool {
	nums := []int{}
	lenout := len(matrix)
	for i := 0; i < lenout; i++ {
		len := len(matrix[i])
		for j := 0; j < len; j++ {
			nums = append(nums, matrix[i][j])
		}
	}
	return searchTrueFalse(nums, target)
}

func searchTrueFalse(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == right {
		if nums[left] == target {
			return true
		}
	}

	return false
}
