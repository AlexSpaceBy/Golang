// Homework for fifth lecture: to cover the code from lecture 3 with unit tests
// Task 1 - find average of [6]int array
// Task 2 - 1) finds the longest word from the []string slice; 2) returns copy of []int64 slice in reverse order;

package main

import (
	"fmt"
	"unicode/utf8"
)

// Task 1 - average of [6]int array
func average(array [6]int) float64 {

	var sum = 0

	for _, value := range array {
		sum += value
	}

	return float64(sum) / float64(len(array))
}

// Task 2 - the longest word from the slice
func max(slice []string) string {

	var bufferLength = utf8.RuneCountInString(slice[0])
	var bufferString = slice[0]

	for _, value := range slice {

		if utf8.RuneCountInString(value) > bufferLength {

			bufferLength = utf8.RuneCountInString(value)
			bufferString = value
		}
	}

	return bufferString
}

// Task 2 - copy of slice in reverse order
func reverse(slice []int64) []int64 {

	var reverseSlice []int64

	for i := len(slice) - 1; i >= 0; i -= 1 {

		reverseSlice = append(reverseSlice, slice[i])
	}

	return reverseSlice
}

func main() {

	// Task 1 - average of array
	array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(array))

	// Task 2 - the longest word from slice
	sliceString := []string{"one", "two", "one", "three", "four"}
	fmt.Println(max(sliceString))

	// Task 2 - copy of slice in reverse order: original slice is not changed
	sliceInt := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(reverse(sliceInt))
	fmt.Println(sliceInt)

}
