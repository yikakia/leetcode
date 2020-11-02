# 剑指 Offer 43. 1～n 整数中 1 出现的次数
[https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/](https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/)
## 规律分析
```go
func countDigitOne(n int) int {
	r := 0
	cur := n % 10
	hi, lo := n/10, 0
	digit := 1
	for cur != 0 || hi != 0 {
		if cur == 0 {
			r += hi * digit
		} else if cur == 1 {
			r += hi*digit + lo + 1
		} else {
			r += (hi + 1) * digit
		}
		lo += cur * digit
		cur = hi % 10
		hi /= 10
		digit *= 10
	}
	return r
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了43.48%的用户

简单地说就是找规律。把数字拆分成 hi cur lo 三个部分。cur表示当前要求的位。
- 当 cur == 1 的时候
    简单地说，就是无论 hi 和 lo 部分取什么值都是出现1。
    那么就很简单了，把 hi * (lo 的位数)，再加上 lo 再加1（因为lo为全0时也可以满足条件）

- 当 cur == 0 的时候
    还是对当前位进行分析。如果要出现 1 那么只能是当前位为 1 。而当前位小于 1 。这个时候就是 hi * (lo 的位数)

- 当 cur > 2 的时候
    这个时候当前位大于 1 ，让高位和低位能够顶格，这个时候就是 (hi+1) * (lo 的位数)