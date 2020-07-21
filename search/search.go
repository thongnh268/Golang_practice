package main

import (
	"fmt"
)

// array := int[]{1, 3 , 4, 6, 7, 12, 34 , 45, 68, 88, 100}

func binarySearch(arraylist []int, lower int, upper int, target int) int {
	if lower > upper {
		return -1
	}
	mid := int(lower + (upper-lower)/2)
	if arraylist[mid] > target {
		return binarySearch(arraylist, lower, mid, target)
	} else if arraylist[mid] < target {
		return binarySearch(arraylist, mid+1, upper, target)
	} else {
		return mid
	}
}

func linearSearch(arraylist []int, target int) int {
	for i, elem := range arraylist {
		if elem == target {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("Binary search:")
	array := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	// index := binarySearch(array, 0, 9, 8)
	index := linearSearch(array, 8)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("array[", index, "] = ", array[index])
	}
}
