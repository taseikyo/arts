/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 14:57:49
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBookList(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}

	return res
}

func reverseBookList(head *ListNode) []int {
	newHead := reverse(head)
	res := []int{}
	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}

    return res
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
