# 1232.缀点成线
[https://leetcode-cn.com/problems/check-if-it-is-a-straight-line/](https://leetcode-cn.com/problems/check-if-it-is-a-straight-line/) 
## 原题
在一个 XY 坐标系中有一些点，我们用数组 `coordinates` 来分别记录它们的坐标，其中 `coordinates[i] = [x, y]` 表示横坐标为 `x`、纵坐标为 `y` 的点。

请你来判断，这些点是否在该坐标系中属于同一条直线上，是则返回 `true`，否则请返回 `false`。

 

**示例 1：** 

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/10/19/untitled-diagram-2.jpg" style="height: 336px; width: 336px;">

```
输入：coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
输出：true

```
**示例 2：** 

**<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/10/19/untitled-diagram-1.jpg" style="height: 336px; width: 348px;">** 

```
输入：coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
输出：false
```
 

**提示：** 
- `2 <= coordinates.length <= 1000`
- `coordinates[i].length == 2`
- `-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4`
- `coordinates` 中不含重复的点
 
**标签**
`几何` `数组` `数学` 


## 
```go
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}
	c0, c1 := coordinates[0], coordinates[1]
	x0, y0 := c0[0], c0[1]
	x1, y1 := c1[0], c1[1]
	for _, c := range coordinates[2:] {
		x2, y2 := c[0], c[1]
		if (y1-y0)*(x2-x0) != (y2-y0)*(x1-x0) {
			return false
		}
	}

	return true
}
```
>执行用时: 4 ms
内存消耗: 3.6 MB

简单的公式推一下就可以了
