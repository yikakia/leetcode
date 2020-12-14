# 5626. 十-二进制数的最少数目
[https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/](https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/) 
## 原题
如果一个十进制数字不含任何前导零，且每一位上的数字不是 `0` 就是 `1` ，那么该数字就是一个 **十-二进制数**  。例如，`101` 和 `1100` 都是 **十-二进制数** ，而 `112` 和 `3001` 不是。
给你一个表示十进制整数的字符串 `n` ，返回和为 `n` 的 **十-二进制数 ** 的最少数目。
 
**示例 1：** 
```
输入：n = "32"
输出：3
解释：10 + 11 + 11 = 32
```
**示例 2：** 
```
输入：n = "82734"
输出：8
```
**示例 3：** 
```
输入：n = "27346209830709182346"
输出：9
```
 
**提示：** 
- `1 <= n.length <= 10^5`
- `n` 仅由数字组成
- `n` 不含任何前导零并总是表示正整数


## 模拟
```go
func minPartitions(n string) int {
	res := 0
	nums := []byte(n)
	for i := range nums {
		nums[i] -= '0'
	}
	hasChanged := true
	for hasChanged {
		hasChanged = false
		for i := range nums {
			if nums[i] > 0 {
				hasChanged = true
				nums[i]--
			}
		}
		if hasChanged {
			res++
		}
	}

	return res

}
```
>执行用时: 32 ms
内存消耗: 6.2 MB

暴力模拟拆数

## 找最大的数位
```go
func minPartitions(n string) int {
	nums := []byte(n)
	for i := range nums {
		nums[i] -= '0'
	}
	res := nums[0]
	for _, num := range nums {
		if num == 9 {
			return 9
		}
		if res < num {
			res = num
		}
	}
	return int(res)
}
```
>执行用时: 8 ms
内存消耗: 6.2 MB

从**方法一**可以看出，本质就是找最大的数位。
