# 452. 用最少数量的箭引爆气球
[https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/) 
## 
```go
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	res := 1
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	for _, point := range points {
		if point[0] > maxRight {
			maxRight = point[1]
			res++
		}
	}
	return res
}
```
>执行用时: 88 ms
内存消耗: 7.3 MB

按照右端点进行升序排列，然后在遍历的时候更新最大的右端点。当遍历到的区间在最大右端点的右边（即不被包括）的时候，就把当前区间的右端点记为最大右端点。

在执行更新操作的时候还要让结果加1。结果的初始值应该为1，因为实际的运行时如果没有进行更新操作也应该射一只箭。不过如果数组为空，就不应该射箭了。所以先提前判断是否为空，不为空再执行射箭的流程。