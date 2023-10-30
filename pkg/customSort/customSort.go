package customSort

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}
