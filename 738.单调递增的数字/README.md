# 738. 单调递增的数字
[https://leetcode-cn.com/problems/monotone-increasing-digits/](https://leetcode-cn.com/problems/monotone-increasing-digits/) 
## 原题
给定一个非负整数 `N`，找出小于或等于 `N` 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
（当且仅当每个相邻位数上的数字 `x` 和 `y` 满足 `x <= y` 时，我们称这个整数是单调递增的。）
**示例 1:** 
```
输入: N = 10
输出: 9
```
**示例 2:** 
```
输入: N = 1234
输出: 1234
```
**示例 3:** 
```
输入: N = 332
输出: 299
```
**说明:**  `N` 是在 `[0, 10^9]` 范围内的一个整数。


## 
```go
func monotoneIncreasingDigits(N int) int {
	strN := strconv.Itoa(N)
	runeN := []rune(strN)
	i := 0
	for i < len(runeN)-1 {
		if runeN[i] > runeN[i+1] {
			break
		}
		i++
	}
	if i < len(runeN)-1 && runeN[i] > runeN[i+1] {
		runeN[i]--
		for ii := i + 1; ii < len(runeN); ii++ {
			runeN[ii] = '9'
		}
	}
	strN = string(runeN)
	intN, _ := strconv.Atoi(strN)
	if intN == N {
		return intN
	}
	tmp := monotoneIncreasingDigits(intN)
	if tmp == intN {
		return intN
	}
	return monotoneIncreasingDigits(tmp)
}
```
>执行用时: 0 ms
内存消耗: 2 MB

简单地说就是找到第一个破坏递增关系的位置`rune[i]<rune[i+1]`，然后把它`rune[i]`本身减 1 ，后面的全部置 9 。不过这个对于 `N == 33332` 的时候还不够，这样的操作结果为 `N = 33329` ，明显不符合要求。这个时候需要不断地往前继续修改，也就是递归调用了。

什么时候递归结束呢？ 就是转换后的结果和转换前的结果相同就可以结束递归了。