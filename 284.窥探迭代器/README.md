# 284.窥探迭代器
[https://leetcode-cn.com/problems/peeking-iterator](https://leetcode-cn.com/problems/peeking-iterator) 
## 原题
请你设计一个迭代器，除了支持 `hasNext` 和 `next` 操作外，还支持 `peek` 操作。

实现 `PeekingIterator` 类：
-  `PeekingIterator(int[] nums)` 使用指定整数数组 `nums` 初始化迭代器。
-  `int next()` 返回数组中的下一个元素，并将指针移动到下个元素处。
-  `bool hasNext()` 如果数组中存在下一个元素，返回 `true` ；否则，返回 `false` 。
-  `int peek()` 返回数组中的下一个元素，但 **不** 移动指针。
 

 **示例：** 

```

输入：
["PeekingIterator", "next", "peek", "next", "next", "hasNext"]
[[[1, 2, 3]], [], [], [], [], []]
输出：
[null, 1, 2, 2, 3, false]

解释：
PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [1,2,3]
peekingIterator.next();    // 返回 1 ，指针移动到下一个元素 [1,2,3]
peekingIterator.peek();    // 返回 2 ，指针未发生移动 [1,2,3]
peekingIterator.next();    // 返回 2 ，指针移动到下一个元素 [1,2,3]
peekingIterator.next();    // 返回 3 ，指针移动到下一个元素 [1,2,3]
peekingIterator.hasNext(); // 返回 False

```
 

 **提示：** 
-  `1 <= nums.length <= 1000` 
-  `1 <= nums[i] <= 1000` 
- 对 `next` 和 `peek` 的调用均有效
-  `next` 、 `hasNext` 和 `peek` 最多调用 `1000` 次
 

 **进阶：** 你将如何拓展你的设计？使之变得通用化，从而适应所有的类型，而不只是整数型？

 
**标签**
`设计` `数组` `迭代器` 


## 预先取出
```go
type PeekingIterator struct {
	iter    *Iterator
	hasPeek bool
	peekNum int
}

func Constructor(iter *Iterator) *PeekingIterator {
	ret := &PeekingIterator{
		iter:    iter,
		hasPeek: iter.hasNext(),
		peekNum: iter.next(),
	}

	return ret
}

func (pt *PeekingIterator) hasNext() bool {
	return pt.hasPeek
}

func (pt *PeekingIterator) next() int {
	tmp := pt.peek()
	pt.hasPeek = pt.iter.hasNext()
	if pt.hasPeek {
		pt.peekNum = pt.iter.next()
	}
	return tmp
}

func (pt *PeekingIterator) peek() int {
	return pt.peekNum
}
```
>执行用时：0 ms
内存消耗：2.5 MB

就是模拟，在 next 的时候预先进行处理，获取下一个元素，用于 peek 调用。