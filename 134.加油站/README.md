# 134. 加油站
[https://leetcode-cn.com/problems/gas-station/](https://leetcode-cn.com/problems/gas-station/) 
## 暴力
```go
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := range gas {
		cur := gas[i] - cost[i]
		for j := 1; j < n+1; j++ {
			pos := (i + j) % n
			if cur < 0 {
				break
			}
			cur += gas[pos] - cost[pos]
		}
		if cur >= 0 {
			return i
		}
	}
	return -1
}
```
>执行用时: 196 ms
内存消耗: 3 MB

就是纯粹模拟的暴力法。明显耗时太长

## 优化遍历方式
```go
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; {
		cnt := 0
		totalGas, totalCost := 0, 0
		for ; totalGas >= totalCost && cnt < n; cnt++ {
			pos := (i + cnt) % n
			totalGas += gas[pos]
			totalCost += cost[pos]
		}
		if totalGas >= totalCost && cnt == n {
			return i
		}
		i += cnt
	}
	return -1
}
```
>执行用时: 0 ms
内存消耗: 2.9 MB

参考这个[题解](https://leetcode-cn.com/problems/gas-station/solution/shi-yong-tu-de-si-xiang-fen-xi-gai-wen-ti-by-cyayc/)的图片![参考](https://pic.leetcode-cn.com/98ee6782654518e1a33852e99825f1537869a542ee26738cf02d5fb6f0f0a899-%E6%97%A0%E6%A0%87%E9%A2%98.png)


把绿色看作是 gas 橙色看作是 cost 黑色看作是 sum(gas-cost) ，很简单的我们可以知道，不论从哪个点开始，这个黑色的折线的形状是不会变的，只会在坐标轴上下移动。那么我们要做的就是不断调整开始的位置，直到跑了一圈之后，所有的黑色点都在 x 轴上方就可以了。为了避免最后一个退出的情况，所以还要再判断最后一个满不满足。