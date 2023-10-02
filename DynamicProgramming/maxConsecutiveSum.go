// Here, we are finding maximum consecutive sum value for a substring, when we can't change the order
// using dynamic programming, we are finding the max sum upto the currect index and then storing it
// it has the timeComplexicity of O(n) and space complexity of O(1)

// In this code, the maxSubarraySum function finds the maximum sum of any consecutive numbers in the array. 
// It maintains two variables, maxSum and currentSum, where currentSum represents the sum of the current subarray,
// and maxSum keeps track of the maximum sum found so far. The loop iterates through the array and updates these variables accordingly.

package main

import (
	"fmt"
)

func maxSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubarraySum(nums)
	fmt.Printf("Maximum sum of any consecutive numbers: %d\n", result)
}
