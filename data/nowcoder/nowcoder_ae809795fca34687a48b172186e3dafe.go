/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2020-12-18 00:04:01
 * @link    github.com/taseikyo
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)
	s := strconv.Itoa(num)
	end := len(s) - 1
	for i := end; i >= 0; i-- {
		fmt.Printf("%c", s[i])
	}
}
