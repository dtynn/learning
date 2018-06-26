package unique_binary_search_tree_2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	return genNodes(nums)
}

func genNodes(nums []int) []*TreeNode {
	res := []*TreeNode{}
	if len(nums) == 0 {
		return res
	}

	if len(nums) == 1 {
		res = append(res, &TreeNode{
			Val: nums[0],
		})
		return res
	}

	if len(nums) == 2 {
		res = append(res, &TreeNode{
			Val: nums[0],
			Right: &TreeNode{
				Val: nums[1],
			},
		}, &TreeNode{
			Val: nums[1],
			Left: &TreeNode{
				Val: nums[0],
			},
		})

		return res
	}

	for i := 0; i < len(nums); i++ {
		lefts := genNodes(nums[:i])
		if len(lefts) == 0 {
			lefts = append(lefts, nil)
		}

		rights := genNodes(nums[i+1:])
		if len(rights) == 0 {
			rights = append(rights, nil)
		}

		for ileft := 0; ileft < len(lefts); ileft++ {
			for iright := 0; iright < len(rights); iright++ {
				res = append(res, &TreeNode{
					Val:   nums[i],
					Left:  lefts[ileft],
					Right: rights[iright],
				})
			}
		}
	}

	return res
}
