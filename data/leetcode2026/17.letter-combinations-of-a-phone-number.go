/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 00:25:20
 * @link    github.com/taseikyo
 */

func letterCombinations(digits string) []string {
	res := []string{}

	table := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	path := ""
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(digits) {
			res = append(res, path)
			return
		}

		for _, ch := range table[digits[start]] {
			path += string(ch)
			backtrack(start + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
