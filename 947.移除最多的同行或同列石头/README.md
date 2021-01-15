# 947.移除最多的同行或同列石头
[https://leetcode-cn.com/problems/most-stones-removed-with-same-row-or-column/](https://leetcode-cn.com/problems/most-stones-removed-with-same-row-or-column/) 
## 原题
`n` 块石头放置在二维平面中的一些整数坐标点上。每个坐标点上最多只能有一块石头。

如果一块石头的 **同行或者同列**  上有其他石头存在，那么就可以移除这块石头。

给你一个长度为 `n` 的数组 `stones` ，其中 `stones[i] = [x<sub>i</sub>, y<sub>i</sub>]` 表示第 `i` 块石头的位置，返回 **可以移除的石子**  的最大数量。

 

**示例 1：** 

```

输入：stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
输出：5
解释：一种移除 5 块石头的方法如下所示：
1. 移除石头 [2,2] ，因为它和 [2,1] 同行。
2. 移除石头 [2,1] ，因为它和 [0,1] 同列。
3. 移除石头 [1,2] ，因为它和 [1,0] 同行。
4. 移除石头 [1,0] ，因为它和 [0,0] 同列。
5. 移除石头 [0,1] ，因为它和 [0,0] 同行。
石头 [0,0] 不能移除，因为它没有与另一块石头同行/列。
```
**示例 2：** 

```

输入：stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
输出：3
解释：一种移除 3 块石头的方法如下所示：
1. 移除石头 [2,2] ，因为它和 [2,0] 同行。
2. 移除石头 [2,0] ，因为它和 [0,0] 同列。
3. 移除石头 [0,2] ，因为它和 [0,0] 同行。
石头 [0,0] 和 [1,1] 不能移除，因为它们没有与另一块石头同行/列。
```
**示例 3：** 

```

输入：stones = [[0,0]]
输出：0
解释：[0,0] 是平面上唯一一块石头，所以不可以移除它。
```
 

**提示：** 
- `1 <= stones.length <= 1000`
- `0 <= x<sub>i</sub>, y<sub>i</sub> <= 10^4`
- 不会有两块石头放在同一个坐标点上
 
**标签**
`深度优先搜索` `并查集` 


## 
```go
func removeStones(stones [][]int) int {
	fa, rank := map[int]int{}, map[int]int{}
	var find func(int) int
	find = func(x int) int {
		if _, has := fa[x]; !has {
			fa[x], rank[x] = x, 1
		}
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	union := func(x, y int) {
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

	for _, p := range stones {
		union(p[0], p[1]+10001)
	}
	ans := len(stones)
	for x, fx := range fa {
		if x == fx {
			ans--
		}
	}
	return ans
}
```
>执行用时: 24 ms
内存消耗: 6.8 MB

这道题就是并查集，但是做了个处理，把横纵坐标进行了映射，一个点`a`的横坐标全部映射到纵坐标上，这样当遇到了同样的横坐标的点`b`的时候，要做的就是把点`a`的纵坐标映射到点`b`的纵坐标上。这样最后要找的就是有几个集合，把总石头数减去集合数（最后留下来的石头数），就是拿去的石头数。