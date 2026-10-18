/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 15:28:36
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
func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }
    // 包含当前节点的路径数 + 不包含当前节点的路径数（在左、右子树中）
    res := dfs(root, targetSum)
    res += pathSum(root.Left, targetSum)
    res += pathSum(root.Right, targetSum)
    return res
}

// dfs 计算从 node 出发，向下寻找有多少条路径和等于 targetSum
func dfs(node *TreeNode, targetSum int) int {
    if node == nil {
        return 0
    }
    res := 0
    if node.Val == targetSum {
        res++
    }
    // 继续向下寻找，目标值减去当前节点的值
    res += dfs(node.Left, targetSum - node.Val)
    res += dfs(node.Right, targetSum - node.Val)
    return res
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    // mp 存储前缀和及其出现次数
    // 初始化 mp[0] = 1 至关重要，它代表“空路径”，用于处理从根节点开始的路径[reference:8]
    mp := map[int]int{0: 1}
    return preorder(root, 0, targetSum, mp)
}

func preorder(node *TreeNode, currSum int, target int, mp map[int]int) int {
    if node == nil {
        return 0
    }

    // 1. 更新当前路径和
    currSum += node.Val

    // 2. 查找是否存在满足条件的前缀和
    //    核心公式：当前前缀和 - target = 之前某个前缀和
    res := mp[currSum - target]

    // 3. 将当前前缀和加入哈希表，供子节点使用
    mp[currSum]++

    // 4. 递归处理左右子树
    res += preorder(node.Left, currSum, target, mp)
    res += preorder(node.Right, currSum, target, mp)

    // 5. 回溯：从哈希表中移除当前前缀和，避免影响其他路径[reference:9]
    mp[currSum]--

    return res
}
