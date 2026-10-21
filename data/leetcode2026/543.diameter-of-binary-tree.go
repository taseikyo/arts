/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-17 00:29:59
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
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        leftDepth := dfs(node.Left)
        rightDepth := dfs(node.Right)

        // 当前节点的直径 = 左子树深度 + 右子树深度，更新全局最大值
        if leftDepth+rightDepth > diameter {
            diameter = leftDepth + rightDepth
        }

        // 返回当前节点的最大深度
        if leftDepth > rightDepth {
            return leftDepth + 1
        }
        return rightDepth + 1
    }
    dfs(root)
    return diameter
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    diameter := 0
    // 栈用于存储节点，map用于存储每个节点对应的最大深度[reference:6]
    stack := []*TreeNode{root}
    depthMap := make(map[*TreeNode]int)

    // 后序遍历
    for len(stack) > 0 {
        node := stack[len(stack)-1]

        // 如果当前节点的左右子树还未处理，则将其入栈
        if node.Left != nil && depthMap[node.Left] == 0 {
            stack = append(stack, node.Left)
        } else if node.Right != nil && depthMap[node.Right] == 0 {
            stack = append(stack, node.Right)
        } else {
            // 左右子树都已处理完毕（或为空），可以计算当前节点
            stack = stack[:len(stack)-1] // 出栈

            leftDepth := depthMap[node.Left]
            rightDepth := depthMap[node.Right]

            // 更新直径
            if leftDepth+rightDepth > diameter {
                diameter = leftDepth + rightDepth
            }

            // 记录当前节点的最大深度
            depthMap[node] = max(leftDepth, rightDepth) + 1
        }
    }
    return diameter
}
