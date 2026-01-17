package problems

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums, return all the triplets
[nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.


Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func CallThreeSum() {

	int1 := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println("received :", int1)

	int2 := threeSum([]int{0, 1, 1})
	fmt.Println("received :", int2)

	int3 := threeSum([]int{0, 0, 0})
	fmt.Println("received :", int3)

}

func threeSum(nums []int) [][]int {
	seen := make(map[[3]int]bool)
	// sort in increasing order
	sort.Ints(nums)
	// empty return value
	retslice := [][]int{}
	// get all the combinations of given slice
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				//fmt.Println(i, j, k)
				if (nums[i] + nums[j] + nums[k]) == 0 {

					arrayKey := [3]int{nums[i], nums[j], nums[k]}
					if _, exists := seen[arrayKey]; !exists {
						seen[arrayKey] = true
						retslice = append(retslice, []int{nums[i], nums[j], nums[k]})
					}

				}
			} // end of k loop
		} // end of j loop
	} // end of i loop

	return retslice
}
