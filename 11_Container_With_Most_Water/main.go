package main

//You are given an integer array height of length n.
//There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

//Find two lines that together with the x-axis form a container,
//such that the container contains the most water.
//Return the maximum amount of water a container can store.
//Notice that you may not slant the container.

func maxArea(height []int) int {
	maxArea := 0
	leftPointer, rightPointer := 0, len(height)-1

	for leftPointer < rightPointer {
		base := rightPointer - leftPointer
		if height[leftPointer] < height[rightPointer] {
			base *= height[leftPointer]
			leftPointer++
		} else {
			base *= height[rightPointer]
			rightPointer--
		}

		if maxArea < base {
			maxArea = base
		}
	}

	return maxArea
}
