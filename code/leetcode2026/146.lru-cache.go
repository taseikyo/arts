/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-16 01:24:45
 * @link    github.com/taseikyo
 */

// 定义双向链表的节点
type Node struct {
	key, val   int
	prev, next *Node
}

// 定义 LRU 缓存结构
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // 虚拟头节点
	tail     *Node // 虚拟尾节点
}

// 构造函数，初始化 LRU 缓存
func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// get 操作：从缓存中获取值
func (this *LRUCache) Get(key int) int {
	// 1. 从哈希表中查找节点
	if node, ok := this.cache[key]; ok {
		// 2. 如果存在，将该节点移动到链表头部 (表示最近使用)
		this.moveToHead(node)
		return node.val
	}
	return -1
}

// put 操作：向缓存中插入或更新值
func (this *LRUCache) Put(key int, value int) {
	// 1. 尝试从哈希表中查找节点
	if node, ok := this.cache[key]; ok {
		// 2. 如果 key 存在，更新值，并将节点移到头部
		node.val = value
		this.moveToHead(node)
		return
	}

	// 3. 如果 key 不存在，检查容量是否已满
	if len(this.cache) >= this.capacity {
		// 4. 如果满了，移除链表尾部的节点 (最久未使用)
		tailNode := this.removeTail()
		delete(this.cache, tailNode.key)
	}

	// 5. 创建新节点，加入哈希表，并添加到链表头部
	newNode := &Node{key: key, val: value}
	this.cache[key] = newNode
	this.addToHead(newNode)
}

// --- 辅助方法 ---

// moveToHead: 将某个已存在的节点移动到链表头部
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

// addToHead: 在虚拟头节点之后添加一个新节点
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// removeNode: 从链表中移除一个节点
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// removeTail: 移除链表尾部（虚拟尾节点之前）的真实节点，并返回它
func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
