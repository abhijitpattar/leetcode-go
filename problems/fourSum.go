package problems

import "fmt"

/*
Given an array nums of n integers,
return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]]
such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]] */

func CallFourSum() {

	int1 := threeSumClosest([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println("expected 2, received :", int1)

	int2 := threeSumClosest([]int{2, 2, 2, 2, 2}, 8)
	fmt.Println("expected 0, received :", int2)

}

func fourSum(nums []int, target int) [][]int {
	retslice := [][]int{}
	return retslice
}
