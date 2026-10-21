/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 16:30:54
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 1. 终止条件：遇到空节点，或找到 p 或 q 就返回
	if root == nil || root == p || root == q {
		return root
	}

	// 2. 在左、右子树中递归查找
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 3. 根据返回值判断
	if left != nil && right != nil {
		// 情况1: p 和 q 分别在左右子树，当前节点即为 LCA
		return root
	}
	if left != nil {
		// 情况2: 都在左子树
		return left
	}
	// 情况3: 都在右子树，或者都没找到
	return right
}
