# 剑指 Offer 31. 栈的压入、弹出序列
[https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/](https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/)

## 辅助栈模拟出栈
```go
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	index := 0
	for _, value := range pushed {
		stack = append(stack, value)
		for len(stack) != 0 && popped[index] == stack[len(stack)-1] {
			index++
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
```
>执行用时：8 ms, 在所有 Go 提交中击败了85.38%的用户
内存消耗：3.8 MB, 在所有 Go 提交中击败了51.19%的用户

开一个栈来模拟压栈的顺序压栈。然后当栈顶元素和当前应该出栈的元素相等的时候，就一直出栈到栈顶元素和当前应该出栈的元素不等时，就继续压栈，重复循环。

最后当没有完全出栈时就是 `false` 否则就是 `true`