# 剑指 Offer 16. 数值的整数次方
[https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/](https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/)
```go
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var r float64 = 1
	if n < 0 {
		x = 1 / x
		n *= -1
	}
	for n > 0 {
		if n%2 == 1 {
			r *= x
		}
		x *= x
		n /= 2
	}
	return r
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了35.06%的用户

有坑，直接计算n次方会超时,因此应该用到快速幂。快速幂的思想是把多少次幂n看成是二进制序列，那么n可以看成是`n=k1*2^1+k2*2^2+……+kp*a^2`，其中`p`代表n的二进制位数，`kp`代表它在这个二进位上的值。那么原式`x^n`就可以转化为`[x^(k1*2^1)]*[x^(k2*2^2)]*……*[x^(kp*2^p)]`这个时候时间复杂度就大大降低了。