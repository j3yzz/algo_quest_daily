package main

import "fmt"

func main() {
	var nums = []int{-1, 1, 2, 3, 1}

	fmt.Printf("result: %v\n", countPairs(nums, 2))
}

func countPairs(nums []int, target int) int {
	var count = 0

	for i := 0; i < len(nums)-1; i++ {
		for j := 1 + i; j < len(nums); j++ {
			if (nums[i] + nums[j]) < target {
				count++
			}
		}
	}

	return count
}
