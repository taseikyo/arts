/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-22 00:32:10
 * @link    github.com/taseikyo
 */

func inventoryManagement(stock []int) int {
	res, vote := 0, 0
	for _, v := range stock {
		if vote == 0 {
			res = v
		}
		if v == res {
			vote++
		} else {
			vote--
		}
	}

	return res
}
