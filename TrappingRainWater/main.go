package main

func main() {}

func trapBruteForce(heights []int) int {
	// 各indexの左と右のheightのうち、低い方を記録しておく
	minHighMap := map[int]int{}
	for i := range heights {
		var leftHigh, rightHigh int
		for j := i + 1; j < len(heights); j++ {
			if heights[j] > rightHigh {
				rightHigh = heights[j]
			}
		}

		for k := i - 1; k >= 0; k-- {
			if heights[k] > leftHigh {
				leftHigh = heights[k]
			}
		}
		if leftHigh >= rightHigh {
			minHighMap[i] = rightHigh
		} else {
			minHighMap[i] = leftHigh
		}
	}
	var water int
	for i, h := range heights {
		w := minHighMap[i] - h
		if w > 0 {
			water += w
		}
	}
	return water
}
