/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 14:07:33
 * @link    github.com/taseikyo
 */

func partition(s string) [][]string {
    res := [][]string{}
    path := []string{}

    // 判断子串 s[l:r+1] 是否为回文串
    isPalindrome := func(l, r int) bool {
        for l < r {
            if s[l] != s[r] {
                return false
            }
            l++
            r--
        }
        return true
    }

    var dfs func(int)
    dfs = func(start int) {
        // 到达字符串末尾，找到一种分割方案
        if start == len(s) {
            // 关键：拷贝一份 path，因为后续回溯会修改它[reference:11]
            tmp := make([]string, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }

        for i := start; i < len(s); i++ {
            // 剪枝：只有当前子串是回文串，才继续递归[reference:12]
            if isPalindrome(start, i) {
                path = append(path, s[start:i+1]) // 做选择
                dfs(i + 1)                        // 递归
                path = path[:len(path)-1]         // 撤销选择（回溯）
            }
        }
    }

    dfs(0)
    return res
}
