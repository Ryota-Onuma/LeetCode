package main

func main() {}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, startIdx, endIdx int) {
	var tmp int
	for i := startIdx; i < endIdx; i++ {
		tmp = nums[endIdx]
		nums[endIdx] = nums[startIdx]
		nums[startIdx] = tmp
		startIdx++
		endIdx--
	}
}
