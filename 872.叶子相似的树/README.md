# 872.叶子相似的树
[https://leetcode-cn.com/problems/leaf-similar-trees](https://leetcode-cn.com/problems/leaf-similar-trees) 
## 原题
请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个  *叶值序列* 。

<img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/16/tree.png" style="height: 240px; width: 300px;" />

举个例子，如上图所示，给定一棵叶值序列为  `(6, 7, 4, 9, 8)`  的树。

如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是  *叶相似* 的。

如果给定的两个根结点分别为  `root1` 和  `root2`  的树是叶相似的，则返回  `true` ；否则返回 `false` 。

 

 **示例 1：** 

<img alt="" src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-1.jpg" style="height: 297px; width: 750px;" />

```

输入：root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
输出：true

```
 **示例 2：** 

```

输入：root1 = [1], root2 = [1]
输出：true

```
 **示例 3：** 

```

输入：root1 = [1], root2 = [2]
输出：false

```
 **示例 4：** 

```

输入：root1 = [1,2], root2 = [2,2]
输出：true

```
 **示例 5：** 

<img alt="" src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-2.jpg" style="height: 165px; width: 450px;" />

```

输入：root1 = [1,2,3], root2 = [1,3,2]
输出：false

```
 

 **提示：** 
- 给定的两棵树可能会有  `1`  到 `200`  个结点。
- 给定的两棵树上的值介于 `0` 到 `200` 之间。
 
**标签**
`树` `深度优先搜索` 


## dfs
```go
func getLeafSlice(root *TreeNode) []int {
	resSlice := []int{}
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		if node.Left == nil && node.Right == nil {
			resSlice = append(resSlice, node.Val)
		}
		helper(node.Right)
	}
	helper(root)
	return resSlice
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1, s2 := getLeafSlice(root1), getLeafSlice(root2)
	return reflect.DeepEqual(s1, s2)
}
```
>执行用时: 0 ms
内存消耗: 2.6 MB

暴力的方法就是先找出所有的叶子节点，然后再判断是否相等。

还有优化的方法就是只找出第一棵树的全部叶子节点，第二棵树遍历的时候就依次进行对比即可。但是这样还需要再维护一个记录当前遍历到多少个的数据。实现起来稍微麻烦一点。