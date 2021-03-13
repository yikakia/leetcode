package leetcode

import (
	"container/list"
)

// 705.设计哈希集合
// https://leetcode-cn.com/problems/design-hashset

const base = 769

type MyHashSet struct {
	buckets []list.List
	hash    func(int) int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{buckets: make([]list.List, base),
		hash: func(key int) int {
			return key % base
		}}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		hash := this.hash(key)
		this.buckets[hash].PushBack(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	hash := this.hash(key)
	for e := this.buckets[hash].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.buckets[hash].Remove(e)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	hash := this.hash(key)
	for e := this.buckets[hash].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// func TestSolution(t *testing.T) {
// 	testCases := []struct {
// 		desc string
// 		want
// 	}{
// 		{
//             want: ,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			get :=
// 			if !reflect.DeepEqual(tC.want,get){
// 				t.Errorf("input: %+v get: %v\n",tC,get)
// 			}
// 		})
// 	}
// }
