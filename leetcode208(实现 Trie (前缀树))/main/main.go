package main

import "fmt"

//Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
//
//请你实现 Trie 类：
//
//Trie() 初始化前缀树对象。
//void insert(String word) 向前缀树中插入字符串 word 。
//boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
//boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false

//提示：
//
//1 <= word.length, prefix.length <= 2000
//word 和 prefix 仅由小写英文字母组成
//insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次

type Trie struct {
	children map[rune]*Trie // 表示该节点的所有子节点列表
	isEnd    bool           // 是否是叶子节点
}

func Constructor() Trie {
	return Trie{
		children: map[rune]*Trie{},
		isEnd:    false,
	}
}

func (t *Trie) Insert(word string) {
	// 如果找到了，就不插入，否则插入新的节点
	node := t
	for k, v := range word { // 遍历字符串字符
		if value, ok := node.children[v]; ok { // 已经有值了
			node = value // 当前节点为找到字符所在的节点
		} else { // 新增节点
			node.children[v] = &Trie{
				children: map[rune]*Trie{},
				isEnd:    false,
			}
			node = node.children[v] // 当前节点为新节点
		}
		if k == len(word)-1 {
			node.isEnd = true // 最后一个字符节点标记为true
		}
	}
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, v := range word { // 遍历字符串字符
		if value, ok := node.children[v]; ok { // 已经有值了
			node = value // 当前节点为找到字符所在的节点
		} else { // 没有找到
			return false
		}
	}
	if !node.isEnd {
		return false
	}
	return true
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, v := range prefix { // 遍历字符串字符
		if value, ok := node.children[v]; ok { // 已经有值了
			node = value // 当前节点为找到字符所在的节点
		} else { // 没有找到
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	obj := Constructor()
	obj.Insert("app")
	obj.Insert("apple")
	obj.Insert("beer")
	obj.Insert("add")
	obj.Insert("jam")
	obj.Insert("rental")
	fmt.Println(obj.Search("apps"))
	fmt.Println(obj.Search("app"))
}
