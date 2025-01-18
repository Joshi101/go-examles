package main

import "fmt"

// Implement array sum by recursion in go lang
func main() {
	nums := []int{2, 5, 2, 3, 4, 5}
	s := sum(nums)
	fmt.Printf("sum of %v is %d", nums, s)
}

func sum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[0] + sum(nums[1:])

}
