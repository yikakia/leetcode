# 剑指 Offer 57 - II. 和为s的连续正数序列
[https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/](https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/)
##
```go
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	for i := int(math.Sqrt(float64(2 * target))); i >= 2; i-- {
		judge := target - i*(i-1)/2
		if judge%i == 0 {
			begin := judge / i
			temp := make([]int, 0)
			for j := 0; j < i; j++ {
				temp = append(temp, begin+j)
			}
			res = append(res, temp)
		}
	}
	return res
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了48.04%的用户

记第一个数为x,一共有k个数在序列里，那么满足公式

$(x+x+k-1)*k/2=target$

我们可以把它化简成：

$target - k(k-1)/2=x*k$

需要注意的是，x 和 k 都是整数，而 k 最大也大不过$\sqrt2(n)$，因为我们可以将其化简成 $k=\sqrt(2 * target+((2 * x-1)/2)^2)-(2 * x-1)/2$

我们将$(2*x-1)/2$看作一个整体的话，那么就是 $k=\sqrt(2 *target+t^2) -t$ 当 target 一定时，t 在$[0,+inf)$上使得 k 趋近于 0 。也就是说当 t = 0 时，k 可以取到最大值 $\sqrt(2 * target)$

而 k 的最小值自然是 2 了。至少由两个数组成嘛。

而当 $(target - k*(k-1)/2) $ 可以被 k 整除时，这个时候就能得到合法的 x 了。因为 x 和 k 都要是整数。