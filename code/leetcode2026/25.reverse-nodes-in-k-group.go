/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-16 00:26:22
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    prev := dummy // 上一组的尾巴，最开始指向虚拟头

    for {
        // 1. 【切】检查从 prev 开始，有没有 k 个节点？
        check := prev
        for i := 0; i < k; i++ {
            check = check.Next
            if check == nil { // 不够？不翻了，直接回家
                return dummy.Next
            }
        }

        // 2. 【翻】翻转这一组
        curr := prev.Next      // curr 是当前组第一个节点（例如例子的 1）
        next := curr.Next      // next 是当前组第二个节点（例如例子的 2）

        // 为什么循环 k-1 次？因为 2 个节点只需要交换 1 次指针
        for i := 0; i < k-1; i++ {
            temp := next.Next   // 先保存 2 后面的节点（防止丢了 3）
            next.Next = curr    // 关键！让 2 指向 1 (反转)
            curr = next         // 指针后移，准备翻下一个
            next = temp         // 指针后移
        }

        // 3. 【接】把翻好的组接回去
        // 注意：此时的 curr 已经是原来最后一个节点（例子中的 2）
        // 注意：此时的 next 已经是下一组的头（例子中的 3）
        groupHead := prev.Next   // 原来的组头（例子中的 1），现在它变成组尾了
        prev.Next = curr         // 上一组尾巴指向新组头（例如 dummy 指向 2）
        groupHead.Next = next    // 新组尾指向下一组头（例如 1 指向 3）
        prev = groupHead         // 移动 prev 到这一组的尾巴（1），准备下一轮
    }
}
