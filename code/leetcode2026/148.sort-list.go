/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-16 01:06:51
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1. 找到中点，分割链表
	mid := findMid(head)
	rightHead := mid.Next
	mid.Next = nil

	// 2. 递归排序
	left := sortList(head)
	right := sortList(rightHead)

	// 3. 合并
	return merge(left, right)
}

// 快慢指针找中点
func findMid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 合并两个有序链表
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

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    // 1. 计算链表长度
    length := 0
    for cur := head; cur != nil; cur = cur.Next {
        length++
    }

    dummy := &ListNode{Next: head}

    // 2. 逐步增加合并的步长：1, 2, 4, 8, ...
    for step := 1; step < length; step <<= 1 {
        prev, cur := dummy, dummy.Next

        for cur != nil {
            // 2.1 找到第一段的头 head1，并截取 step 个节点
            head1 := cur
            for i := 1; i < step && cur.Next != nil; i++ {
                cur = cur.Next
            }

            // 2.2 找到第二段的头 head2
            head2 := cur.Next
            cur.Next = nil // 切断第一段
            cur = head2

            // 2.3 截取第二段的 step 个节点
            for i := 1; i < step && cur != nil && cur.Next != nil; i++ {
                cur = cur.Next
            }

            // 2.4 保存下一段的起点
            next := (*ListNode)(nil)
            if cur != nil {
                next = cur.Next
                cur.Next = nil // 切断第二段
            }

            // 2.5 合并 head1 和 head2，接到 prev 后面
            prev.Next = merge(head1, head2)

            // 2.6 移动 prev 到合并后链表的尾部
            for prev.Next != nil {
                prev = prev.Next
            }

            cur = next // 处理下一组
        }
    }

    return dummy.Next
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
