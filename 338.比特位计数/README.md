# 338.比特位计数
[https://leetcode-cn.com/problems/counting-bits](https://leetcode-cn.com/problems/counting-bits) 
## 原题
给定一个非负整数 **num** 。对于 **0 ≤ i ≤ num** 范围中的每个数字 **i** ，计算其二进制数中的 1 的数目并将它们作为数组返回。

 **示例 1:** 

```
输入: 2
输出: [0,1,1]
```
 **示例 2:** 

```
输入: 5
输出: [0,1,1,2,1,2]
```
 **进阶:** 
- 给出时间复杂度为 **O(n*sizeof(integer))** 的解答非常容易。但你可以在线性时间 **O(n)** 内用一趟扫描做到吗？
- 要求算法的空间复杂度为 **O(n)** 。
- 你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 **__builtin_popcount** ）来执行此操作。
 
**标签**
`位运算` `动态规划` 


## 利用最高有效位
```go
func countBits(num int) (res []int) {
	res = make([]int, num+1)
	for i, highBits := 1, 0; i <= num; i++ {
		// i 为 2 的整数次幂
		if i&(i-1) == 0 {
			highBits = i
		}
		res[i] = res[i-highBits] + 1
	}
	return
}
```
>执行用时: 4 ms
内存消耗: 4.6 MB

具体来说就是把 `x` 转为二进制数的话，可以知道最左边的第一个 `1` 的位置。那么 `x` 减去第一个 `1`  ，它的 `1` 的个数就是之前那个数减一。

## 利用最低位
```go
func countBits(num int) (res []int) {
	res = make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return
}
```

我们上面做的其实就是去除最高位。这个时候用 `x&(x-1)` 可以去除掉最低位的 `1` 。这个时候 `1` 的个数也比 `x` 少一个。

## 总结
就是利用之前的结果，不过要定一个规则，看是比哪个数多一个 `1` 。