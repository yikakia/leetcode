# 5. 最长回文子串
[https://leetcode-cn.com/problems/longest-palindromic-substring/](https://leetcode-cn.com/problems/longest-palindromic-substring/) 
## 从中心向两边拓展
```go
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	res := s[:1]
	tmpRes := ""
	for i := 1; i < len(s); i++ {
		width := 1
		l, r := i-width, i+width
		for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			tmpRes = s[l : r+1]
		}
		if len(tmpRes) > len(res) {
			res = tmpRes
		}
		tmpRes = ""
		l, r = i-width, i
		for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			tmpRes = s[l : r+1]
		}
		if len(tmpRes) > len(res) {
			res = tmpRes
		}
		tmpRes = ""
	}
	return res
}
```
>执行用时: 12 ms
内存消耗: 2.6 MB

就是很简单地不断判断是不是回文而已，结果还是官方的解法之一。
