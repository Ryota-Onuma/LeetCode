package main

func main() {}

func majorityElement(nums []int) int {
	numCountMap := make(map[int]int)
	for _, num := range nums {
		numCountMap[num]++
	}
	var majority, majorityCount int
	for num, count := range numCountMap {
		if count > majorityCount {
			majority = num
			majorityCount = count
		}
	}
	return majority
}

// 空間計算量がO(1)になるように修正
func majorityElement2(nums []int) int {
	var candidate, count int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}
