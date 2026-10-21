/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-14 16:55:53
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    p := reverseList(head.Next)
    // 将当前 head 节点，添加到已反转链表的末尾
    // 此时 head.Next 是已反转链表的尾节点（即原链表的倒数第二个节点）
    head.Next.Next = head
    head.Next = nil
    return p
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
