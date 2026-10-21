/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 01:04:48
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
func sortedArrayToBST(nums []int) *TreeNode {
    size := len(nums)
    if size == 0 {
        return nil
    }
    if size == 1 {
        return &TreeNode{
            Val: nums[0],
        }
    }

    rootIdx := size / 2
    root := &TreeNode{Val: nums[rootIdx]}
    root.Left = sortedArrayToBST(nums[:rootIdx])
    if rootIdx < size-1 {
        root.Right = sortedArrayToBST(nums[rootIdx+1:])
    }

    return root
}
