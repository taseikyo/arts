/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-15 00:47:07
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    p, q := head, head
    newHead := &ListNode{
        Next: head,
    }
    tmp := newHead
    for p != nil && p.Next != nil {
        q = p.Next

        p.Next = q.Next
        q.Next = p
        tmp.Next = q

        tmp = p
        p = p.Next
    }

    return newHead.Next
}
