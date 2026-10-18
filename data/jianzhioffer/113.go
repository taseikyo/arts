/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 00:39:45
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
func pathSum(root *TreeNode, targetSum int) [][]int {
    var result [][]int
    var path []int

    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, remaining int) {
        // 1. 如果节点为空，直接返回
        if node == nil {
            return
        }

        // 2. 将当前节点加入路径
        path = append(path, node.Val)
        // 更新剩余目标值
        remaining -= node.Val

        // 3. 检查是否到达叶子节点且路径和为目标值
        if node.Left == nil && node.Right == nil && remaining == 0 {
            // 关键：拷贝一份 path 的副本，因为 path 后续会被修改
            temp := make([]int, len(path))
            copy(temp, path)
            result = append(result, temp)
        }

        // 4. 递归遍历左右子树
        dfs(node.Left, remaining)
        dfs(node.Right, remaining)

        // 5. 回溯：将当前节点从路径中移除
        path = path[:len(path)-1]
    }

    dfs(root, targetSum)
    return result
}
