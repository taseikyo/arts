/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 16:04:16
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
func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32

    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        // 1. 后序遍历：先计算左右子树的贡献值
        //    如果子树贡献为负，则舍弃，贡献值记为 0
        leftGain := max(dfs(node.Left), 0)
        rightGain := max(dfs(node.Right), 0)

        // 2. 计算“经过当前节点”的最大路径和，并更新全局最大值
        //    这个路径由 node.Val + leftGain + rightGain 组成
        currentPathSum := node.Val + leftGain + rightGain
        if currentPathSum > maxSum {
            maxSum = currentPathSum
        }

        // 3. 返回当前节点能向父节点提供的最大“贡献值”
        //    父节点只能选择当前节点的一条路径，所以是 node.Val + max(leftGain, rightGain)
        return node.Val + max(leftGain, rightGain)
    }

    dfs(root)
    return maxSum
}
