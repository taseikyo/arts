/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-16 01:14:35
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }
    if len(lists) == 2 {
        return merge(lists[0], lists[1])
    }

    return merge(mergeKLists(lists[0: len(lists)/2]), mergeKLists(lists[len(lists)/2: len(lists)]))
}

func merge(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }
        tail = tail.Next
    }

    if l1 != nil {
        tail.Next = l1
    } else {
        tail.Next = l2
    }

    return dummy.Next
}
