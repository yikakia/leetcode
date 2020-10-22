# 剑指 Offer 15. 二进制中1的个数
[https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/](https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/)

```go
func hammingWeight(num uint32) int {
	r := 0
	for num != 0 {
		r += int(num & 1)
		num = num >> 1
	}
	return r
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了66.00%的用户

两种思路，一个是辗转相除法，把它转为二进制数的同时获取`1`的个数。

还有一种思路就是我这种，进行位运算，直接看成二进制数，获取`1`的个数