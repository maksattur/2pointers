package main

import (
	"sort"
)

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.

// Input: nums = [-1,0,1,2,-1,-4] => [-4, -1, -1, 0, 1, 2]
// Output: [[-1,-1,2],[-1,0,1]]

// Сложность O(n^2)
// Память O(1)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	numsSize := len(nums)
	for i := 0; i < numsSize-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		left, right := i+1, numsSize-1

		for left < right {
			sum := nums[left] + nums[right]
			if sum > -a {
				right--
			} else if sum < -a {
				left++
			} else {
				result = append(result, []int{a, nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			}
		}
	}

	return result
}
