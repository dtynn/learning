package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i1, i2 int
	for i1 < m+n && i2 < n {
		if nums2[i2] < nums1[i1] {
			copy(nums1[i1+1:], nums1[i1:])
			nums1[i1] = nums2[i2]
			i2++
		}

		i1++

	}

	if i2 < n {
		copy(nums1[m+i2:], nums2[i2:])
	}
}
