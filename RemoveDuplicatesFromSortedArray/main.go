package main

func main() {}

func removeDuplicates(nums []int) int {
	record := map[int]struct{}{}
	var k int
	for i, num := range nums {
		if _, ok := record[num]; ok {
			continue
		}
		record[num] = struct{}{}
		nums[k] = nums[i]
		k++
	}

	return k
}
