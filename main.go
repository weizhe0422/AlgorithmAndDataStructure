package main

import (
	"fmt"
	tree2 "github.com/weizhe0422/algorithm/tree"

)

func main() {
	var (
		//unsorted []int
		tree *tree2.Tree
	)

	//unsorted = []int{4,2,7,9,10,1,5,3}

	//fmt.Println("Bubble Sort:")
	//fmt.Println(sort.BubbleSort(unsorted))

	//fmt.Println("Exchange Sort:")
	//fmt.Println(sort.ExchangeSort(unsorted))

	//fmt.Println("Selection Sort:")
	//fmt.Println(sort.SelectionSort(unsorted))

	//fmt.Println("Quick Sort (Recursive):")
	//fmt.Println(sort.QuickSortRecursive(unsorted,0,len(unsorted)))

	//fmt.Println("Merge Sort (Recursive):")
	//fmt.Println(sort.MergeSort(unsorted))

	//	     4
	//	  2      7
	//	1   3  5   9
	//	             10
	tree = &tree2.Tree{}
	tree.Inert(4)
	tree.Inert(2)
	tree.Inert(7)
	tree.Inert(9)
	tree.Inert(10)
	tree.Inert(1)
	tree.Inert(5)
	tree.Inert(3)

	//fmt.Println("PreOrder", tree.PreOrder())
	fmt.Println("InOrder", tree.InOrder())

}
