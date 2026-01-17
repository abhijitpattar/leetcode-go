package problems

import (
	"fmt"
)

/*
You have a set of integers s, which originally contains all the numbers from 1 to n.
Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set,
which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.

Example 1:

Input: nums = [1,2,2,4]
Output: [2,3]
Example 2:

Input: nums = [1,1]
Output: [1,2]
*/

func CallFindErrorNums() {
	retvalue_1 := findErrorNums([]int{1, 2, 2, 4}) // [2,3]
	fmt.Println("returned value -", retvalue_1)

	retvalue_2 := findErrorNums([]int{1, 1}) //[1,2]
	fmt.Println("returned value -", retvalue_2)

	retvalue_3 := findErrorNums([]int{2, 2}) //[2, 1]
	fmt.Println("returned value -", retvalue_3)

	retvalue_4 := findErrorNums([]int{3, 2, 3, 4, 6, 5}) //[3, 1]
	fmt.Println("returned value -", retvalue_4)

	retvalue_5 := findErrorNums([]int{3, 2, 2}) //[2, 1]
	fmt.Println("returned value -", retvalue_5)

	retvalue_6 := findErrorNums([]int{3, 3, 2}) //[3, 1]
	fmt.Println("returned value -", retvalue_6)
}

func findErrorNums(nums []int) []int {

	seenmap := make(map[int]bool, len(nums))
	duplicate, missing := 0, 0

	for _, val := range nums {
		if !seenmap[val] {
			seenmap[val] = true
		} else {
			duplicate = val
		}
	}

	for i := 1; i <= len(nums); i++ {
		if !seenmap[i] {
			missing = i
			break
		}
	}

	return []int{duplicate, missing}
}
