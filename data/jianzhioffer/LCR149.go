/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-25 00:52:52
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
func decorateRecord(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) > 0 {
		size := len(queue)
		for _, node := range queue[:size] {
			res = append(res, node.Val)
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
