package main

import (
	"fmt"
	"hash/fnv"
)

type entry[K comparable, V interface{}] struct {
	key   K
	value V
	next  *entry[K, V]
}

type HashTable[K comparable, V interface{}] struct {
	buckets []*entry[K, V]
	size    int
}

func NewHashTable[K comparable, V interface{}](size int) *HashTable[K, V] {
	return &HashTable[K, V]{
		buckets: make([]*entry[K, V], size),
		size:    size,
	}
}

func (ht *HashTable[K, V]) hash(key K) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(fmt.Sprintf("%v", key)))
	if err != nil {
		panic("could not convert to hash")
	}
	return int(h.Sum32()) % ht.size
}

func (ht *HashTable[K, V]) Put(key K, value V) {
	index := ht.hash(key)
	head := ht.buckets[index]
	current := head
	for current != nil {
		if current.key == key {
			current.value = value
			return
		}
		current = current.next
	}

	newEntry := &entry[K, V]{
		key:   key,
		value: value,
		next:  head,
	}
	ht.buckets[index] = newEntry
}

func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	index := ht.hash(key)
	head := ht.buckets[index]
	current := head
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}

	var zeroVal V
	return zeroVal, false
}

func (ht *HashTable[K, V]) Delete(key K) bool {
	index := ht.hash(key)
	current := ht.buckets[index]
	var prev *entry[K, V]

	for current != nil {
		if current.key == key {
			if prev == nil {
				ht.buckets[index] = current.next
			} else {
				prev.next = current.next
			}
			return true
		}
		prev = current
		current = current.next
	}
	return false
}
