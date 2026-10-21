/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-14 17:13:42
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	count := 0
	p := head
	for p != nil {
		p = p.Next
		count++
	}
	count = (count + 1) / 2
	p, q := head, head
	for count > 0 {
		q = q.Next
		count--
	}
	q = reverseList(q)
	for q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}
	return newHead
}
