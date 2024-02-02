package peak_mountain_array

func peakOfMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid+1] > arr[mid] {
			left = mid + 1
		} else if arr[mid-1] > arr[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
