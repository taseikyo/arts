/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 01:16:53
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
func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)

	res := 0
	for root != nil || len(stack) > 0 {
		if root == nil {
			root = stack[len(stack)-1]
			k--
			if k == 0 {
				res = root.Val
				break
			}
			stack = stack[:len(stack)-1]
			root = root.Right
		} else {
			stack = append(stack, root)
			root = root.Left
		}
	}

	return res
}
