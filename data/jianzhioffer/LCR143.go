/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-25 00:54:04
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

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 约定空树不是任何树的子结构[reference:15]
	if A == nil || B == nil {
		return false
	}

	// 辅助函数：判断以 A 为根的树是否包含 B
	var recur func(a, b *TreeNode) bool
	recur = func(a, b *TreeNode) bool {
		// 1. 如果 B 已经遍历完，说明匹配成功[reference:16][reference:17]
		if b == nil {
			return true
		}
		// 2. 如果 A 已经遍历完，但 B 还没完，匹配失败[reference:18]
		if a == nil {
			return false
		}
		// 3. 如果当前节点值不同，匹配失败[reference:19]
		if a.Val != b.Val {
			return false
		}
		// 4. 当前节点相同，继续递归比较左右子树[reference:20]
		return recur(a.Left, b.Left) && recur(a.Right, b.Right)
	}

	// 主逻辑：[reference:21][reference:22]
	// 1. 检查当前 A 节点是否包含 B
	// 2. 递归检查 A 的左子树
	// 3. 递归检查 A 的右子树
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
