package main

import "fmt"

// this will be the main function from where all the leetcode calls will happen
func main() {
	CallTwoSum()
}

func CallTwoSum() {
	// retvalue_1 := TwoSum([]int{2, 7, 11, 15}, 9) //[0,1]
	// fmt.Println(retvalue_1)

	// retvalue_2 := TwoSum([]int{3, 2, 4}, 6) //[1,2]
	// fmt.Println(retvalue_2)

	// retvalue_3 := TwoSum([]int{3, 3}, 6) //[0,1]
	// fmt.Println(retvalue_3)

	retvalue_4 := TwoSum_optimized([]int{3, 2, 4}, 6) //[1,2]
	fmt.Println(retvalue_4)
}
