# 剑指 Offer 14- II. 剪绳子 II
[https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/](https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/)

```go
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	r := 1
	for n > 4 {
		r = r * 3 % 1000000007
		n -= 3
	}
	return r * n % 1000000007
}
```
这次因为要用到了取余，所以之前的先乘法后除法的行不通了。因为不一定可以被`3`整除了。这个地方应该注意。