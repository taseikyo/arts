/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-12 16:55:55
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
    
    tempHead := &ListNode{}
    p := tempHead
    pre := 0

    for l1 != nil || l2 != nil || pre != 0{

        val1, val2 := 0, 0
        if l1 != nil {
            val1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            val2 = l2.Val
            l2 = l2.Next
        }
        sum := val1 + val2 + pre
        pre = sum / 10
        sum = sum % 10
        p.Next = &ListNode{
            Val: sum,
        }
        p = p.Next
    }

    return tempHead.Next
}
