/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 00:51:37
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
func findTargetNode(root *TreeNode, cnt int) int {
	stack := []*TreeNode{}
	res := []int{}
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			res = append(res, root.Val)
			if len(res) == cnt {
				break
			}
			stack = stack[:len(stack)-1]
			root = root.Left
		}
	}
	return res[len(res)-1]
}
