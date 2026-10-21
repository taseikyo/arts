/*
* @Date:   2025-03-12 23:07:35
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
func addToArrayForm(num []int, k int) []int {
    sum := 0
    for _, val := range num {
        sum = sum*10+val
    }
    sum += k

    fmt.Println(sum)
    res := make([]int, 0)
    for sum > 0 {
        res = append([]int{sum%10}, res...)
        sum /= 10
    }

    return res
}

func addToArrayForm(num []int, k int) []int {
    carry := 0
    idx := len(num) - 1
    for k > 0 && idx >= 0 {
        carry = carry + k%10 + num[idx]
        num[idx] = carry % 10
        carry = carry / 10
        k = k/10
        idx--
    }
    for k > 0 {
        carry = carry + k%10
        num = append([]int{carry % 10}, num...)
        carry = carry / 10
        k = k/10
    }
    for idx >= 0 {
        carry = carry + num[idx]
        num[idx] = carry % 10
        carry = carry / 10
        idx--
    }
    if carry > 0 {
        num = append([]int{1}, num...)
    }


    return num
}
