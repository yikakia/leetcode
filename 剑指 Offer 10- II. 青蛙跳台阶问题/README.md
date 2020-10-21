# 剑指 Offer 10- II. 青蛙跳台阶问题
[https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/](https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/)

## 水题
```go
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	a, b, r := 1, 1, 0
	const max = 1000000007
	for i := 2; i <= n; i++ {
		r = (a + b) % max
		a = b
		b = r
	}
	return r
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了67.75%的用户

和前面的斐波那契数列有区别么？初始值不同而已。水题.