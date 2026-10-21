/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 16:33:28
 * @link    github.com/taseikyo
 */

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

// Constructor 初始化 Trie 对象
func Constructor() Trie {
	return Trie{}
}

// Insert 插入一个单词
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

// Search 查找一个完整的单词是否存在
func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd
}

// StartsWith 检查是否存在以给定前缀开头的单词
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}
