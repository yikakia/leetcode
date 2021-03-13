# 705.设计哈希集合
[https://leetcode-cn.com/problems/design-hashset](https://leetcode-cn.com/problems/design-hashset) 
## 原题
不使用任何内建的哈希表库设计一个哈希集合（HashSet）。

实现 `MyHashSet` 类：
-  `void add(key)` 向哈希集合中插入值 `key` 。
-  `bool contains(key)` 返回哈希集合中是否存在这个值 `key` 。
-  `void remove(key)` 将给定值 `key` 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。

 

 **示例：** 

```

输入：
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
输出：
[null, null, null, true, false, null, true, null, false]

解释：
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // 返回 True
myHashSet.contains(3); // 返回 False ，（未找到）
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // 返回 True
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // 返回 False ，（已移除）
```
 

 **提示：** 
-  `0 <= key <= 10^6` 
- 最多调用 `10^4` 次 `add` 、 `remove` 和 `contains` 。
 

 **进阶：** 你可以不使用内建的哈希集合库解决此问题吗？

 
**标签**
`设计` `哈希表` 


## 
```go

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
```
>执行用时: 84 ms
内存消耗: 8.4 MB

用的是链地址法。为了能够具有拓展性于是内置了哈希函数用于确定 key 映射的值。

其他的就是链地址法的实现了。因为用了标准库，所以还是比较简单的。
