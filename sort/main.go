package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Best: O(n) Worse:O(n2)
func BubbleSort(numList []int) []int {
	isSwitch := false
	for i := 0; i <= len(numList)-1; i++ {
		isSwitch = false
		for j := i + 1; j <= len(numList)-1; j++ {
			if numList[i] > numList[j] {
				numList[i], numList[j] = numList[j], numList[i]
				isSwitch = true
			}
		}
		if isSwitch == false {
			return numList
		}
	}

	return numList
}

//Best: O(nlogn) Worse: O(nlogn)
func MergeSort(numList []int) []int {

	if len(numList) < 2 {
		return numList
	}
	midIdx := len(numList) / 2

	left := MergeSort(numList[:midIdx])
	right := MergeSort(numList[midIdx:])
	return Merge(left, right)
}

func Merge(left []int, right []int) []int {

	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func QuickSort(numList []int) []int {
	if len(numList) < 2 {
		return numList
	}

	median := numList[rand.Intn(len(numList))]

	low_part := make([]int, 0, len(numList))
	median_part := make([]int, 0, len(numList))
	high_part := make([]int, 0, len(numList))

	for _, value := range numList {
		switch {
		case value < median:
			low_part = append(low_part, value)
		case value > median:
			high_part = append(high_part, value)
		case value == median:
			median_part = append(median_part, value)
		}
	}

	low_part = QuickSort(low_part)
	high_part = QuickSort(high_part)

	low_part = append(low_part, median_part...)
	low_part = append(low_part, high_part...)

	return low_part

}

func BuildMaxHeap(numList []int) []int {
	heapList := []int{}

	for _, value := range numList {
		heapList = append(heapList, value)
		heapList = MaxHeapifyBottomUp(heapList)
	}

	return heapList
}

func MaxHeapifyBottomUp(heapList []int) []int {
	newValueItem := len(heapList) - 1

	for {
		parantIdx := (newValueItem+1)/2 - 1
		if parantIdx < 0 {
			break
		}

		if heapList[newValueItem] > heapList[parantIdx] {
			heapList[newValueItem], heapList[parantIdx] = heapList[parantIdx], heapList[newValueItem]
			newValueItem = parantIdx
		} else {
			break
		}
	}
	return heapList
}

func HeapSort(numList []int) []int {
	curLastNode := len(numList) - 1
	for {
		if curLastNode <= 0 {
			break
		}
		if curLastNode > 0 {
			numList[0], numList[curLastNode] = numList[curLastNode], numList[0]
			curLastNode--
			MaxHeapifyTopDown(curLastNode, numList)
		}
	}
	return numList
}

func MaxHeapifyTopDown(curLastNode int, numList []int) {
	curEleIdx := 0
	for {
		child := 2*curEleIdx + 1
		if child > curLastNode {
			break
		}
		if child+1 <= curLastNode && numList[child] < numList[child+1] {
			child++
		}
		if numList[child] < numList[curEleIdx] {
			return
		}
		numList[curEleIdx], numList[child] = numList[child], numList[curEleIdx]
		curEleIdx = child
	}
}

func main() {
	numList := make([]int, 10)
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 10; i++ {
		numList[i] = rand.Intn(1000)
	}

	fmt.Println(numList)

	//fmt.Println(BubbleSort(numList))
	//fmt.Println(MergeSort(numList))
	//fmt.Println(QuickSort(numList))

	fmt.Println(HeapSort(BuildMaxHeap(numList)))
}
