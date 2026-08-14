/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-14 17:20:10
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
        if fast == slow {
			return true
		}
	}

	return false
}
