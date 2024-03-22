package easy

//https://leetcode.com/problems/plus-one/

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for digits[i] == 9 {
		digits[i] = 0
		i--
		if i < 0 {
			newDigits := make([]int, 1, len(digits)+1)
			digits = append(newDigits, digits...)
			i = 0
		}
	}
	digits[i]++
	return digits
}
