# 66.加一
[https://leetcode-cn.com/problems/plus-one/](https://leetcode-cn.com/problems/plus-one/) 
## 原题
给定一个由 **整数** 组成的 **非空**  数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储**单个** 数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

 

**示例 1：** 

```

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。

```
**示例 2：** 

```

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。

```
**示例 3：** 

```

输入：digits = [0]
输出：[1]

```
 

**提示：** 
- `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`
 
**标签**
`数组` 


## 模拟
```go
func plusOne(digits []int) []int {
	n := len(digits)
	overflow := 0
	digits[n-1]++
	for i := n - 1; i >= 0; i-- {
		digits[i] += overflow
		if digits[i] > 9 {
			overflow = 1
			digits[i] %= 10
		} else {
			overflow = 0
			break
		}
	}
	if overflow > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
```
>执行用时: 0 ms
内存消耗: 2 MB

没啥说的，就是模拟十进制加法。