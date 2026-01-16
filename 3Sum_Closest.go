package main

/*
Given an integer array nums of length n and an integer target, find three integers at distinct indices in nums such that the sum is closest to target.
Return the sum of the three integers.
You may assume that each input would have exactly one solution.



Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
Example 2:

Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
*/

import (
	"fmt"
)

func callthreeSumClosest() {
	int1 := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	fmt.Println("expected 2, received :", int1) // 2

	int2 := threeSumClosest([]int{0, 0, 0}, 1)
	fmt.Println("expected 0, received :", int2)

	int3 := threeSumClosest([]int{10, 20, 30, 40, 50, 60, 70, 80, 90}, 1)
	fmt.Println("expected 60, received :", int3)

	int4 := threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2)
	fmt.Println("expected -2, received :", int4)

	int5 := threeSumClosest([]int{1, 1, 1, 0}, -100)
	fmt.Println("expected 2, received :", int5)

	int6 := threeSumClosest([]int{0, 3, 97, 102, 200}, 300)
	fmt.Println("expected 300, received :", int6)

	int7 := threeSumClosest([]int{408, -583, -665, 709, 729, -390, 850, -718, 419, 987, -374, -548, 526, -733, -186, 478, 381, -240, 279, -539, -949, -19, 621, -251, -72, 874, -227, 612, -891, -43, 106, 775, 153, 676, 971, 362, -744, 215, 986, 781, -977, 325, 229, 236, 885, -922, 100, -29, -76, -57, 622, 883, -576, 266, 37, -832, 316, -46, -764, -569, -930, -366, -669, -102, -469, -552, -230, 114, 139, 422, -307, 895, -348, 359, -235, -792, -808, 123, -627, -493, -441, -756, -735, 866, -762, -429, -758, 633, 539, -101, 174, -799, -781, 142, -709, 776, -708, 294, -70, -498, 133, 390, 515, -915, -274, -564, -850, 444, -214, -244, 341, 560, -600, -621, -147, 516, -296, -636, 658, -510, 773, 832, -457, 436, -74, 250, -109, 772, -866, 86, 805, 175, 533, -864, -863, 881, 295, 975, -535, 578, 957, 548, 744, -921, 83, -588, 45, 33, -937, 61, 183, 536, 574, -425, -677, 204, -461, 84, 172, 2, -681, -925, -515, -750, -361, 322, -31, 743, 634, -971, -194, 484, -451, 936, 465, -761, -69, 513, -20, 431, 137, 774, 528, -965, 277, -662, -738, 690, -325, -816, 503, -167, 809, 163, 912, -406, -879, -272, 338, -715, 940, -18, -354, -459, -846, -673, -395, -534, -444, 562, 352, -233, 267, 371, -215, -837, 851, -258, 717, 891, 208, 931, -411, 878, -90, -953, 91, -299, 903, 962, 561, -356, 915, -670, -424, -367, 984, 287, -591, 835, -155, 918, 132, -262, -231, 103, -934, 646, -135, -917, 797, -198, -528, 138, -123, -126, 69, -173, 252, 487, -452, -972, -163, -828, 779, 576, 583, 333, -268, 412, 289, -27, -815, 712, -716, 56, -948, -156, 304, -386, 470, -989, 207, 616, 593, 112, -513, 580, -398, 264, 906, 278, 405, -447, -801, 618, -181, -113, -954, -956, -501, 421, -336, 877, -975, -810, 388, -202, 558, 698, -712, 360, -488, 367, -668, -587, -696, 506, -686, -415, 749, 177, -913, -263, -769, -984, -995, -939, 220, -14, -162, -741, -92, -580, -10, -237, 128, 235, -892, -22, -284, -603, -966, 73, -499, 750, -688, 281, -413, 541, 519, 988, -865, -300, 245, -674, 707, -383, 344, 32, 424, -343, 393, 715, -772, 452, -248, -265, 575, -189, -99, 339, 93, -532, -704, -687, 923, 467, -768, 425, -403, -529, 741, -908, -905, -464, -472, 365}, 1526)
	fmt.Println("expected 1526, received :", int7)

}

func threeSumClosest(nums []int, target int) int {
	retnum := 0
	// handle case when the input array is less than 3 char
	if len(nums) <= 3 {
		for _, val := range nums {
			retnum += val
		}
		return retnum
	}

	// if the array has more than 3 numbers
	retnum = nums[0] + nums[1] + nums[2]
	total := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				//fmt.Println(i, j, k)
				total = nums[i] + nums[j] + nums[k]
				if AbsInt(total-target) < AbsInt(retnum-target) {
					retnum = total
				}
			}
		}

	}
	return retnum
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
