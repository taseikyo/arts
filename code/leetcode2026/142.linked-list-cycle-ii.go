/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-14 17:36:01
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    // 1. 第一阶段：判断是否有环
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            // 有环，进入第二阶段
            // 2. 第二阶段：找到环的入口
            ptr := head
            for ptr != slow {
                ptr = ptr.Next
                slow = slow.Next
            }
            return ptr // 环的入口节点
        }
    }

    return nil
}
