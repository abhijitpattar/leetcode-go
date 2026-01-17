package problems

import "fmt"

/*Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

func TwoSum(nums []int, target int) []int {
	// iterating thorugh entire array
	for i := range nums {
		//fmt.Println("**** i - ", i, ", nums- ", nums[i])
		// start inner loop with one index higher than outerloop
		for j := i + 1; j < len(nums); j++ {
			//fmt.Println("j - ", j, ", nums- ", nums[j])
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSum_optimized(nums []int, target int) []int {
	index := make(map[int]int)
	//{3, 2, 4}, 6
	for ii, num := range nums {
		//fmt.Println("****************num: ", num)
		//fmt.Println(index)
		//fmt.Println("target: ", target)
		diff := target - num
		//fmt.Println("diff: ", diff)
		//fmt.Println("index[diff]: ", index[diff])
		if _, ok := index[diff]; ok {
			//fmt.Println(index, "final")
			return []int{ii, index[diff]}
		}

		index[num] = ii
	}
	return nil
}

func CallTwoSum() {
	retvalue_1 := TwoSum([]int{2, 7, 11, 15}, 9) //[0,1]
	fmt.Println(retvalue_1)

	retvalue_2 := TwoSum([]int{3, 2, 4}, 6) //[1,2]
	fmt.Println(retvalue_2)

	retvalue_3 := TwoSum([]int{3, 3}, 6) //[0,1]
	fmt.Println(retvalue_3)

	retvalue_4 := TwoSum_optimized([]int{3, 2, 4}, 6) //[1,2]
	fmt.Println(retvalue_4)
}
