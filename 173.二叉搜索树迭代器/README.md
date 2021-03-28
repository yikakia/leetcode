# 173.二叉搜索树迭代器
[https://leetcode-cn.com/problems/binary-search-tree-iterator](https://leetcode-cn.com/problems/binary-search-tree-iterator) 
## 原题
实现一个二叉搜索树迭代器类 `BSTIterator` ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
-  `BSTIterator(TreeNode root)` 初始化 `BSTIterator` 类的一个对象。BST 的根节点 `root` 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
-  `boolean hasNext()` 如果向指针右侧遍历存在数字，则返回 `true` ；否则返回 `false` 。
-  `int next()` 将指针向右移动，然后返回指针处的数字。
注意，指针初始化为一个不存在于 BST 中的数字，所以对 `next()` 的首次调用将返回 BST 中的最小元素。
你可以假设  `next()`  调用总是有效的，也就是说，当调用 `next()`  时，BST 的中序遍历中至少存在一个下一个数字。

 

 **示例：** 
<img alt="" src="https://assets.leetcode.com/uploads/2018/12/25/bst-tree.png" style="width: 189px; height: 178px;" />
```

输入
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
输出
[null, 3, 7, true, 9, true, 15, true, 20, false]

解释
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // 返回 3
bSTIterator.next();    // 返回 7
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 9
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 15
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 20
bSTIterator.hasNext(); // 返回 False

```
 

 **提示：** 
- 树中节点的数目在范围 `[1, 10^5]` 内
-  `0 <= Node.val <= 10^6` 
- 最多调用 `10^5` 次 `hasNext` 和 `next` 操作
 

 **进阶：** 
- 你可以设计一个满足下述条件的解决方案吗？ `next()` 和 `hasNext()` 操作均摊时间复杂度为 `O(1)` ，并使用 `O(h)` 内存。其中 `h` 是树的高度。
 
**标签**
`栈` `树` `设计` 


## 自己实现递归
```go
type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		cur: root,
	}
}

func (this *BSTIterator) Next() int {
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	this.cur, this.stack = this.stack[len(this.stack)-1], this.stack[:len(this.stack)-1]
	res := this.cur.Val
	this.cur = this.cur.Right
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) != 0
}
```
>执行用时: 24 ms
内存消耗: 9.5 MB

二叉树中序遍历最熟悉的就是用递归了，而递归的实现原理是使用了栈，在进行递归之前先把当前寄存器中的值压入栈中，然后再执行对应的函数。于是这里就用栈来实现递归。

从中序遍历的过程中我们可以看到，实际上是先找到最左的节点，然后再访问该节点（最左的），再访问它的父亲，再访问它的父亲的右儿子。

那么用栈来表示的时候就是要栈顶储存接下来可以访问的节点。这里很巧妙的一点就是用另外一个变量来储存当前要访问的节点，然后对这个节点不断地找左儿子，并且压栈。这个循环中这个节点如果是空的话就不用进行操作了。而为空可以是一开始的时候就为空，也可以是循环过程中为空，但是结果都是一样的，说明现在已经找到了最左的元素，这个时候就要把栈顶的元素弹出，然后找它的右儿子。

这里为什么是找右儿子呢？因为之前的循环已经找到了最左的节点了，即栈顶的元素，这个时候就是中序遍历中的 `左 中 右` 中的  `中` 的环节了。

于是这时候就是找右儿子了。这个时候还是按照中序遍历的思想，继续调用这个  `Next()` 函数进行查找。