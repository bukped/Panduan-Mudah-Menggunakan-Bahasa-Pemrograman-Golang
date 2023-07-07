package main

import "testing"

func CountElements(slice []int) int {
	count := 0
	for range slice {
		count++
	}
	return count
}

func TestCountElements(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expectedCount := 5

	actualCount := CountElements(slice)

	if actualCount != expectedCount {
		t.Errorf("Expected count: %d, but got: %d", expectedCount, actualCount)
	}
}
