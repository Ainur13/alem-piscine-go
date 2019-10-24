package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var nums []int
	if n == 0 {
		nums = append(nums, n)
	}
	l := 0
	for n > 0 {
		nums = append(nums, n%10)
		n = n / 10
		l++
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				t := nums[i]
				nums[i] = nums[j]
				nums[j] = t
			}
		}
	}
	for i := range nums {
		z01.PrintRune(rune(48 + nums[i]))
	}
}
