/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 16:43:15
 * @link    github.com/taseikyo
 */

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1 / myPow(x, -n)
	} else if n == 1 {
        return x
    }

	res := myPow(x, n/2)
	if n%2 == 0 {
		return res * res
	}
	return res * res * x
}
