/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 00:58:35
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
func calculateDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		res++
		size := len(queue)
		for _, node := range queue[:size] {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
	}
	return res

}
