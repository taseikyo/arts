/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-16 00:47:15
 * @link    github.com/taseikyo
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 1. 建立原节点到新节点的映射
	cache := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		cache[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	// 2. 连接新节点的 Next 和 Random 指针
	cur = head
	for cur != nil {
		// 从 map 中取出当前节点对应的新节点
		newNode := cache[cur]
		// 设置新节点的 Next 指针
		newNode.Next = cache[cur.Next]
		// 设置新节点的 Random 指针
		newNode.Random = cache[cur.Random]
		cur = cur.Next
	}

	// 返回原头节点对应的新节点
	return cache[head]
}
