package main

func smallerNumbersThanCurrent(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		tempCount := 0
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] > nums[j] {
				tempCount++
			}
		}
		result = append(result, tempCount)
	}
	return result
}
