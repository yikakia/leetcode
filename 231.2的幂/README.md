# 231.2的幂
[https://leetcode-cn.com/problems/power-of-two](https://leetcode-cn.com/problems/power-of-two) 
## 原题
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

 **示例 1:** 

```
输入: 1
输出: true
解释: 2^0 = 1
```
 **示例 2:** 

```
输入: 16
输出: true
解释: 2^4 = 16
```
 **示例 3:** 

```
输入: 218
输出: false
```
 
**标签**
`位运算` `数学` 


## 
```go
func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0 && n != 0 && n != math.MinInt64
}
```
>执行用时: 0 ms
内存消耗: 2.1 MB

原理很简单，就是用 n&(n-1) 这个运算的结果。如果 n 不是 2 的幂次的话就不会为0。只有当 n 为 2 的幂次的时候才会为 0 。同时还要注意 n 为 0 的时候要额外判断，因为此时 n&(n-1) 依旧为 0。

而 n 为负数的时候由于 n 是通过补码保存的，因此只会在 n = math.MinInt64 的时候出现 n&(n-1) 为 0 的情况。

综上，排除掉 n = 0 和 n = math.MinInt64 的情况即可。

不过还有个更简单的写法
```go
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
```
排除掉为负数的情况即可。