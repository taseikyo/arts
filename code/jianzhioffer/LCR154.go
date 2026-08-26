/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 15:27:56
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

// A -> A' -> B -> B'
func copyRandomList(head *Node) *Node {
	cur := head
	// 生成复制的节点
	for cur != nil {
		tmp := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next = tmp
		cur = tmp.Next
	}
	// 复制 random 指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 生成新链表，还原老链表
	dummy := &Node{}
	p := dummy
	cur = head
	for cur != nil {
		p.Next = cur.Next
		cur.Next = cur.Next.Next
		p = p.Next
		cur = cur.Next
	}

	return dummy.Next
}
