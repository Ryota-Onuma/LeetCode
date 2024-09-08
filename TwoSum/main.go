package main

import "fmt"

func main() {
	results := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(results)
}

type idx = int

func twoSum(nums []int, target int) []int {
	numAndIndices := map[idx][]idx{}
	for i, num := range nums {
		numAndIndices[num] = append(numAndIndices[num], i)
	}

	for i, num := range nums {
		wantSecondVal := target - num
		indices, ok := numAndIndices[wantSecondVal]
		if !ok {
			continue
		}
		for _, index := range indices {
			if i == index {
				continue
			}
			return []int{i, index}
		}
	}

	return []int{}
}
