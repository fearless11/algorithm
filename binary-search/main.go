package main

import "fmt"

// https://leetcode-cn.com/problems/binary-search/

// key point:
// 1. move middle number
// 2. exit

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	fmt.Printf("result: %v\n", search(nums, target))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		mid := left + (right-left)/2

		if left > right {
			return -1
		}

		if left == right {
			if nums[left] == target {
				return left
			} else {
				return -1
			}
		}

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}
}
