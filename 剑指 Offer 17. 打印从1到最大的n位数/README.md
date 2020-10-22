# 剑指 Offer 17. 打印从1到最大的n位数
[https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/](https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/)

```go
func printNumbers(n int) []int {
	tmp := n
	n = 0
	for tmp != 0 {
		n *= 10
		n += 9
		tmp--
	}
	r := make([]int, n)
	for i := 1; i <= n; i++ {
		r[i-1] = i
	}
	return r
}
```
>执行用时：12 ms, 在所有 Go 提交中击败了52.17%的用户
内存消耗：7 MB, 在所有 Go 提交中击败了51.18%的用户

看了下这个内存和耗时，感觉主要是这个计算最大值的时候的方法不对？我这是直接算`99999……`了，而不是算`1000……`最后减`1`。最低内存和最短耗时都是用的内置的阶乘。这个可能会更快？