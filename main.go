package main

import (
	"fmt"
	"leetcode/binarysearch"
)

func main() {
	fmt.Println("start leetcode")

	// example
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	fmt.Printf("result: %v\n", binarysearch.Search(nums, target))
}
