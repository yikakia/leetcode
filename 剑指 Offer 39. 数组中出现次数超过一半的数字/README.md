# 剑指 Offer 39. 数组中出现次数超过一半的数字
[https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/](https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/)

## 投票法

```go
func majorityElement(nums []int) int {
	r := 0
	count := 0
	for _, value := range nums {
		if count == 0 {
			r = value
		}
		if r != value {
			count--
		} else {
			count++
		}

	}
	return r
}
```
>执行用时：20 ms, 在所有 Go 提交中击败了87.69%的用户
内存消耗：6.1 MB, 在所有 Go 提交中击败了26.57%的用户


简单地说就是一个个对比，和当前暂时的最大进行对比，相同就计数加一，不同就计数减一。但是这个只能找出超过一半的数，放在一般化的找最多的就应该用`map`来计数了。