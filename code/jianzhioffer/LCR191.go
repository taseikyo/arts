/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 20:59:56
 * @link    github.com/taseikyo
 */

func statisticalResult(arrayA []int) []int {
	n := len(arrayA)
	if n < 1 {
		return nil
	}
	p, r := make([]int, n), make([]int, n)

	p[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		p[i] = p[i+1] * arrayA[i+1]
	}

	r[0] = 1
	for i := 1; i < n; i++ {
		r[i] = r[i-1] * arrayA[i-1]
	}

	for i := range p {
		p[i] *= r[i]
	}

	return p
}
