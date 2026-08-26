/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 16:28:07
 * @link    github.com/taseikyo
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果 p 和 q 的值都小于当前节点，去左子树找
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	// 如果 p 和 q 的值都大于当前节点，去右子树找
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	// 否则，当前节点就是最近公共祖先
	return root
}
