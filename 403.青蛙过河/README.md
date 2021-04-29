# 403.青蛙过河
[https://leetcode-cn.com/problems/frog-jump](https://leetcode-cn.com/problems/frog-jump) 
## 原题
一只青蛙想要过河。 假定河流被等分为若干个单元格，并且在每一个单元格内都有可能放有一块石子（也有可能没有）。 青蛙可以跳上石子，但是不可以跳入水中。

给你石子的位置列表 `stones` （用单元格序号 **升序** 表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一块石子上）。

开始时， 青蛙默认已站在第一块石子上，并可以假定它第一步只能跳跃一个单位（即只能从单元格 1 跳至单元格 2 ）。

如果青蛙上一步跳跃了  `k` 个单位，那么它接下来的跳跃距离只能选择为  `k - 1` 、 `k` 或 `k + 1` 个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。

 

 **示例 1：** 

```

输入：stones = [0,1,3,5,6,8,12,17]
输出：true
解释：青蛙可以成功过河，按照如下方案跳跃：跳 1 个单位到第 2 块石子, 然后跳 2 个单位到第 3 块石子, 接着 跳 2 个单位到第 4 块石子, 然后跳 3 个单位到第 6 块石子, 跳 4 个单位到第 7 块石子, 最后，跳 5 个单位到第 8 个石子（即最后一块石子）。
```
 **示例 2：** 

```

输入：stones = [0,1,2,3,4,8,9,11]
输出：false
解释：这是因为第 5 和第 6 个石子之间的间距太大，没有可选的方案供青蛙跳跃过去。
```
 

 **提示：** 
-  `2 <= stones.length <= 2000` 
-  `0 <= stones[i] <= 2^31 - 1` 
-  `stones[0] == 0` 
 
**标签**
`动态规划` 


## 动规

### MLE
```go
func canCross(stones []int) bool {
	pre := make([][]int, len(stones))
	for i := range pre[1:] {
		pre[1+i] = []int{}
	}
	pre[0] = []int{-1}
	for i := range stones {
		for _, index := range pre[i] {
			k := 0
			if i != 0 {
				k = stones[i] - stones[index]
			}
			for t := k - 1; t <= k+1; t++ {
				if t < 1 {
					continue
				}
				targetStone := t + stones[i]
				targetIndex := sort.Search(len(stones), func(ii int) bool {
					return stones[ii] >= targetStone
				})
				if targetIndex == len(stones) || stones[targetIndex] != targetStone {
					continue
				}
				pre[targetIndex] = append(pre[targetIndex], i)
			}
		}
	}
	return len(pre[len(stones)-1]) > 0
}
```

思路就是每次把可能的结果都存上。如果最后能够到最后一格，则解法一定大于 0 。这里的 pre 记录的是可以跳到这个点的点的下标。

### 正解
```go
func canCross(stones []int) bool {
	n := len(stones)
	for i := 1; i < n; i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == n-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}
```
>执行用时: 12 ms
内存消耗: 8.5 MB

上一个解法之所以会 MLE 是因为最后可能会存在指数级的解法。而范围是 `0~2000`,$2000! = 3.31e5735$ 这么大肯定会爆内存。

因此另辟蹊径。可以知道第 `i` 个点和第 `j` 个点之间的距离差最大可以为 `j+1` 。若是比 `j-i` 还大，则无论如何都跳不过去。
>大概举个例子就是 0,1,3,6,10 每次跳的距离递增 1 。如果比这个最大的递增序列还大，则不可能跳到终点

于是用 `dp[i][k]` 记录第 `i` 个点是否能够由跳 `k` 个距离的点跳到。

可以知道能够到 `k` 的状态的点，这个时候记为`j` ，此时 `dp[j][k-1]`、`dp[j][k]`、`dp[j][k+1]`  只要有一者可以跳到，那么 `dp[i][k]` 就可以被 `j` 经过 `k` 跳到。当然了，还需要满足 `k = stones[i] - stones[j]`