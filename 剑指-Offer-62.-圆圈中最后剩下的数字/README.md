# 剑指 Offer 62. 圆圈中最后剩下的数字
[https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/](https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/)
## 约瑟夫环
```go
func lastRemaining(n int, m int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (res + m) % i
	}
	return res
}
```
>执行用时：8 ms, 在所有 Go 提交中击败了99.18%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了78.28%的用户


[这个题解](https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/solution/huan-ge-jiao-du-ju-li-jie-jue-yue-se-fu-huan-by-as/)讲得很清楚，主要是对于公式的推导是怎么来的讲得很透彻。
