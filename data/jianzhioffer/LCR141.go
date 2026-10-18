/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 15:10:33
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainningPlan(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := trainningPlan(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
