package merge_sorted_array

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.
The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := len(nums1) - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		tmp := nums1[m]
		tmp2 := nums2[n]
		_, _ = tmp, tmp2
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}
	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
}
