/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-22 01:02:37
 * @link    github.com/taseikyo
 */

func trainingPlan(actions []int) []int {
	if len(actions) < 2 {
		return actions
	}

	l, r := 0, len(actions)-1
	for l < r {
		for r >= 0 && actions[r]%2 == 0 {
			r--
		}
		for l < len(actions) && actions[l]%2 != 0 {
			l++
		}
		if l > r {
			break
		}
		actions[l], actions[r] = actions[r], actions[l]
		l++
		r--
	}

	return actions
}
