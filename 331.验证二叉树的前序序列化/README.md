# 331.验证二叉树的前序序列化
[https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree](https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree) 
## 原题
序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 `#` 。

```
     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #

```
例如，上面的二叉树可以被序列化为字符串 `"9,3,4,#,#,1,#,#,2,#,6,#,#"` ，其中 `#` 代表一个空节点。

给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。

每个以逗号分隔的字符或为一个整数或为一个表示 `null` 指针的 `';#';` 。

你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 `"1,,3"` 。

 **示例 1:** 

```
输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
输出: true
```
 **示例 2:** 

```
输入: "1,#"
输出: false

```
 **示例 3:** 

```
输入: "9,#,#,1"
输出: false
```
 
**标签**
`栈` 


## 标记还可以存在多少节点
```go
func isValidSerialization(preorder string) bool {
	split := strings.Split(preorder, ",")
	freeNodes := 1
	for _, s := range split[:] {
		if freeNodes == 0 {
			return false
		}
		if s != "#" {
			freeNodes++
		} else {
			freeNodes--
		}
	}
	return freeNodes == 0
}
```
>执行用时: 0 ms
内存消耗: 2.7 MB

定义一个值，用来表示这个二叉树里面还可以容纳多少节点。现在把空节点也看做是一个节点，那么可以知道，空节点后面不能再有节点，因此遇到一个空节点的时候就要把可以容纳的值减一。而遇到非空的节点的时候，就要把节点的值加二，因为自己还要占一个位置，所以还要把可以容纳的值减一，综合起来就是把节点的值加一。

因此我们在循环过程中如果遇到一个值在插入之前可以容纳的值就为 0 的时候，那么就说明这个序列化是非法的。因为整个序列是这个时候二叉树已经满了。

特别地，我们要让这个可以容纳的值的初始值为 1 。可以把它理解为这个需要求的序列是二叉树

```
   _1_
  /   \
  #   waitToCheck
```

的右子树的前序遍历。