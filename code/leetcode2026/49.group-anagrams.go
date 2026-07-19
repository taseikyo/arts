/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-19 23:23:44
 * @link    github.com/taseikyo
 */

func groupAnagrams(strs []string) [][]string {
    cache := make(map[string][]string)
    for _, str := range strs {
        bstr := []byte(str)
        sort.Slice(bstr, func(i, j int) bool {
            return bstr[i] < bstr[j]
        })
        sstr := string(bstr)
        if _, ok := cache[sstr]; ok {
            cache[sstr] = append(cache[sstr], str)
        } else {
            cache[sstr] = []string{str}
        }
    }
    res := make([][]string, 0)
    for _, v := range cache {
        res = append(res, v)
    }

    return res
}
