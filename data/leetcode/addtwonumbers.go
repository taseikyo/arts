/*
* @Date:   2025-03-02 18:51:26
* @Author: Lewis Tian (taseikyo@gmail.com)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	overflow := (l1.Val + l2.Val) / 10
	res := &ListNode{}
    res.Val = (l1.Val + l2.Val) % 10
	head := res
    l1 = l1.Next
    l2 = l2.Next
	for l1 != nil || l2 != nil || overflow > 0 {
        res.Next = &ListNode{}
        res = res.Next
        res.Val += overflow
        if l1 != nil {
            res.Val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            res.Val += l2.Val
            l2 = l2.Next
        }
		overflow = res.Val / 10
		res.Val = res.Val % 10
	}

	return head
}
