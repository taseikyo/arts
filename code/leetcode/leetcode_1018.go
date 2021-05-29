/**
 * @date    2021-01-14 21:30:56
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

package main

import "fmt"


func prefixesDivBy5(A []int) []bool {
	num := 0
	ans := make([]bool, len(A))

	for i := 0; i < len(A); i++ {
		num = (num * 2 + A[i]) % 5
		ans[i] = num == 0
	}

	return ans
}

func prefixesDivBy5(A []int) []bool {
	ans := make([]bool, 0, len(A))
	x := 0
	for _, v := range A {
		x = (x << 1 | v) % 5
		ans = append(ans, x == 0)
	}
	return ans
}