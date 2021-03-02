# 1012.至少有 1 位重复的数字
[https://leetcode-cn.com/problems/numbers-with-repeated-digits](https://leetcode-cn.com/problems/numbers-with-repeated-digits) 
## 原题
给定正整数 `N` ，返回小于等于 `N` 且具有至少 1 位重复数字的正整数的个数。

 

 **示例 1：** 

```
输入：20
输出：1
解释：具有至少 1 位重复数字的正数（<= 20）只有 11 。

```
 **示例 2：** 

```
输入：100
输出：10
解释：具有至少 1 位重复数字的正数（<= 100）有 11，22，33，44，55，66，77，88，99 和 100 。

```
 **示例 3：** 

```
输入：1000
输出：262

```
 

 **提示：** 
-  `1 <= N <= 10^9` 
 
**标签**
`数学` `动态规划` 


## 求出所有不重复的数字
```go
func numDupDigitsAtMostN(N int) int {
	digits := []int{}
	for tn := N; tn != 0; tn /= 10 {
		digits = append(digits, tn%10)
	}
	n := len(digits)
	total := 0

	for i := 1; i < n; i++ {
		total += 9 * A(9, i-1)
	}

	used := [10]int{}
	for i := n - 1; i >= 0; i-- {
		num := digits[i]
		j := 0
		if i == n-1 {
			j = 1
		}

		for ; j < num; j++ {
			if used[j] != 0 {
				continue
			}
			total += A(10-(n-i), i)
		}
		used[num]++
		if used[num] > 1 {
			break
		}
		if i == 0 {
			total++
		}
	}

	return N - total
}

func A(n, m int) int {
	return fact(n) / fact(n-m)
}

var f map[int]int

func fact(n int) int {
	if f == nil {
		f = map[int]int{}
	}
	if v, ok := f[n]; ok {
		return v
	}
	if n < 2 {
		f[n] = 1
		return 1
	}
	f[n] = n * fact(n-1)
	return f[n]
}
```
>执行用时: 0 ms
内存消耗: 1.9 MB

简单地思想就是求出所有不重复的数，然后减去即可。