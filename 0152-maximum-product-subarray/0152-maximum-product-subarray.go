

func maxProduct(nums []int) int {
if len(nums) == 0 {
    return 0
}
maxP, minP, result := nums[0], nums[0], nums[0]

for i := 1; i < len(nums); i++ {
    if nums[i] < 0 {
        maxP, minP = minP, maxP
    }
    maxP = max(nums[i], maxP*nums[i])
    minP = min(nums[i], minP*nums[i])

    result = max(result, maxP)
}
return result
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
