/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-16 01:37:48
 * @link    github.com/taseikyo
 */

package main

// 定义双向链表的节点
type Node struct {
    key, val   int
    freq       int
    prev, next *Node
}

// 定义双向链表
type DoubleList struct {
    head, tail *Node
}

// 创建一个新的双向链表，带有虚拟头尾节点
func NewDoubleList() *DoubleList {
    head, tail := &Node{}, &Node{}
    head.next, tail.prev = tail, head
    return &DoubleList{head: head, tail: tail}
}

// 在链表头部添加节点 (表示最近使用)
func (this *DoubleList) AddFirst(node *Node) {
    node.next = this.head.next
    node.prev = this.head
    this.head.next.prev = node
    this.head.next = node
}

// 从链表中移除一个节点
func (this *DoubleList) Remove(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
    // 清空被移除节点的指针，帮助GC
    node.prev, node.next = nil, nil
}

// 移除链表尾部节点 (表示最久未使用) 并返回它
func (this *DoubleList) RemoveLast() *Node {
    if this.IsEmpty() {
        return nil
    }
    last := this.tail.prev
    this.Remove(last)
    return last
}

// 检查链表是否为空
func (this *DoubleList) IsEmpty() bool {
    return this.head.next == this.tail
}

// 定义 LFU 缓存结构
type LFUCache struct {
    capacity int
    size     int
    minFreq  int
    // key -> Node 的映射
    keyNodeMap map[int]*Node
    // freq -> DoubleList 的映射
    freqListMap map[int]*DoubleList
}

// 构造函数，初始化 LFU 缓存
func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity:    capacity,
        keyNodeMap:  make(map[int]*Node),
        freqListMap: make(map[int]*DoubleList),
    }
}

// get 操作：从缓存中获取值
func (this *LFUCache) Get(key int) int {
    if this.capacity == 0 {
        return -1
    }
    // 1. 从 keyNodeMap 中查找节点
    if node, ok := this.keyNodeMap[key]; ok {
        // 2. 如果存在，增加其频次 (内部处理链表移动)
        this.increaseFreq(node)
        return node.val
    }
    return -1
}

// put 操作：向缓存中插入或更新值
func (this *LFUCache) Put(key int, value int) {
    if this.capacity == 0 {
        return
    }
    // 1. 尝试从 keyNodeMap 中查找节点
    if node, ok := this.keyNodeMap[key]; ok {
        // 2. 如果 key 存在，更新值，并增加其频次
        node.val = value
        this.increaseFreq(node)
        return
    }

    // 3. 如果 key 不存在，检查容量是否已满
    if this.size >= this.capacity {
        // 4. 如果满了，从 freqListMap[minFreq] 链表中移除尾部节点 (最久未使用)
        list := this.freqListMap[this.minFreq]
        deletedNode := list.RemoveLast()
        // 从 keyNodeMap 中删除该节点的 key
        delete(this.keyNodeMap, deletedNode.key)
        this.size--
    }

    // 5. 创建新节点，频次为 1
    newNode := &Node{key: key, val: value, freq: 1}
    this.keyNodeMap[key] = newNode

    // 6. 将新节点加入 freqListMap[1] 的链表头部
    if this.freqListMap[1] == nil {
        this.freqListMap[1] = NewDoubleList()
    }
    this.freqListMap[1].AddFirst(newNode)

    // 7. 更新最小频次为 1
    this.minFreq = 1
    this.size++
}

// --- 辅助方法：增加节点的频次 ---
func (this *LFUCache) increaseFreq(node *Node) {
    // 1. 从原频次对应的链表中移除该节点
    oldFreq := node.freq
    oldList := this.freqListMap[oldFreq]
    oldList.Remove(node)

    // 2. 如果原链表为空，则从 freqListMap 中删除它
    if oldList.IsEmpty() {
        delete(this.freqListMap, oldFreq)
        // 如果移除的正好是最小频次，则 minFreq 需要增加
        if this.minFreq == oldFreq {
            this.minFreq++
        }
    }

    // 3. 更新节点的频次
    node.freq++

    // 4. 将节点添加到新频次对应链表的头部
    newFreq := node.freq
    if this.freqListMap[newFreq] == nil {
        this.freqListMap[newFreq] = NewDoubleList()
    }
    this.freqListMap[newFreq].AddFirst(node)
}
