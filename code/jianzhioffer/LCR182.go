/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 01:24:50
 * @link    github.com/taseikyo
 */

func dynamicPassword(password string, target int) string {
	if target > len(password) {
		target = target % len(password)
	}
	if target == 0 {
		return password
	}

	res := []byte(password)
	reverse(res, 0, len(password)-1)
	reverse(res, 0, len(password)-target-1)
	reverse(res, len(password)-target, len(password)-1)
	return string(res)
}

func reverse(s []byte, l, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
