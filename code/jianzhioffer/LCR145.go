/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-25 00:47:36
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
func checkSymmetricTree(root *TreeNode) bool {
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
