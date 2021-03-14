package leetcode

import (
	"container/list"
	"testing"
)

// 706.设计哈希映射
// https://leetcode-cn.com/problems/design-hashmap

const base = 769

type node struct {
	key   int
	value int
}

type MyHashMap struct {
	buckets []list.List
	hash    func(int) int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		buckets: make([]list.List, base),
		hash: func(key int) int {
			return key % base
		},
	}
}

func (this *MyHashMap) contain(key int) *list.Element {
	idx := this.hash(key)
	for e := this.buckets[idx].Back(); e != nil; e = e.Next() {
		if e.Value.(node).key == key {
			return e
		}
	}
	return nil
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	idx := this.hash(key)
	if e := this.contain(key); e != nil {
		e.Value = node{key: key, value: value}
		return
	}
	this.buckets[idx].PushBack(node{key: key, value: value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if e := this.contain(key); e != nil {
		return e.Value.(node).value
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	e := this.contain(key)
	if e != nil {
		idx := this.hash(key)
		this.buckets[idx].Remove(e)
	}
}

func TestSolution(t *testing.T) {
	type op struct {
		name string
		k, v int
	}
	testCases := []struct {
		desc string
		ops  []op
	}{
		{
			ops: []op{
				{name: "get", k: 15, v: -1},
				{name: "put", k: 1, v: 1},
				{name: "put", k: 10, v: 1},
				{name: "put", k: 2, v: 2},
				{name: "get", k: 10, v: 1},
				{name: "get", k: 1, v: 1},
				{name: "get", k: 2, v: 2},
				{name: "put", k: 2, v: 1},
				{name: "get", k: 2, v: 1},
				{name: "remove", k: 2, v: -1},
				{name: "get", k: 2, v: -1},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m := Constructor()
			for _, o := range tC.ops {
				switch o.name {
				case "put":
					m.Put(o.k, o.v)
				case "get":
					if get := m.Get(o.k); get != o.v {
						t.Error("get Wrong\n", o, get)
					}
				case "remove":
					m.Remove(o.k)
				default:
					panic("wrong op name ")
				}
			}
		})
	}
}
