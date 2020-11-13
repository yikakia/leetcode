# 剑指 Offer 56 - I. 数组中数字出现的次数
[https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/)
## 异或的用法
```go

func singleNumbers(nums []int) []int {
	n := 0
	for _, value := range nums {
		n ^= value
	}
	mask := 1
	for n&mask == 0 {
		mask <<= 1
	}
	a, b := n, n
	for _, value := range nums {
		if mask&value == 0 {
			a ^= value
		} else {
			b ^= value
		}
	}
	return []int{a, b}
}
```
>执行用时：20 ms, 在所有 Go 提交中击败了84.26%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了35.97%的用户

简单地说就是异或是不同为 1 ，相同为 0 。那么对整个数组只有一个数字不成对的话，如果全部异或一遍，得到的就是那一个单独的数字。但是这里我们要得到的是两个单独的数字，这个时候就要想办法让原来的数组分成两个数组，对每个数组分别进行异或。这样就能找到那个值了。

为了分类，为了在区分出两个不同的值的基础上分类，我们就得对先前得到的数`n`分别与`1` `10` `100` …… 相与，直到结果 `mask&n`为非 0 ，这个时候我们得到的值就可以用来区别两个值了。因为若值为 0 的话，就说明这最低的几位还不能用于区别两个不同的数。

之后就可以再遍历一遍数组了，按照与`mask`相与的结果是否为 1 来简单地分类。这里我是让两个数的初始值为`n`，这样就相当于抵消掉了满足这个规则的数。当然初始值也可以为 `0` 因为`0` 与任何数相与都是任何数本身。不过设置其他的值为初始值也是可以的，不过最后输出的时候还得再与这个初始值相与。就结果而言，前后相与的结果本来就为 `0` ，不如初始值就设成 `0` 就行了。