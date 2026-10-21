/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 15:20:09
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	root := &TreeNode{Val: preorder[0]}
	rootIdx := func() int {
		for idx, v := range inorder {
			if v == preorder[0] {
				return idx
			}
		}
		return 0
	}()
	root.Left = buildTree(preorder[1:rootIdx+1], inorder[0:rootIdx])
	root.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
	return root
}
