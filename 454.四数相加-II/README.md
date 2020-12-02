# 454. 四数相加 II
[https://leetcode-cn.com/problems/4sum-ii/](https://leetcode-cn.com/problems/4sum-ii/) 
## 哈希表
```go
func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	abPlus := make(map[int]int, 0)
	for _, x := range A {
		for _, y := range B {
			abPlus[x+y]++
		}
	}
	for _, x := range C {
		for _, y := range D {
			res += abPlus[-x-y]
		}
	}
	return res
}
```
>执行用时: 84 ms
内存消耗: 22.7 MB

才发现这个之前没写。那就写一写吧。

简单来说就是通过合并两个数组来形成一个哈希表。如果**A**和**B**相加的结果能被**C**和**D**相加的结果抵消，那么就结果加一。而为什么要拆成平均的两份呢？因为原本是$n^4$的复杂度，拆成两份之后就是$n^2$的时间复杂度了。而我们拆成**A B C**一组，**D**一组可以么？也可以，但是这样就是$n^3$的时间复杂度了，明显不如平分的时间复杂度。