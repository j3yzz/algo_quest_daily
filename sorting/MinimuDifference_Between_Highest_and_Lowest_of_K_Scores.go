package main

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	sort.Ints(nums)

	minDiff := 100000000
	for i := 0; i <= len(nums)-k; i++ {
		diff := nums[i+k-1] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func main() {
	nums1 := []int{90}
	k1 := 1
	fmt.Println("Example 1 Output:", minimumDifference(nums1, k1)) // Expected 0

	nums2 := []int{9, 4, 1, 7}
	k2 := 2
	fmt.Println("Example 2 Output:", minimumDifference(nums2, k2)) // Expected 2
}
