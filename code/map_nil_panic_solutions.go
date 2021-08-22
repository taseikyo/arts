/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2021-08-22 22:33:30
 * @link    github.com/taseikyo
 */

package main

import "fmt"

func main() {
	var dict1 map[string]string
	dict2 := map[string]string{}
	dict3 := make(map[string]string)

	// map[] map[] map[]
	fmt.Println(dict1, dict2, dict3)
	// true false false
	fmt.Println(dict1 == nil, dict2 == nil, dict3 == nil)
	// 0x0, 0xc000084150, 0xc000084180
	fmt.Printf("%p, %p, %p\n", dict1, dict2, dict3)
}
