# 1030. 距离顺序排列矩阵单元格
[https://leetcode-cn.com/problems/matrix-cells-in-distance-order/](https://leetcode-cn.com/problems/matrix-cells-in-distance-order/) 
## 排序
```go
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, R*C)
	for i := range res {
		res[i] = []int{i % R, i / R}
	}
	sort.Slice(res, func(i, j int) bool {
		a := abs(res[i][0]-r0) + abs(res[i][1]-c0)
		b := abs(res[j][0]-r0) + abs(res[j][1]-c0)
		return a < b
	})
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

```
>执行用时: 20 ms
内存消耗: 6.9 MB

生成序列然后排序。
