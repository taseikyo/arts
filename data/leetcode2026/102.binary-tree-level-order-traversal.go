/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 00:39:00
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
		queue = queue[size:]
	}

	return res
}
