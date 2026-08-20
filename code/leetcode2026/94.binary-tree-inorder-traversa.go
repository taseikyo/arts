/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-16 23:28:19
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
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    }

    res := make([]int, 0)
    left := inorderTraversal(root.Left)
    right := inorderTraversal(root.Right)

    res = append(res, left...)
    res = append(res, root.Val)
    res = append(res, right...)

    return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}
