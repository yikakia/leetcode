# 1047.删除字符串中的所有相邻重复项
[https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string) 
## 原题
给出由小写字母组成的字符串 `S` ， **重复项删除操作** 会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

 

 **示例：** 

```
输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。

```
 

 **提示：** 
-  `1 <= S.length <= 20000` 
-  `S` 仅由小写英文字母组成。
 
**标签**
`栈` 


## 
```go
func removeDuplicates(S string) string {
	stack := make([]byte, 0, len(S)/2)
	for i := range S {
		var changed bool
		for len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
			changed = true
		}
		if !changed {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
```
>执行用时: 4 ms
内存消耗: 6.4 MB

有点类似于单调栈的感觉。
