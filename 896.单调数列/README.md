# 896.单调数列
[https://leetcode-cn.com/problems/monotonic-array](https://leetcode-cn.com/problems/monotonic-array) 
## 原题
如果数组是单调递增或单调递减的，那么它是 *单调的* 。

如果对于所有 `i <= j` ， `A[i] <= A[j]` ，那么数组 `A` 是单调递增的。 如果对于所有 `i <= j` ， `A[i]> = A[j]` ，那么数组 `A` 是单调递减的。

当给定的数组 `A` 是单调数组时返回 `true` ，否则返回 `false` 。

 
 **示例 1：** 

```
输入：[1,2,2,3]
输出：true

```
 **示例 2：** 

```
输入：[6,5,4,4]
输出：true

```
 **示例 3：** 

```
输入：[1,3,2]
输出：false

```
 **示例 4：** 

```
输入：[1,2,4,5]
输出：true

```
 **示例 5：** 

```
输入：[1,1,1]
输出：true

```
 

 **提示：** 
-  `1 <= A.length <= 50000` 
-  `-100000 <= A[i] <= 100000` 
 
**标签**
`数组` 


## 置标志位
```go
func isMonotonic(A []int) bool {
	n := len(A)
	if n < 3 {
		return true
	}
	flag := -1
	for i := 1; i < n; i++ {
		switch {
		case A[i] > A[i-1]:
			if flag == 2 {
				return false
			}
			flag = 1
		case A[i] < A[i-1]:
			if flag == 1 {
				return false
			}
			flag = 2
		default:
		}
	}
	return true
}
```
>执行用时: 68 ms
内存消耗: 8.1 MB

简单地说就是用一个 `flag` 表示现在存在的状态是增还是减。

### 二进制状态压缩 (负优化)
```go

const (
	Inc = iota + 1
	Dec
)

func isMonotonic(A []int) bool {
	n := len(A)
	if n < 3 {
		return true
	}
	flag := 0
	for i := 1; i < n; i++ {
		switch {
		case A[i] > A[i-1]:
			flag |= Inc
		case A[i] < A[i-1]:
			flag |= Dec
		default:
		}
		if flag == Inc|Dec {
			return false
		}
	}
	return true
}
```
>执行用时: 80 ms
内存消耗: 8.1 MB

想了下用二进制状态表示会不会快一点。结果表明并不会。不过感觉这个用时应该是正常的波动吧。