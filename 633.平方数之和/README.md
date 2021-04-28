# 633.平方数之和
[https://leetcode-cn.com/problems/sum-of-square-numbers](https://leetcode-cn.com/problems/sum-of-square-numbers) 
## 原题
给定一个非负整数 `c` ，你要判断是否存在两个整数 `a` 和 `b` ，使得 `a^2 + b^2 = c` 。

 

 **示例 1：** 

```
输入：c = 5
输出：true
解释：1 * 1 + 2 * 2 = 5

```
 **示例 2：** 

```
输入：c = 3
输出：false

```
 **示例 3：** 

```
输入：c = 4
输出：true

```
 **示例 4：** 

```
输入：c = 2
输出：true

```
 **示例 5：** 

```
输入：c = 1
输出：true
```
 

 **提示：** 
-  `0 <= c <= 2^31 - 1` 
 
**标签**
`数学` 


## 夹逼
```go
func judgeSquareSum(c int) bool {
	a, b := 0, int(math.Sqrt(float64(c)))

	for a <= b {
		t := a*a + b*b
		switch {
		case t == c:
			return true
		case t > c:
			b--
		case t < c:
			a++
		}
	}

	return false
}
```
>执行用时: 0 ms
内存消耗: 1.9 MB



## 遍历
```go
func judgeSquareSum(c int) bool {

	for a := 0; a*a <= c; a++ {
		t := math.Sqrt(float64(c - a*a))
		if t == math.Floor(t) {
			return true
		}
	}

	return false
}
```
>执行用时: 0 ms
内存消耗: 1.9 MB
