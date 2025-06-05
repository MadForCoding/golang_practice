package main

import (
	"fmt"
)

func main() {
	var constructMode int
	fmt.Println("Pls input the testCase mode(1:manual, 2: direct)")
	fmt.Scanf("%d", &constructMode)
	var testCase []int
	if constructMode == 1 {
		var arrayLen int
		fmt.Println("Pls type data length")
		fmt.Scanf("%d", &arrayLen)
		testCase = constructByCmd(arrayLen)
	} else if constructMode == 2 {
		testCase = constructByDirectTestCase()
	} else {
		fmt.Println("unsupprted testCase")
		return
	}
	fmt.Println("Before sort")
	fmt.Println(testCase)
	fmt.Println("After sort")
	fmt.Println(mergeSort(testCase))
}

func constructByDirectTestCase() []int {
	return []int{3, 5, 38, 20, 18, 1, 2, 10, 8, 6, 4}
}

func constructByCmd(dataLength int) []int {
	var res = make([]int, dataLength)
	for i := 0; i < dataLength; i++ {
		tmp := 0
		fmt.Scanf("%d", &tmp)
		res[i] = tmp
	}
	return res
}

func mergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	middle := len(list) / 2

	return mergeTwoList(mergeSort(list[:middle]), mergeSort(list[middle:]))
}

func mergeTwoList(list1, list2 []int) []int {
	res := make([]int, len(list1)+len(list2))
	index1 := 0
	index2 := 0
	index := 0
	for index1 < len(list1) && index2 < len(list2) {
		if list1[index1] < list2[index2] {
			res[index] = list1[index1]
			index++
			index1++
		} else {
			res[index] = list2[index2]
			index++
			index2++
		}
	}
	if index1 < len(list1) {
		for index1 < len(list1) {
			res[index] = list1[index1]
			index++
			index1++
		}
	}
	if index2 < len(list2) {
		for index2 < len(list2) {
			res[index] = list2[index2]
			index++
			index2++
		}
	}
	return res
}
