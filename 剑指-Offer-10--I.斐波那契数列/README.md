# 剑指 Offer 10- I. 斐波那契数列

[https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/)

## 递归法
```go
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1)%1000000007 + fib(n-2)%1000000007
}
```
本来想说这是水题的，没想到直接`TLE`了。没办法那就递推把。超时的主要原因就是对于我们只想求一个值而言，我们前面的结果是可以复用的。而递归不能做到复用，所以超时了。

## 递推
```go
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b%1000000007, (a+b)%1000000007
	}
	return b
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了75.84%的用户

这里用到了 `golang` 的多返回值的特性，所以内存消耗大了一丢丢。于是改成下面这样

### 优化
```go
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 2; i <= n; i++ {
		c = (a + b) % 1000000007
		a = b
		b = c
	}
	return b
}
```

>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了82.54%的用户

可以看到消耗又小了一丢丢，再优化一下

### 再优化
```go
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a := 0
	b := 1
	c := 0
	for i := 2; i <= n; i++ {
		c = (a + b) % 1000000007
		a = b
		b = c
	}
	return b
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了56.01%的用户

啊这，怎么比之前差远了。看来可能是赋初值的多值返回更省内存？`golang`的底层等我刷完了一遍再来研究吧。