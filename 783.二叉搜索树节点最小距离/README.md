# 783.二叉搜索树节点最小距离
[https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes](https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes) 
## 原题
给你一个二叉搜索树的根节点 `root` ，返回 **树中任意两不同节点值之间的最小差值** 。

 **注意：** 本题与 530：<a href="https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/">https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/</a> 相同

 
 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/05/bst1.jpg" style="width: 292px; height: 301px;" />
```

输入：root = [4,2,6,1,3]
输出：1

```
 **示例 2：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/05/bst2.jpg" style="width: 282px; height: 301px;" />
```

输入：root = [1,0,48,null,null,12,49]
输出：1

```
 

 **提示：** 
- 树中节点数目在范围 `[2, 100]` 内
-  `0 <= Node.val <= 10^5` 
 
**标签**
`树` `深度优先搜索` `递归` 


## 利用二叉搜索树的性质
```go
func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if res > num {
			res = num
		}
	}
	return res
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 1e6
	}
	res := min(minDiffInBST(root.Left), minDiffInBST(root.Right))
	lMax := maxLeft(root)
	if lMax != nil {
		res = min(res, root.Val-lMax.Val)
	}
	rMin := minRight(root)
	if rMin != nil {
		res = min(res, rMin.Val-root.Val)
	}
	return res
}

func maxLeft(root *TreeNode) *TreeNode {
	cur := root.Left
	if cur == nil {
		return nil
	}
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

func minRight(root *TreeNode) *TreeNode {
	cur := root.Right
	if cur == nil {
		return nil
	}
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}
```
>执行用时: 0 ms
内存消耗: 2.4 MB

思路就是利用二叉搜索树的性质，和一个数最接近的数一定在它的两边，即左子树的最右，和右子树的最左。此时就能够求出对这个节点的最小差值了。

题目要求是整个树的最小差值，因此需要对每个节点都求一次，先对左右子树求，再求自己的最小差值，最后返回里面的最小的差值。

## 中序遍历
```go
func minDiffInBST(root *TreeNode) int {
    ans, pre := math.MaxInt64, -1
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        if pre != -1 && node.Val-pre < ans {
            ans = node.Val - pre
        }
        pre = node.Val
        dfs(node.Right)
    }
    dfs(root)
    return ans
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/solution/er-cha-sou-suo-shu-jie-dian-zui-xiao-ju-8u87w/
// 来源：力扣（LeetCode）
```
>执行用时: 0 ms
内存消耗: 2.3 MB

这就是用了中序遍历的思路，和二叉搜索树的性质。左子树中序遍历的最后结果一定是左子树的最右节点，这就是实现了之前所说的左子树的最右。这里就判断 `ans = ans > node.Val - pre?node.Val - pre:ans` 能不能更新答案，并且更新 `pre` 节点的值为当前节点。因为右子树的遍历的第一个结果一定是右子树的最左。此时当访问到的时候一样能够减去 `pre` 节点。