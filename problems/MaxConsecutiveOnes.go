package problems

import "fmt"

/*
Given a binary array nums, return the maximum number of consecutive 1's in the array.

Example 1:

Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
The maximum number of consecutive 1s is 3.

Example 2:
Input: nums = [1,0,1,1,0,1]
Output: 2
*/

func CallFindMaxConsecutiveOnes() {
	retvalue_1 := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}) //3
	fmt.Println("returned value -", retvalue_1)

	retvalue_2 := findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}) //2
	fmt.Println("returned value -", retvalue_2)

}

func findMaxConsecutiveOnes(nums []int) int {
	retval := 0
	counter := 0
	for i := range nums {

		if nums[i] == 1 {
			counter++
			if counter > retval {
				retval = counter
			}
		}
		if nums[i] == 0 {
			counter = 0
		}
	}
	return retval
}
