/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-14 23:57:11
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := make([]int, 0)
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			res = append(res, list1.Val)
			list1 = list1.Next
		} else {
			res = append(res, list2.Val)
			list2 = list2.Next
		}
	}
	for list1 != nil {
		res = append(res, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		res = append(res, list2.Val)
		list2 = list2.Next
	}

	var head, p *ListNode
	for _, v := range res {
		if head == nil {
			head = &ListNode{
				Val: v,
			}
			p = head
		} else {
			p.Next = &ListNode{
				Val: v,
			}
			p = p.Next
		}
	}

	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := &ListNode{}
    node := head
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            node.Next = list1
            list1 = list1.Next
        } else {
            node.Next= list2
            list2 = list2.Next
        }
        node = node.Next
    }
    if list1 != nil {
        node.Next = list1
    } else {
        node.Next = list2
    }
    return head.Next
}
