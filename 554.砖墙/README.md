# 554.砖墙
[https://leetcode-cn.com/problems/brick-wall](https://leetcode-cn.com/problems/brick-wall) 
## 原题
你的面前有一堵矩形的、由 `n` 行砖块组成的砖墙。这些砖块高度相同（也就是一个单位高）但是宽度不同。每一行砖块的宽度之和应该相等。

你现在要画一条 **自顶向下** 的、穿过 **最少** 砖块的垂线。如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。 **你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。** 

给你一个二维数组 `wall` ，该数组包含这堵墙的相关信息。其中， `wall[i]` 是一个代表从左至右每块砖的宽度的数组。你需要找出怎样画才能使这条线 **穿过的砖块数量最少** ，并且返回 **穿过的砖块数量** 。

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2021/04/24/cutwall-grid.jpg" style="width: 493px; height: 577px;" />
```

输入：wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
输出：2

```
 **示例 2：** 

```

输入：wall = [[1],[1],[1]]
输出：3

```

 

 **提示：** 
-  `n == wall.length` 
-  `1 <= n <= 10^4` 
-  `1 <= wall[i].length <= 10^4` 
-  `1 <= sum(wall[i].length) <= 2 * 10^4` 
- 对于每一行 `i` ， `sum(wall[i])` 应当是相同的
-  `1 <= wall[i][j] <= 2^31 - 1` 
 
**标签**
`哈希表` 


## 伪差分数组
```go
func leastBricks(wall [][]int) (res int) {
	brickStart := map[int]int{}

	for _, row := range wall {
		pre := 0
		for _, brick := range row {
			brickStart[pre]++
			pre += brick
		}
	}

	n := len(wall)
	res = n
	for k, v := range brickStart {
		if k > 0 {
			res = min(res, n-v)
		}
	}

	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```
>执行用时: 28 ms
内存消耗: 7.1 MB

和差分数组的思路差不多。不过这个只记录了有多少块砖在这个点开始出现。因为如果边缘也算的话无论如何都是 `len(wall)` 个长度。因此减去在这个点开始的砖块数量即可知道横跨这个点的砖块有多少。
