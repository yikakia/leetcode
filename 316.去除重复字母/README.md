# 316. 去除重复字母
[https://leetcode-cn.com/problems/remove-duplicate-letters/](https://leetcode-cn.com/problems/remove-duplicate-letters/) 
## 原题
给你一个字符串 `s` ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 **返回结果的字典序最小** （要求不能打乱其他字符的相对位置）。
**注意：** 该题与 1081 <a href="https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters">https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters</a> 相同
 
**示例 1：** 
```
输入：s = "bcabc"
输出："abc"
```
**示例 2：** 
```
输入：s = "cbacdcbc"
输出："acdb"
```
 
**提示：** 
- `1 <= s.length <= 10^4`
- `s` 由小写英文字母组成


## 单调栈变种
```go
func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}
```
>执行用时: 0 ms
内存消耗: 2.1 MB

和单调栈相差不大，主要是出栈的时候还要判断这个字符在之后有没有，如果有那么就可以出栈，如果没有就不行（题意要求每个都要至少出现一次）。
