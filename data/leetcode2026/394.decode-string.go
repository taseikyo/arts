/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 00:06:52
 * @link    github.com/taseikyo
 */

func decodeString(s string) string {
    // 数字栈，存储重复次数
    numStack := []int{}
    // 字符串栈，存储括号外的字符串片段
    strStack := []string{}
    // 当前正在构建的字符串
    currStr := ""
    // 当前的重复次数
    num := 0

    for _, ch := range s {
        if ch >= '0' && ch <= '9' {
            // 1. 遇到数字：可能有多位，累积计算
            num = num*10 + int(ch-'0')
        } else if ch == '[' {
            // 2. 遇到 '['：将当前数字和字符串入栈，然后重置
            numStack = append(numStack, num)
            strStack = append(strStack, currStr)
            num = 0
            currStr = ""
        } else if ch == ']' {
            // 3. 遇到 ']'：弹出栈顶，构建重复字符串
            // 弹出重复次数
            repeatTimes := numStack[len(numStack)-1]
            numStack = numStack[:len(numStack)-1]
            // 弹出之前的字符串
            prevStr := strStack[len(strStack)-1]
            strStack = strStack[:len(strStack)-1]

            // 构建重复后的字符串
            repeated := ""
            for i := 0; i < repeatTimes; i++ {
                repeated += currStr
            }
            // 将重复后的字符串拼接到之前的字符串后面
            currStr = prevStr + repeated
        } else {
            // 4. 遇到普通字母：直接拼接到当前字符串
            currStr += string(ch)
        }
    }
    return currStr
}
