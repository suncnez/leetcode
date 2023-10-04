func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
if idx, ok := m[target - nums[i]]; ok {
    return []int{i, idx}
}
m[nums[i]] = i
}
return []int{-1, -1}
}
