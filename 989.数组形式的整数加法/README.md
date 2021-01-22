# 989.数组形式的整数加法
[https://leetcode-cn.com/problems/add-to-array-form-of-integer/](https://leetcode-cn.com/problems/add-to-array-form-of-integer/) 
## 原题
对于非负整数 `X` 而言，*`X`* 的*数组形式*是每位数字按从左到右的顺序形成的数组。例如，如果 `X = 1231`，那么其数组形式为 `[1,2,3,1]`。

给定非负整数 `X` 的数组形式 `A`，返回整数 `X+K` 的数组形式。

 
**示例 1：** 

```
输入：A = [1,2,0,0], K = 34
输出：[1,2,3,4]
解释：1200 + 34 = 1234

```
**示例 2：** 

```
输入：A = [2,7,4], K = 181
输出：[4,5,5]
解释：274 + 181 = 455

```
**示例 3：** 

```
输入：A = [2,1,5], K = 806
输出：[1,0,2,1]
解释：215 + 806 = 1021

```
**示例 4：** 

```
输入：A = [9,9,9,9,9,9,9,9,9,9], K = 1
输出：[1,0,0,0,0,0,0,0,0,0,0]
解释：9999999999 + 1 = 10000000000

```
 

**提示：** 
- `1 <= A.length <= 10000`
- `0 <= A[i] <= 9`
- `0 <= K <= 10000`
- 如果 `A.length > 1`，那么 `A[0] != 0`
 
**标签**
`数组` 


## 
```go

func addToArrayForm(A []int, K int) []int {
	if K == 0 {
		return A
	}
	ks := []int{}
	for K != 0 {
		ks = append(ks, K%10)
		K /= 10
	}
	reverse(A)
	if len(ks) > len(A) {
		A, ks = ks, A
	}
	A = append(A, 0)
	overflow := 0
	for i := range A {
		if i < len(ks) {
			A[i] += ks[i]
		} else if overflow == 0 {
			break
		}
		A[i] += overflow
		if A[i] > 9 {
			overflow = 1
			A[i] %= 10
		} else {
			overflow = 0
		}
	}
	if A[len(A)-1] == 0 {
		A = A[:len(A)-1]
	}
	reverse(A)
	return A
}

func reverse(a []int) {
	n := len(a)
	end := n / 2
	for i := 0; i < end; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

```
>执行用时: 28 ms
内存消耗: 6.6 MB

就是模拟运算，最后来个判断即可
