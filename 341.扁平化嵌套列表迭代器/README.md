# 341.扁平化嵌套列表迭代器
[https://leetcode-cn.com/problems/flatten-nested-list-iterator](https://leetcode-cn.com/problems/flatten-nested-list-iterator) 
## 原题
给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。

列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。

 

 **示例 1:** 

```
输入: [[1,1],2,[1,1]]
输出: [1,1,2,1,1]
解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
```
 **示例 2:** 

```
输入: [1,[4,[6]]]
输出: [1,4,6]
解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。

```
 
**标签**
`栈` `设计` 


## 
```go
type NestedIterator struct {
	stacks [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		stacks: [][]*NestedInteger{nestedList},
	}
}

func (this *NestedIterator) Next() int {
	last := this.stacks[len(this.stacks)-1]
	val := last[0].GetInteger()
	this.stacks[len(this.stacks)-1] = last[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stacks) > 0 {
		last := this.stacks[len(this.stacks)-1]
		if len(last) == 0 {
			this.stacks = this.stacks[:len(this.stacks)-1]
			continue
		}
		nest := last[0]
		if nest.IsInteger() {
			return true
		}
		this.stacks[len(this.stacks)-1] = last[1:]
		this.stacks = append(this.stacks, nest.GetList())
	}
	return false
}
```
>执行用时: 4 ms
内存消耗: 5.1 MB

原理就是在 HasNext() 的时候求出下一个。这样 Next 就能直接返回了。