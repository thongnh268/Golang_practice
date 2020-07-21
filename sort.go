package main

import (
	"fmt"
	"math/rand"
)

//Merge sort
func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	var middle = len(array) / 2
	var a = mergeSort(array[:middle])
	var b = mergeSort(array[middle:])
	return merge(a, b)
	// merge()
}

func merge(a []int, b []int) []int {
	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}
	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}
	return r
}

//Quick sort
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

//Insertion sort
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		in := i
		for ; in > 0 && arr[in-1] >= temp; in-- {
			arr[in] = arr[in-1]
		}
		arr[in] = temp
	}
	return arr
}

//Selection sort
func selectionSort(arr []int) []int {
	var min int
	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
	return arr
}

//Bubble sort
func bubbleSort(arr []int) []int {
	for i := (len(arr) - 1); i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if arr[j-1] > arr[j] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func main() {
	array := []int{4, 7, 1, 0, 13, 43, 54, 67, 87, 10, 23, 25, 56}
	// var a []int = mergeSort(array)
	// var a []int = quickSort(array)
	// var a []int = insertionSort(array)
	// var a []int = selectionSort(array)
	var a []int = bubbleSort(array)
	fmt.Println("Sorted Array:", a)

}
