func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := seen[target - nums[i]]; ok {
			return []int{j, i}
		}
		seen[nums[i]] = i
	}
	return []int{}
}
