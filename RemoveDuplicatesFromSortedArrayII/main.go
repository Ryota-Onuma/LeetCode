package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	fmt.Println(k)
	fmt.Println(nums)
	fmt.Println(nums[:k])
}

func removeDuplicates(nums []int) int {
	var k int
	// 1 <= nums.length <= 3 * 104 なので、[0]はOK
	current := nums[0]
	count := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == current && count >= 2 {
			continue
		}

		if nums[i] == current && i != 0 {
			count++
		}

		if nums[i] != current {
			current = nums[i]
			count = 1
		}

		nums[k] = nums[i]
		k++
	}
	return k
}
