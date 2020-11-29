# 剑指 Offer 66. 构建乘积数组
[https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/](https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/)
## 用位运算模拟除法
```go
func constructArr(a []int) []int {
	all := 1
	hasZero := false
	hasNZero := false
	for _, v := range a {
		if v != 0 {
			all *= v
		} else {
			if hasZero {
				hasNZero = true
				break
			}
			hasZero = true
		}
	}
	b := make([]int, len(a))
	if !hasNZero {
		if hasZero {
			for i := range b {
				if a[i] != 0 {
					b[i] = 0
				} else {
					b[i] = all
				}
			}
		} else {
			for i := range b {
				b[i], _ = div(all, a[i])
			}
		}
	}
	return b
}

// div 是进行 a/b 的除法运算，返回一个商和余数
func div(a, b int) (int, int) {
	// 被除数
	dividend := a
	if a <= 0 {
		dividend = -a
	}
	// 除数
	divisor := b
	if b <= 0 {
		divisor = -b
	}
	// 商
	quotient := 0
	// 余数
	remainder := 0
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			quotient = quotient + 1<<i
			dividend = dividend - divisor<<i
		}
	}
	if a^b < 0 {
		quotient = -quotient
	}
	if b > 0 {
		remainder = dividend
	} else {
		remainder = -dividend
	}
	return quotient, remainder
}
```
>执行用时：36 ms, 在所有 Go 提交中击败了14.43%的用户
内存消耗：8.1 MB, 在所有 Go 提交中击败了52.52%的用户

模拟除法其实也算除法。所以还得用别的方法。

## 两遍遍历
```go
func constructArr(a []int) []int {
	n := len(a)
	b := make([]int, n)
	tmp := 1
	for i := range a {
		b[i] = tmp
		tmp *= a[i]
	}
	tmp = 1
	for i := n - 1; i >= 0; i-- {
		b[i] *= tmp
		tmp *= a[i]
	}
	return b
}
```
>执行用时：28 ms, 在所有 Go 提交中击败了78.11%的用户
内存消耗：8.1 MB, 在所有 Go 提交中击败了52.52%的用户

简单地说就是先从左往右乘，再从右往左乘。