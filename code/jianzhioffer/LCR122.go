/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 21:20:44
 * @link    github.com/taseikyo
 */

func pathEncryption(path string) string {
	res := make([]byte, len(path))
	for i := range path {
		if path[i] == '.' {
			res[i] = ' '
		} else {
			res[i] = path[i]
		}
	}

	return string(res)
}
