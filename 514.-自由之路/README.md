# 514. 自由之路
[https://leetcode-cn.com/problems/freedom-trail/](https://leetcode-cn.com/problems/freedom-trail/)
## 动态规划
```go
func findRotateSteps(ring string, key string) int {
	const inf = math.MaxInt32 / 2

	pos := make(map[byte][]int)
	for i := range ring {
		pos[ring[i]] = append(pos[ring[i]], i)
	}

	m := len(ring)
	n := len(key)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	for _, p := range pos[key[0]] {
		dp[0][p] = min(p, m-p) + 1
	}

	for i := 1; i < n; i++ {
		for _, j := range pos[key[i]] {
			for _, k := range pos[key[i-1]] {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), m-abs(j-k))+1)
			}
		}
	}
	return min(dp[n-1]...)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
```
>执行用时：12 ms, 在所有 Go 提交中击败了45.45%的用户
内存消耗：6.7 MB, 在所有 Go 提交中击败了27.27%的用户

有点像最短路径的求法，不断地松弛能够到这个点的边。不过这里是松弛的是离得最近的点。另外就是移动可以是双向的，所以要判断两种情况。