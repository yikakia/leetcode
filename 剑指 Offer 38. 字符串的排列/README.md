# 剑指 Offer 38. 字符串的排列
[https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/)
## 深搜，对于每一位去掉同样的字符
```go
func permutation(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	r := make([]string, 0)
	hasApper := make(map[rune]bool, 0)
	for index, char := range s {
		if hasApper[char] {
			continue
		}
		tmpr := permutation(s[:index] + s[index+1:])
		for i, word := range tmpr {
			tmpr[i] = string(char) + word
		}
		r = append(r, tmpr...)
		hasApper[char] = true
	}
	return r
}
```
>执行用时：80 ms, 在所有 Go 提交中击败了30.55%的用户
内存消耗：8.2 MB, 在所有 Go 提交中击败了19.66%的用户

简单地说就是不断地把 `s` 拆分，把拆分出的 `char` 作为头添加在没有 `char` 的 `s` 的全排列里面。因为要去重，所以添加一个`map[rune]bool`用来判断这个拆分出的`char`在之前有没有出现过。