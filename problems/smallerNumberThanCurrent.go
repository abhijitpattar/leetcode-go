package problems

import "fmt"

func CallsmallerNumbersThanCurrent() {
	retvalue_1 := smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3})
	fmt.Println("expected [4,0,1,1,3], returned value - ", retvalue_1)

	retvalue_2 := smallerNumbersThanCurrent([]int{6, 5, 4, 8})
	fmt.Println("expected [2,1,0,3], returned value -", retvalue_2)

	retvalue_3 := smallerNumbersThanCurrent([]int{7, 7, 7, 7})
	fmt.Println("expected [0,0,0,0], returned value -", retvalue_3)

}

func smallerNumbersThanCurrent(nums []int) []int {
	retval := []int{}

	for i := range nums {
		count := 0
		for j := range nums {
			if i != j && nums[j] < nums[i] {
				count++
			}
		}
		retval = append(retval, count)
	}
	return retval
}
