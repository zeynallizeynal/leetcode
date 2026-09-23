package main

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	cur := t.root

	for _, r := range word {
		if cur.children[r] == nil {
			cur.children[r] = NewTrieNode()
		}
		cur = cur.children[r]
	}
	cur.isEnd = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, r := range word {
		if cur.children[r] == nil {
			return false
		}
		cur = cur.children[r]
	}
	return cur.isEnd
}

func (t *Trie) StartsWith(word string) bool {
	cur := t.root
	for _, r := range word {
		if cur.children[r] == nil {
			return false
		}
		cur = cur.children[r]
	}
	return true
}
