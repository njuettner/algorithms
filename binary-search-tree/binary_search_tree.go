package binary

import (
	"math"
)

// time O(log(n)), space O(log(n))
func BinarySearch(arr []int, target int) bool {
	return binarySearchHelper(arr, target, 0, len(arr)-1)
}

func binarySearchHelper(arr []int, target, left, right int) bool {
	potentialMiddleTarget := int(math.Floor(float64((left + right) / 2)))
	if left > right {
		return false
	}

	if target == arr[potentialMiddleTarget] {
		return true
	}
	if target < arr[potentialMiddleTarget] {
		return binarySearchHelper(arr, target, left, potentialMiddleTarget-1)
	}
	return binarySearchHelper(arr, target, potentialMiddleTarget+1, right)
}
