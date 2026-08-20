/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 00:17:35
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric2(root.Left, root.Right)
}

func isSymmetric2(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right // 只有当 left 和 right 同时为 nil 时才 true
	}
	return left.Val == right.Val &&
		isSymmetric2(left.Left, right.Right) &&
		isSymmetric2(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    queue := []*TreeNode{root.Left, root.Right}
    for len(queue) > 0 {
        l, r := queue[0], queue[1]
        queue = queue[2:]
        if l == nil && r == nil {
            continue
        }
        if l == nil || r == nil || l.Val != r.Val {
            return false
        }
        queue = append(queue, l.Left, r.Right, l.Right, r.Left)
    }
    return true
}
