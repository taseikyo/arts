/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-15 00:33:27
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	pre := &ListNode{
		Next: head,
	}
	res := pre
	for n > 0 {
		q = q.Next
		n--
	}
	for q != nil {
		pre = p
		p = p.Next
		q = q.Next
	}
	pre.Next = pre.Next.Next

	return res.Next
}
