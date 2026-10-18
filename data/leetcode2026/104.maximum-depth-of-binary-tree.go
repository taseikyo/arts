/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-16 23:42:01
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
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }

    left := maxDepth(root.Left) + 1
    right := maxDepth(root.Right) + 1

    return max(left, right)
}

// BFS 采用层序遍历的方式，像剥洋葱一样，一层一层地遍历树的所有节点。
// 每遍历完一层，深度计数器就加 1。当队列为空时，计数器的值就是树的最大深度。

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    depth := 0
    queue := []*TreeNode{root} // 初始化队列，将根节点放入

    for len(queue) > 0 {
        // 记录当前层的节点数
        levelSize := len(queue)

        // 将当前层的所有节点出队，并将它们的子节点入队
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:] // 出队

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        // 当前层遍历完毕，深度+1
        depth++
    }
    return depth
}



// DFS 采用前序遍历的方式，沿着一条路径走到黑，再回溯。
// 我们可以用栈来模拟这个过程，栈中每个元素除了节点本身，还需记录该节点当前的深度。
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    maxDepth := 0
    // 栈中存储的是 [节点, 当前深度]
    stack := [][2]interface{}{{root, 1}}

    for len(stack) > 0 {
        // 弹出栈顶元素
        item := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        node := item[0].(*TreeNode)
        currentDepth := item[1].(int)

        // 更新最大深度
        if currentDepth > maxDepth {
            maxDepth = currentDepth
        }

        // 将子节点入栈，深度+1
        if node.Left != nil {
            stack = append(stack, [2]interface{}{node.Left, currentDepth + 1})
        }
        if node.Right != nil {
            stack = append(stack, [2]interface{}{node.Right, currentDepth + 1})
        }
    }
    return maxDepth
}
