package MedianOfTwoSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	all := len(nums1) + len(nums2)
	needed := all/2 + 1
	// 归并排序的最后一步合并, 优化之后我们只需要前 needed 个元素
	merged := make([]int, 0, all)
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] < nums2[0] {
			merged = append(merged, nums1[0])
			nums1 = nums1[1:]
		} else {
			merged = append(merged, nums2[0])
			nums2 = nums2[1:]
		}

		if len(merged) == needed {
			goto OUTPUT
		}
	}

	// nums1 或 nums2 中的至少一个已经用尽, 但是不一定获得需要的所有元素
	if len(merged) < needed {
		if len(nums1) > 0 {
			merged = append(merged, nums1[:needed-len(merged)]...)
		} else if len(nums2) > 0 {
			merged = append(merged, nums2[:needed-len(merged)]...)
		}
	}

OUTPUT:
	if all&1 == 0 {
		return float64(merged[len(merged)-2]+merged[len(merged)-1]) / 2
	}

	return float64(merged[len(merged)-1])
}
