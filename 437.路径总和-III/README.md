# 437.路径总和 III
[https://leetcode-cn.com/problems/path-sum-iii](https://leetcode-cn.com/problems/path-sum-iii) 
## 原题
给定一个二叉树的根节点 `root`  ，和一个整数 `targetSum` ，求该二叉树里节点值之和等于 `targetSum` 的 **路径** 的数目。

 **路径** 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

 

 **示例 1：** 

<img src="https://assets.leetcode.com/uploads/2021/04/09/pathsum3-1-tree.jpg" style="width: 452px; " />

```

输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。

```
 **示例 2：** 

```

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3

```
 

 **提示:** 
- 二叉树的节点个数的范围是 `[0,1000]` 
- <meta charset="UTF-8" /> `-10^9 <= Node.val <= 10^9`  
-  `-1000 <= targetSum <= 1000`  
 
**标签**
`树` `深度优先搜索` `二叉树` 


## 深搜
```go
type TreeNodeWithCount struct {
	Val   int
	Left  *TreeNodeWithCount
	Right *TreeNodeWithCount
	count map[int]int
}

var res = 0
var _targetSum = 0

func pathSum(root *TreeNode, targetSum int) int {
	res = 0
	_targetSum = targetSum
	ret := conv(root)
	pathSum2(ret)
	return res
}

func conv(root *TreeNode) *TreeNodeWithCount {
	if root == nil {
		return nil
	}

	ret := &TreeNodeWithCount{}
	ret.count = map[int]int{}

	ret.count[root.Val]++
	ret.Left = conv(root.Left)
	ret.Right = conv(root.Right)
	ret.Val = root.Val

	return ret
}

func pathSum2(root *TreeNodeWithCount) {
	if root == nil {
		return
	}
	pathSum2(root.Left)
	pathSum2(root.Right)
	res += root.count[_targetSum]
	add(root, root.Left)
	add(root, root.Right)
}

func add(root, son *TreeNodeWithCount) {
	if son != nil {
		for k, v := range son.count {
			root.count[k+root.Val] += v
			if k+root.Val == _targetSum {
				res += v
			}
		}
	}
}
```
>执行用时：68 ms
内存消耗：32 MB

思路就是先看自己满不满足，满足的话再看自己加上孩子满不满足条件。

不过这么写会有内存占用过大的问题，主要是我这么写保存了大量没必要保存的数据。虽然比起便利来说理论上时间会更快，因为保存了多种情况下的数据，但是实际上并不能加速，应该是 map 会频繁进行扩容导致的。

下面的是优化版本

### 优化版
```go
func pathSum(root *TreeNode, targetSum int) int {
    if root==nil{
        return 0
    }
    var res int
    res = findPath(root, targetSum)
    res += pathSum(root.Left, targetSum)
    res += pathSum(root.Right, targetSum)
    return res
}
func findPath(root *TreeNode, target int) int{
    if root==nil{
        return 0
    }
    var ans int
    if root.Val==target{
        ans += 1
    }
    ans += findPath(root.Left, target-root.Val)
    ans += findPath(root.Right, target-root.Val)
    return ans
}
```