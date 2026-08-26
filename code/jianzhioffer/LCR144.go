/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-25 00:41:19
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
func flipTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	flipTree(root.Left)
	flipTree(root.Right)

	return root
}
