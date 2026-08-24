/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 15:08:53
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainingPlan(head *ListNode, cnt int) *ListNode {
	fast, slow := head, head
	for cnt > 0 {
		fast = fast.Next
		cnt--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
