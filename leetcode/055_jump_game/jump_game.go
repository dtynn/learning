package jump_game

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	last := len(nums) - 1
	for prev := last - 1; prev >= 0; prev-- {
		distance := nums[prev]
		if distance >= last-prev {
			return canJump(nums[:prev+1])
		}
	}

	return false
}
