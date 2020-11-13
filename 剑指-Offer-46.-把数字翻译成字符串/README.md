# 剑指 Offer 46. 把数字翻译成字符串
[https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/](https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/)

## 
```go
func translateNum(num int) int {
	snum := strconv.Itoa(num)
	// p 表示 当前元素的前前的元素的表示方法
	// q 表示 当前元素的前元素的表示方法
	p, q, r := 0, 0, 1
	for index := range snum {
		p, q = q, r

		if index == 0 {
			continue
		}
		doubleChar := snum[index-1 : index+1]
		if doubleChar <= "25" && doubleChar >= "10" {
			r += p
		}
	}
	return r
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了21.27%的用户


和青蛙跳台阶一样，不过有些台阶不能跳两格。
但是这个内存消耗大不知道到底是为什么。明明我看`dfs`对内存的消耗应该更大吧。