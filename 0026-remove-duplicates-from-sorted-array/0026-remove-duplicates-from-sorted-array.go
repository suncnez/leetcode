func removeDuplicates(nums []int) int {
    length := len(nums)

    if length <= 1 {
        return length
    }
i := 0

for j := 1; j < length; j++ {
    if nums[i] != nums[j] {
        i++
        nums[i] = nums[j]
    }
}
return i + 1
}