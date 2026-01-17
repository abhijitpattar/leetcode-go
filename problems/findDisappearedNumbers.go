package problems

import (
	"fmt"
	"sort"
)

func CallfindDisappearedNumbers() {

	retvalue_1 := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println("expected [5,6], returned value - ", retvalue_1)

	retvalue_2 := findDisappearedNumbers([]int{1, 1})
	fmt.Println("expected [2], returned value -", retvalue_2)

}
func findDisappearedNumbers(nums []int) []int {
	retval := []int{}

	sort.Ints(nums)
	fmt.Println(nums)
	for i := range nums {
		if nums[i] != i+1 {
			retval = append(retval, i+1)
		}
	}
	return retval
}
