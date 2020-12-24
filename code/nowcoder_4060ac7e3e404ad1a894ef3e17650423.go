/**
 * @date    2020-12-24 11:31:45
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

package main

import "strings"


func ReplaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1);
}