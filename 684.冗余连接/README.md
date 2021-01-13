# 684.冗余连接
[https://leetcode-cn.com/problems/redundant-connection/](https://leetcode-cn.com/problems/redundant-connection/) 
## 原题
在本问题中, 树指的是一个连通且无环的**无向** 图。

输入一个图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N) 的树及一条附加的边构成。附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。

结果图是一个以`边`组成的二维数组。每一个`边`的元素是一对`[u, v]` ，满足 `u < v`，表示连接顶点`u` 和`v`的**无向** 图的边。

返回一条可以删去的边，使得结果图是一个有着N个节点的树。如果有多个答案，则返回二维数组中最后出现的边。答案边 `[u, v]` 应满足相同的格式 `u < v`。

**示例 1：** 

```
输入: [[1,2], [1,3], [2,3]]
输出: [2,3]
解释: 给定的无向图为:
  1
 / \
2 - 3

```
**示例 2：** 

```
输入: [[1,2], [2,3], [3,4], [1,4], [1,5]]
输出: [1,4]
解释: 给定的无向图为:
5 - 1 - 2
    |   |
    4 - 3

```
**注意:** 
- 输入的二维数组大小在 3 到 1000。
- 二维数组中的整数在1到N之间，其中N是输入数组的大小。
**更新(2017-09-26):** <br>
我们已经重新检查了问题描述及测试用例，明确图是**无向**图。对于有向图详见<a href="https://leetcodechina.com/problems/redundant-connection-ii/description/">冗余连接II</a>。 对于造成任何不便，我们深感歉意。

 
**标签**
`树` `并查集` `图` 


## 
```go
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	fa := make([]int, n+1)
	rank := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) {
		fx, fy := find(x), find(y)
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx
	}

	res := make([]int, 2)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if find(x) == find(y) {
			res = edge
		}
		merge(x, y)
	}

	return res
}
```
>执行用时: 4 ms
内存消耗: 3.1 MB

草 怎么又是并查集。

思路很清楚，先搞清楚怎么判断是否冗余简单地说就是遍历 **edge** ，看现在更新的这两个点是不是已经是同一个集合的了，如果是的话就说明这条边是冗余的（因为没这条边还是在同一个图里面）。

这个题不同的地方就是在合并之前先看两个是不是已经是同一集合的，如果是的话就记录下来。

题目要求返回最后的，其实返回第一个找到的就可以了，因为只有一条附加边，所以直接返回问题也不大。

优化一下的话就是下面这样，我顺便还改了一下 **merge** ，这样可以实现节省一点点时间
```go
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	fa := make([]int, n+1)
	rank := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx
	}

	res := make([]int, 2)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if find(x) == find(y) {
			res = edge
			return res
		}
		merge(x, y)
	}

	return nil
}
```

修改后时间和内存占用都没变，应该是最优了。0 ms 的那些应该是早期用的是小数据。