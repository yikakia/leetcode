# 1584.连接所有点的最小费用
[https://leetcode-cn.com/problems/min-cost-to-connect-all-points/](https://leetcode-cn.com/problems/min-cost-to-connect-all-points/) 
## 原题
给你一个`points` 数组，表示 2D 平面上的一些点，其中 `points[i] = [x<sub>i</sub>, y<sub>i</sub>]` 。

连接点 `[x<sub>i</sub>, y<sub>i</sub>]` 和点 `[x<sub>j</sub>, y<sub>j</sub>]` 的费用为它们之间的 **曼哈顿距离**  ：`|x<sub>i</sub> - x<sub>j</sub>| + |y<sub>i</sub> - y<sub>j</sub>|` ，其中 `|val|` 表示 `val` 的绝对值。

请你返回将所有点连接的最小总费用。只有任意两点之间 **有且仅有**  一条简单路径时，才认为所有点都已连接。

 

**示例 1：** 

<img alt="" src="https://assets.leetcode.com/uploads/2020/08/26/d.png" style="height:268px; width:214px" />

```

输入：points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
输出：20
解释：
<img alt="" src="https://assets.leetcode.com/uploads/2020/08/26/c.png" style="height:268px; width:214px" />
我们可以按照上图所示连接所有点得到最小总费用，总费用为 20 。
注意到任意两个点之间只有唯一一条路径互相到达。

```
**示例 2：** 

```

输入：points = [[3,12],[-2,5],[-4,1]]
输出：18

```
**示例 3：** 

```

输入：points = [[0,0],[1,1],[1,0],[-1,1]]
输出：4

```
**示例 4：** 

```

输入：points = [[-1000000,-1000000],[1000000,1000000]]
输出：4000000

```
**示例 5：** 

```

输入：points = [[0,0]]
输出：0

```
 

**提示：** 
- `1 <= points.length <= 1000`
- `-10^6 <= x<sub>i</sub>, y<sub>i</sub> <= 10^6`
- 所有点 `(x<sub>i</sub>, y<sub>i</sub>)` 两两不同。
 
**标签**
`并查集` 


## 最小生成树 
### Kruskal 算法
```go
func minCostConnectPoints(points [][]int) int {
	type edge struct {
		a, b int
		dis  int
	}
	n := len(points)
	edges := []edge{}
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range points {
		pa := points[i]
		fa[i] = i
		rank[i] = 1
		for j := i + 1; j < n; j++ {
			pb := points[j]
			edges = append(edges, edge{a: i, b: j, dis: calDis(pa, pb)})
		}
	}
	// for _, e := range edges {
	// 	fmt.Printf("%+v\n", e)
	// }
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

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dis < edges[j].dis
	})

	res := 0
	for _, e := range edges {
		fx, fy := find(e.a), find(e.b)
		if fx == fy {
			continue
		}
		res += e.dis
		merge(e.a, e.b)
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func calDis(a, b []int) int {
	if len(a) != 2 || len(b) != 2 {
		return 0
	}
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}
```
>执行用时: 452 ms
内存消耗: 39.3 MB

就是简单的 Kruskal 算法 。

### Prim 算法
```go
func minCostConnectPoints(points [][]int) int {
	type Point struct {
		local []int
		dis   int
	}
	n := len(points)
	_points := make([]Point, n)
	for i := range points {
		_points[i].local = points[i]
		_points[i].dis = calDis(points[0], points[i])
	}

	notvisited := map[int]bool{}
	for i := range _points[1:] {
		notvisited[i+1] = true
	}

	res := 0
	for range _points[1:] {
		nextPoint := -1
		minDist := math.MaxInt32
		for index := range notvisited {
			pa := _points[index]
			if pa.dis < minDist {
				nextPoint = index
				minDist = pa.dis
			}
		}
		res += minDist
		delete(notvisited, nextPoint)
		for i := range notvisited {
			p := _points[i]
			_points[i].dis = min(p.dis, calDis(p.local, points[nextPoint]))
		}
	}

	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func calDis(a, b []int) int {
	if len(a) != 2 || len(b) != 2 {
		return 0
	}
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}
```
>执行用时: 88 ms
内存消耗: 4.4 MB

稠密图果然还是 **Prim** 算法能够更快啊，特别是这种强连通图，时间和空间都是 **Kruskal** 算法的$\frac{1}{10}$特别是我还专门另外开结构体储存了点的坐标和距已知的点集的最小距离。

总的来说就是维护一个未访问过的点集，每次找到其中的最小的距离，然后更新对应的距离即可。


那我不另外开结构体是不是更加省空间？

### Prim 的小修改版
```go
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dis := make([]int, n)
	for i := range points {
		dis[i] = calDis(points[0], points[i])
	}

	notvisited := map[int]bool{}
	for i := range dis[1:] {
		notvisited[i+1] = true
	}

	res := 0
	for range dis[1:] {
		nextPoint := -1
		minDist := math.MaxInt32
		for index := range notvisited {
			if dis[index] < minDist {
				nextPoint = index
				minDist = dis[index]
			}
		}
		res += minDist
		delete(notvisited, nextPoint)
		for i := range notvisited {
			dis[i] = min(dis[i], calDis(points[i], points[nextPoint]))
		}
	}

	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func calDis(a, b []int) int {
	if len(a) != 2 || len(b) != 2 {
		return 0
	}
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}
```
>执行用时: 80 ms
内存消耗: 4.2 MB

区别并不大啊。

### Prim 不用 map 
```go
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dis := make([]int, n)
	for i := range points {
		dis[i] = calDis(points[0], points[i])
	}

	notvisited := make([]bool, n)
	for i := range dis[1:] {
		notvisited[i+1] = true
	}

	res := 0
	for range dis[1:] {
		nextPoint := -1
		minDist := math.MaxInt32
		for index, v := range notvisited {
			if v == false {
				continue
			}
			if dis[index] < minDist {
				nextPoint = index
				minDist = dis[index]
			}
		}
		res += minDist
		notvisited[nextPoint] = false
		for i, v := range notvisited {
			if v == false {
				continue
			}
			dis[i] = min(dis[i], calDis(points[i], points[nextPoint]))
		}
	}

	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func calDis(a, b []int) int {
	if len(a) != 2 || len(b) != 2 {
		return 0
	}
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}
```
>执行用时: 20 ms
内存消耗: 3.7 MB

果然，在用下标做键的情况下还是用数组要快得多，即使在需要遍历的情况下也是如此。只有没有下标的情况下用哈希表才比较好。