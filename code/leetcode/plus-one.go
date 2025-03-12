/*
* @Date:   2025-03-12 23:16:01
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0 ; i-- {
        digits[i] += 1
        if digits[i] < 10 {
            return digits
        }
        
        digits[i] = digits[i]%10
        if i == 0 {
            digits = append([]int{1}, digits...)
        }
    }

    return digits
}
