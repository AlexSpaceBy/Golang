// This is for testing the code from hometask5 file

package main

import "testing"

// Testing of 'average' function
type testAverage struct {
	values  [6]int
	average float64
}

// List of test cases to cover possible outcomes
var casesAverage = []testAverage{
	{[6]int{1, 2, 3, 4, 5, 6}, 3.5},
	{[6]int{0, 0, 3, 4, 5, 12}, 4.0},
	{[6]int{0, 0, 0, 0, 0, 60}, 10.0},
}

// Test function for 'average' function
func TestAverage(t *testing.T) {
	for _, value := range casesAverage {
		result := average(value.values)
		if result != value.average {
			t.Errorf("average(%q) = %v", value.values, result)
		}
	}
}

// Testing of 'max' function
type testMax struct {
	values []string
	max    string
}

// List of cases to cover possible outcomes
var casesMax = []testMax{
	{[]string{"one", "two", "three"}, "three"},
	{[]string{"test", "tests", "thelongestwordintheslice", "four", "five"}, "thelongestwordintheslice"},
	{[]string{" ", "test", " ", " ", " "}, "test"},
}

// Test function for 'max' function
func TestMax(t *testing.T) {
	for _, value := range casesMax {
		result := max(value.values)
		if result != value.max {
			t.Errorf("average(%q) = %v", value.values, result)
		}
	}
}

// Testing of 'reverse' function
type testReverse struct {
	values  []int64
	reverse []int64
}

// List of cases to cover possible outcomes
var casesReverse = []testReverse{
	{[]int64{2, 3, 4, 5}, []int64{5, 4, 3, 2}},
	{[]int64{2, 3, 2}, []int64{2, 3, 2}},
	{[]int64{0, 0, 0, 0, 0}, []int64{0, 0, 0, 0, 0}},
}

// Test function for 'reverse' function
func TestReverse(t *testing.T) {
	for _, value := range casesReverse {
		result := reverse(value.values)
		for ind, val := range result {
			if val != value.reverse[ind] {
				t.Errorf("average(%v) = %v", value.values, result)
			}
		}
	}
}
