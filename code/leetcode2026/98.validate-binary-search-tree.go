/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 01:05:26
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
func isValidBST(root *TreeNode) bool {
	// 初始范围设置为最大和最小整数
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	// 当前节点值必须在 (min, max) 开区间内
	if node.Val <= min || node.Val >= max {
		return false
	}
	// 递归检查左子树（更新上界）和右子树（更新下界）
	return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}
