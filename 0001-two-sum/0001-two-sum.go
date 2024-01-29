func twoSum(nums []int, target int) []int {
    result := []int{}

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if target == nums[i] + nums[j] {
                result = []int{i, j}
                return result
            }
        }
    }

    return result
}