# 815.公交路线
[https://leetcode-cn.com/problems/bus-routes](https://leetcode-cn.com/problems/bus-routes) 
## 原题
给你一个数组 `routes` ，表示一系列公交线路，其中每个 `routes[i]` 表示一条公交线路，第 `i` 辆公交车将会在上面循环行驶。
- 例如，路线 `routes[0] = [1, 5, 7]` 表示第 `0` 辆公交车会一直按序列 `1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ...` 这样的车站路线行驶。
现在从 `source` 车站出发（初始时不在公交车上），要前往 `target` 车站。 期间仅可乘坐公交车。

求出 **最少乘坐的公交车数量** 。如果不可能到达终点车站，返回 `-1` 。

 

 **示例 1：** 

```

输入：routes = [[1,2,7],[3,6,7]], source = 1, target = 6
输出：2
解释：最优策略是先乘坐第一辆公交车到达车站 7 , 然后换乘第二辆公交车到车站 6 。 

```
 **示例 2：** 

```

输入：routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
输出：-1

```
 

 **提示：** 
-  `1 <= routes.length <= 500` .
-  `1 <= routes[i].length <= 10^5` 
-  `routes[i]` 中的所有值 **互不相同** 
-  `sum(routes[i].length) <= 10^5` 
-  `0 <= routes[i][j] < 10^6` 
-  `0 <= source, target < 10^6` 
 
**标签**
`广度优先搜索` `数组` `哈希表` 


## 
```go
func numBusesToDestination(routes [][]int, source, target int) int {
    if source == target {
        return 0
    }

    n := len(routes)
    edge := make([][]bool, n)
    for i := range edge {
        edge[i] = make([]bool, n)
    }
    rec := map[int][]int{}
    for i, route := range routes {
        for _, site := range route {
            for _, j := range rec[site] {
                edge[i][j] = true
                edge[j][i] = true
            }
            rec[site] = append(rec[site], i)
        }
    }

    dis := make([]int, n)
    for i := range dis {
        dis[i] = -1
    }
    q := []int{}
    for _, bus := range rec[source] {
        dis[bus] = 1
        q = append(q, bus)
    }
    for len(q) > 0 {
        x := q[0]
        q = q[1:]
        for y, b := range edge[x] {
            if b && dis[y] == -1 {
                dis[y] = dis[x] + 1
                q = append(q, y)
            }
        }
    }

    ans := math.MaxInt32
    for _, bus := range rec[target] {
        if dis[bus] != -1 && dis[bus] < ans {
            ans = dis[bus]
        }
    }
    if ans < math.MaxInt32 {
        return ans
    }
    return -1
}
```
>执行用时: 108 ms
内存消耗: 21.6 MB

原理就是 `bfs` 但是提前建了图，把站点之间的 `bfs` 转化为了公交线路之间的 `bfs` 。相当于降低了复杂度，不然的会会在最后一个数据 `TLE` 。

原理不难，但是没有想到优化建图这种方法，还是经验不够。