func binarySearch(input []int, target int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		if array[mid] == target {
			// optional early return
		} else if ***comparison*** {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
