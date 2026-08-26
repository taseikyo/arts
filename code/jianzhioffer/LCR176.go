/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 16:18:21
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 1. 计算当前节点左右子树的高度
	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)

	// 2. 检查当前节点是否平衡，并递归检查其左右子树
	return abs(leftHeight-rightHeight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

// maxDepth 计算树的最大深度[reference:4]
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	// 如果高度返回 -1，则表示树不平衡
	return height(root) >= 0
}

// height 返回树的高度；若不平衡则返回 -1[reference:11]
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	// 剪枝：如果左子树已经不平衡，直接返回 -1
	if leftHeight == -1 {
		return -1
	}

	rightHeight := height(root.Right)
	// 剪枝：如果右子树已经不平衡，直接返回 -1
	if rightHeight == -1 {
		return -1
	}

	// 如果当前节点的左右子树高度差超过1，则当前树不平衡
	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	// 返回当前节点作为根节点的树的高度
	return max(leftHeight, rightHeight) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
