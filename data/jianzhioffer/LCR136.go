/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 15:05:44
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	if head.Val == val {
		return head.Next
	}

	fast, slow := head.Next, head
	for fast != nil {
		if fast.Val == val {
			slow.Next = fast.Next
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	return head
}
