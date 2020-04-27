// The package performs two tasks from homework 2:
//      1) median - find median of sorted array
//      2) structure - find end, perimeter, area of square

package main

import (
	"fmt"
	"sort"
)

// Task 1 - Find median of sorted array
func median(i []int) float64 {

	sort.Ints(i)

	if len(i)%2 != 0 {

		return float64(i[len(i)/2])
	}

	return (float64(i[len(i)/2-1]) + float64(i[len(i)/2])) / 2
}

// Task 2 - Implementation of Point structure
type Point struct {
	x, y int
}

// Task 2 - Implementation of Square structure
type Square struct {
	start Point
	a     uint
}

// Task 2 - Implementation of method for Square structure to find end point
func (s Square) End() Point {

	return Point{s.start.x + int(s.a), s.start.y - int(s.a)}
}

// Task 2 - Implementation of method for Square structure to find perimeter
func (s Square) Perimeter() int {
	return int(s.a * 4)
}

// Task 2 - Implementation of method for Square structure to find area
func (s Square) Area() int {
	return int(s.a * s.a)
}

func main() {

	// Task 1 - median check even
	arrayEven := []int{1, 2, 3, 5, 9, 9, 9, 9}
	fmt.Println(median(arrayEven))

	// Task 1 - median check odd
	arrayOdd := []int{1, 2, 3, 5, 9, 9, 9}
	fmt.Println(median(arrayOdd))

	// Task 2 - structure
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())

}
