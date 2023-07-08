package main

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number
// sorted in non-decreasing order.

//Constraints:
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums is sorted in non-decreasing order.

// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	l, r := 0, len(nums)-1

	for i := len(result) - 1; i >= 0; i-- {
		if abc(nums[l]) > abc(nums[r]) {
			result[i] = nums[l] * nums[l]
			l++
		} else {
			result[i] = nums[r] * nums[r]
			r--
		}
	}

	return result
}

func abc(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
