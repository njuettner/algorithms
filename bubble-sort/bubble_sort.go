package bubble

func BubbleSort(arr []int) []int {
	isSorted := false
	counter := 0
	for !isSorted {
		for i := 0; i < len(arr)-1-counter; i++ {
			isSorted = true
			if arr[i] > arr[i+1] {
				// swap positions
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		counter++
	}
	return arr
}
