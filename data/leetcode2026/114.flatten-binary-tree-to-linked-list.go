/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 15:05:53
 * @link    github.com/taseikyo
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	stack := []*TreeNode{}
	pre := root
	for root != nil || len(stack) > 0 {
		if root != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root.Right = root.Left
			root.Left = nil
			pre = root
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pre.Right = root
		}
	}
}

func flatten(root *TreeNode) {
    cur := root
    for cur != nil {
        // 1. 如果当前节点有左子树，才需要处理
        if cur.Left != nil {
            // 2. 寻找左子树中的最右节点（前驱节点）
            pre := cur.Left
            for pre.Right != nil {
                pre = pre.Right
            }

            // 3. 将当前节点的右子树接到前驱节点的右边
            pre.Right = cur.Right

            // 4. 将左子树移到右边，并将左指针置空
            cur.Right = cur.Left
            cur.Left = nil
        }
        // 5. 移动到下一个节点（即当前节点的右子节点）
        cur = cur.Right
    }
}
