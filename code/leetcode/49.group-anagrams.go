/*
* @Date:   2025-04-10 22:45:53
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
func groupAnagrams(strs []string) [][]string {
    if len(strs) == 1 {
        return [][]string{strs}
    }

    maps := make(map[string][]string, 0)
    for _, str := range strs {
        bytes := []byte(str)
        sort.Slice(bytes, func(i, j int) bool {
            return bytes[i] < bytes[j]
        })

        newStr := string(bytes)
        maps[newStr] = append(maps[newStr], str)
    }

    res := make([][]string, 0)
    for _, v := range maps {
        res = append(res, v)
    }

    return res
}
