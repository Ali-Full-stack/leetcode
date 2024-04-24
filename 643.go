// Leetcode problems 643 : Find maximum Average in subarray
package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 6, -3, 1, 8, 1, 7}
	k := 4
	// max := FindMaxSub(nums, k)
	max :=FindMaxArrSliding(nums, k)
	fmt.Println(max)

}
func FindMaxArrSliding(nums []int, k int)(max float64){
	total := 0
	max = math.Inf(-1)
	start := 0

	for end:=0; end < len(nums); end++{
		total += nums[end]

		if (end - start) == k - 1{
			max =math.Max(max, float64(total)/float64(k))
			total -= nums[start]
			start++
		}
	}
	return max

}

func FindMaxSub(nums []int, k int) (max float64) {
	max = math.Inf(-1)
	for i := 0; i < len(nums); i++ {
		total := 0
		for j := 0; j < k; j++ {
			if len(nums) > i+(k-1) {
				total += nums[j+i]
			} else {
				i = len(nums)
				j = k
			}
		}
		if total != 0 {
			max = math.Max(max, float64(total)/float64(k))
		}
	}
	return max
}
