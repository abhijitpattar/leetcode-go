package problems

import "fmt"

/*Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
Return the array in the form [x1,y1,x2,y2,...,xn,yn].*/

func Callshuffle() {

	ret1 := shuffle([]int{1, 2, 3, 4}, 2)
	fmt.Println("returned value - ", ret1)

	ret2 := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	fmt.Println("returned value - ", ret2)

	ret3 := shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4)
	fmt.Println("returned value - ", ret3)

	ret4 := shuffle([]int{1, 1, 2, 2}, 2)
	fmt.Println("returned value - ", ret4)

}

func shuffle(nums []int, n int) []int {
	retval := []int{}
	for i := range n {
		retval = append(retval, nums[i], nums[n+i])
	}
	return retval
}
