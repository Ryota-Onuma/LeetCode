package main

import "fmt"

func main() {
	vals1 := []int{2, 7, 11, 15, 0, 0, 0, 0}
	vals2 := []int{1, 2, 3, 4}
	merge(vals1, 4, vals2, 4)
	fmt.Println(vals1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var usedIdx1, usedIdx2 int
	results := make([]int, 0, m+n)
	for i := 0; i < m+n; i++ {
		// 先にnum1のm個の要素を使い切ったとき
		// num2から取るしかない
		if !shouldContinue(usedIdx1, m) && shouldContinue(usedIdx2, n) {
			results = append(results, nums2[usedIdx2])
			usedIdx2++
			continue
		}
		// 先にnum2のn個の要素を使い切ったとき
		// num1から取るしかない
		if shouldContinue(usedIdx1, m) && !shouldContinue(usedIdx2, n) {
			results = append(results, nums1[usedIdx1])
			usedIdx1++
			continue
		}

		if nums1[usedIdx1] <= nums2[usedIdx2] {
			results = append(results, nums1[usedIdx1])
			usedIdx1++
			continue
		}

		if nums1[usedIdx1] > nums2[usedIdx2] {
			results = append(results, nums2[usedIdx2])
			usedIdx2++
			continue
		}
	}
	for i, v := range results {
		nums1[i] = v
	}
}

func shouldContinue(currentIdx, length int) bool {
	return currentIdx < length
}
