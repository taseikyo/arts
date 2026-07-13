/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-12 16:33:52
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry int = 0
    res := l1
    tmp := l1
    for l1 != nil && l2 != nil {
        carry = carry + l1.Val + l2.Val
        l1.Val = carry % 10
        carry = carry / 10
        tmp = l1
        l1 = l1.Next
        l2 = l2.Next
    }
    for l2 != nil {
        tmp.Next = l2
        if carry == 0 {
            break
        }
        carry = carry + l2.Val
        l2.Val = carry % 10
        carry = carry / 10
        tmp = l2
        l2 = l2.Next
    }
    for l1 != nil {
        carry = carry + l1.Val
        l1.Val = carry % 10
        carry = carry / 10
        tmp = l1
        l1 = l1.Next
    }
    if carry > 0 {
        tmp.Next = new(ListNode)
        tmp.Next.Val = 1
    }

    return res
}