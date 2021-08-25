# 787.K 站中转内最便宜的航班
[https://leetcode-cn.com/problems/cheapest-flights-within-k-stops](https://leetcode-cn.com/problems/cheapest-flights-within-k-stops) 
## 原题
有 `n` 个城市通过一些航班连接。给你一个数组 `flights` ，其中 `flights[i] = [from<sub>i</sub>, to<sub>i</sub>, price<sub>i</sub>]` ，表示该航班都从城市 `from<sub>i</sub>` 开始，以价格 `to<sub>i</sub>` 抵达 `price<sub>i</sub>` 。

现在给定所有的城市和航班，以及出发城市 `src` 和目的地 `dst` ，你的任务是找到出一条最多经过 `k` 站中转的路线，使得从 `src` 到 `dst` 的 **价格最便宜** ，并返回该价格。 如果不存在这样的路线，则输出 `-1` 。

 

 **示例 1：** 

```

输入: 
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
输出: 200
解释: 
城市航班图如下
<img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/02/16/995.png" style="height: 180px; width: 246px;" />

从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
```
 **示例 2：** 

```

输入: 
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
输出: 500
解释: 
城市航班图如下
<img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/02/16/995.png" style="height: 180px; width: 246px;" />

从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。
```
 

 **提示：** 
-  `1 <= n <= 100` 
-  `0 <= flights.length <= (n * (n - 1) / 2)` 
-  `flights[i].length == 3` 
-  `0 <= from<sub>i</sub>, to<sub>i</sub> < n` 
-  `from<sub>i</sub> != to<sub>i</sub>` 
-  `1 <= price<sub>i</sub> <= 10^4` 
- 航班没有重复，且不存在自环
-  `0 <= src, dst, k < n` 
-  `src != dst` 
 
**标签**
`深度优先搜索` `广度优先搜索` `图` `动态规划` `最短路` `堆（优先队列）` 


## 
```go
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := make([][]int, n)

	for i := range cost {
		cost[i] = make([]int, k+2)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt64 / 2
		}
	}
	cost[src][0] = 0
	for i := 1; i <= k+1; i++ {
		for _, flight := range flights {
			to := flight[1]
			from := flight[0]
			cost[to][i] = min(cost[to][i], cost[from][i-1]+flight[2])
		}
	}
	res := math.MaxInt64 >> 1
	for _, d := range cost[dst][1:] {
		res = min(res, d)
	}
	if res == math.MaxInt64/2 {
		return -1
	}

	return res
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if num < res {
			res = num
		}
	}
	return res
}
```
>执行用时：12 ms
内存消耗：4.9 MB

思路就是按照到每个点的次数进行遍历，然后不断更新从源店到它的最短距离。