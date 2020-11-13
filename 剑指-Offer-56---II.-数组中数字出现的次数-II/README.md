# 剑指 Offer 56 - II. 数组中数字出现的次数 II
[https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/)
## 位操作，真值表
```go
func singleNumber(nums []int) int {
	two, one := 0, 0
	for _, value := range nums {
		one = one ^ value&^two
		two = two ^ value&^one
	}
	return one
}
```
>执行用时：24 ms, 在所有 Go 提交中击败了98.65%的用户
内存消耗：6.4 MB, 在所有 Go 提交中击败了49.55%的用户

建议参考这个题解 [面试题56 - II. 数组中数字出现的次数 II（位运算 + 有限状态自动机，清晰图解）](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/solution/mian-shi-ti-56-ii-shu-zu-zhong-shu-zi-chu-xian-d-4/)

这个写的是真的清楚，简单地说就是从位运算的角度来看，把每一位都分别累加起来，因为除了一个数字出现了一次之外，其他的数字都出现了三次，所以如果对三取余的话我们就能得到这个数字。

那么简单的思路就是开一个`32`位的数组，然后把每一个数都拆分成一位一位的累加进对应的位里面，最后再对三取余，得到的就是这个单独的数字的每一位。最后再不断地左移并且与当前位相与就可以得到这个结果。

一般人能做到这里就不错了，但是还有优化的方法，那就是写出真值表。

我们用两个位来表示我们暂时存的值是`0`,`1`,`2`中的哪一个，然后根据现在遇到的`n`来得到之后的值是`0`,`1`,`2`中的哪一个。这样就可以列出一个真值表。然后我们可以根据真值表写出方程，这样就得到了求单独的数的方法。

然后因为对于每一位的操作都是一样的，所以可以用一个`one int32`和一个`two int32`来直接对每一位进行计算。


真的是牛逼疯了，虽然我也学过数电，但没想到可以这么用。这才是真正的学以致用，这样的人才是真正的计算机学科的人才吧。
