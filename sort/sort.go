package sort

import "log"

//                       Worst       Avg        Best      Space
//Bubble Sort            O(n2)       O(n2)      O(n)       O(1)
//Exchange Sort          O(n2)       O(n2)      O(n2)      O(1)
//Selection Sort         O(n2)       O(n2)      O(n2)      O(1)
//Insertion Sort
//Quick Sort		     O(n2)      O(nlogn)   O(nlogn)    Recur: O(logn)/Iter:O(n)
//Merge Sort            O(nlogn)    O(nlogn)   O(nlogn)    O(n)

func BubbleSort(input []int) (result []int, calRound int) {
	var (
		idx1, idx2 int
		sorted     bool
		calCnt     int
	)
	calCnt = 0
	for idx1 = len(input) - 1; idx1 > 0; idx1-- {

		sorted = true

		for idx2 = 0; idx2 < idx1; idx2++ {
			calCnt++
			log.Println(input)
			if input[idx2] > input[idx2+1] {
				input[idx2], input[idx2+1] = input[idx2+1], input[idx2]
				sorted = false
			}
		}
		log.Println("ROUND: ", input)
		if sorted {
			break
		}
	}

	return input, calCnt
}

func ExchangeSort(input []int) (result []int, calRound int) {
	var (
		idx1, idx2 int
		calCnt     int
	)

	log.Println("Exchange Sort:", input)
	for idx1 = 0; idx1 < len(input)-1; idx1++ {
		for idx2 = idx1 + 1; idx2 < len(input); idx2++ {
			calCnt++
			if input[idx1] > input[idx2] {
				input[idx1], input[idx2] = input[idx2], input[idx1]
			}
			log.Println(input)
		}
		log.Println("Round:", input)
	}

	return input, calCnt
}

func SelectionSort(input []int) (result []int, calRound int) {
	var (
		idx1, idx2 int
		min        int
		calCnt     int
	)

	min = -1
	for idx1 = 0; idx1 < len(input)-1; idx1++ {
		min = idx1
		for idx2 = idx1 + 1; idx2 < len(input); idx2++ {
			if input[min] > input[idx2] {
				min = idx2
			}
		}

		if idx1 != min {
			input[idx1], input[min] = input[min], input[idx1]
		}
		calCnt++
		log.Println(input)
	}

	return input, calCnt
}

func QuickSortRecursive(input []int, startIdx, endIndx int) (result []int, calRound int) {
	var (
		pivot                            int
		start, end                       int
		leftStart, leftEnd, leftLimit    int
		rightStart, rightEnd, rightLimit int
	)

	pivot = input[startIdx]
	start = startIdx + 1
	end = endIndx - 1

	log.Println(input, "pivot", pivot)
	for {
		log.Println("START")
		for end > startIdx && input[end] >= pivot {
			end--
		}
		log.Println(end, ":", input[end])
		for start < end && input[start] <= pivot {
			start++
		}
		log.Println(start, ":", input[start])

		if start < end {
			input[start], input[end] = input[end], input[start]
		} else {
			if end > startIdx {
				input[startIdx], input[end] = input[end], input[startIdx]
				log.Println("over", input)
			}
			break
		}
		log.Println("RESULT:", input)
	}

	leftStart = startIdx
	leftEnd = end
	leftLimit = leftEnd - leftStart

	rightStart = end + 1
	rightEnd = endIndx
	rightLimit = rightEnd - rightStart

	if leftLimit > 1 {
		QuickSortRecursive(input, leftStart, leftEnd)
	}

	if rightLimit > 1 {
		QuickSortRecursive(input, rightStart, rightEnd)
	}

	return input, 0

}

func MergeSort(input []int) []int {
	var (
		midPos int
	)

	if len(input) < 2 {
		return input
	}

	midPos = len(input) / 2
	return Merge(MergeSort(input[:midPos]), MergeSort(input[midPos:]))
}
func Merge(left, right []int) []int {
	var (
		leftIdx, rightIdx, tmpIdx, bufferSize int
		buffer                                []int
	)

	bufferSize = len(left) + len(right)
	buffer = make([]int, bufferSize)

	for tmpIdx = 0; tmpIdx < bufferSize; tmpIdx++ {
		if leftIdx > len(left)-1 && rightIdx <= len(right)-1 {
			buffer[tmpIdx] = right[rightIdx]
			rightIdx++
		} else if rightIdx > len(right)-1 && leftIdx <= len(left)-1 {
			buffer[tmpIdx] = left[leftIdx]
			leftIdx++
		} else if left[leftIdx] < right[rightIdx] {
			buffer[tmpIdx] = left[leftIdx]
			leftIdx++
		} else {
			buffer[tmpIdx] = right[rightIdx]
			rightIdx++
		}
	}
	return buffer

}
