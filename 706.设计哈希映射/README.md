# 706.设计哈希映射
[https://leetcode-cn.com/problems/design-hashmap](https://leetcode-cn.com/problems/design-hashmap) 
## 原题
不使用任何内建的哈希表库设计一个哈希映射（HashMap）。

实现 `MyHashMap` 类：
-  `MyHashMap()` 用空映射初始化对象
-  `void put(int key, int value)` 向 HashMap 插入一个键值对 `(key, value)` 。如果 `key` 已经存在于映射中，则更新其对应的值 `value` 。
-  `int get(int key)` 返回特定的 `key` 所映射的 `value` ；如果映射中不包含 `key` 的映射，返回 `-1` 。
-  `void remove(key)` 如果映射中存在 `key` 的映射，则移除 `key` 和它所对应的 `value` 。
 

 **示例：** 

```

输入：
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
输出：
[null, null, null, 1, -1, null, 1, null, -1]

解释：
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // myHashMap 现在为 [[1,1]]
myHashMap.put(2, 2); // myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(1);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(3);    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
myHashMap.put(2, 1); // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
myHashMap.get(2);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
myHashMap.remove(2); // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
myHashMap.get(2);    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]

```
 

 **提示：** 
-  `0 <= key, value <= 10^6` 
- 最多调用 `10^4` 次 `put` 、 `get` 和 `remove` 方法
 

 **进阶：** 你能否不使用内置的 HashMap 库解决此问题？

 
**标签**
`设计` `哈希表` 


## 
```go
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

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	idx := this.hash(key)
	for e := this.buckets[idx].Front(); e != nil; e = e.Next() {
		if tnode := e.Value.(node); tnode.key == key {
			e.Value = node{key: key, value: value}
			return
		}
	}
	this.buckets[idx].PushBack(node{key: key, value: value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	idx := this.hash(key)
	for e := this.buckets[idx].Front(); e != nil; e = e.Next() {
		if et := e.Value.(node); et.key == key {
			return et.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	idx := this.hash(key)
	for e := this.buckets[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(node).key == key {
			this.buckets[idx].Remove(e)
		}
	}
}
```
>执行用时: 132 ms
内存消耗: 8.4 MB

和昨天的[705.设计哈希集合](../705.设计哈希集合/README.md)差不多，只不过储存的值变了一下而已。
