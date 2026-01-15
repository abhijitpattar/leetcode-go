package main

import "fmt"

func callBubbleSort() {
	unList1 := []int{3, 7, 8, 2, 5}
	bubbleSort(&unList1)
	fmt.Println("ordered List1 :", unList1)

	unList2 := []int{9834, 9832, 93, 934, 234, 4582}
	bubbleSort(&unList2)
	fmt.Println("ordered List2 :", unList2)
}

func bubbleSort(unList *[]int) {
	deUnList := *unList
	//fmt.Println(deUnList)
	for range deUnList {
		//fmt.Println(deUnList[i])
		for j := range deUnList {
			if (j - 1) >= 0 {
				if deUnList[j-1] > deUnList[j] {
					// swap the numbers
					i := deUnList[j-1]
					deUnList[j-1] = deUnList[j]
					deUnList[j] = i
				}
			}
		}
	}
}
