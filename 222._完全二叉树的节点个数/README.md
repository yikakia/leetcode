# 222. 完全二叉树的节点个数
[https://leetcode-cn.com/problems/count-complete-tree-nodes/](https://leetcode-cn.com/problems/count-complete-tree-nodes/) 
## 深搜
```go
func countNodes(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	res += countNodes(root.Left) + countNodes(root.Right) + 1
	return res
}
```
>执行用时: 20 ms
内存消耗: 7.1 MB

只是一般的解法，没有用到完全二叉树的性质。


## 利用二叉树的性质
### 利用满二叉树判断
```go
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	levelLeft := 1
	for node := root; node.Left != nil; node = node.Left {
		levelLeft++
	}
	levelRight := 1
	for node := root; node.Right != nil; node = node.Right {
		levelRight++
	}
	if levelLeft == levelRight {
		return (1 << levelLeft) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
```
>执行用时: 16 ms
内存消耗: 7.1 MB

这里是用了满二叉树的性质，即层数`h`与整个二叉树的节点之和`SumNodes`满足$SumNodes=2^h-1$的关系（这里设根节点所在的层为 **0** 层）。如果是满二叉树，那就不用继续找了，直接算出来返回就好。相当于是在基础的**bfs**上面做了剪枝的操作。

### 官方题解

相当于二分，范围是 $1$~$2^{h+1}-1$二分的时候是判断现在找的这个点存不存在，而如果我们把根节点记为 1，然后按照每层从左至右递增的顺序累加的话，我们会发现把要找的第 n 个节点转成二进制的话，它的每一位上的值似乎能够代表应该怎么找到它。比如`6`转换成二进制的话是 $110_{b}$ ，去除首位的`1`不看，从第二位开始

- 当该位为 `1` 的时候，去右子树 （即 node = node.Left）
- 当该位为 `0` 的时候，去左子树 （即 node = node.Right）

按照这个规则，我们可以很快地遍历到这个节点。那么要进行二分的话，就是判断这个节点是否存在，即节点$2^h$到节点$2^{h+1}-1$之间是否能够根据这个规则找到这个节点。

初始的条件可以是 $l=1$ 而 $r=2^h$每次都找$mid=\frac{r-l}{2} + l$看这个节点是否存在，

- 如果存在就`l=mid+1`，因为这个时候**l**不一定就是我们要找的那个节点。

- 如果不存在就`r=mid`，因为要找的那个节点肯定在它的左边，右边的都可以去掉。

```go
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}
```