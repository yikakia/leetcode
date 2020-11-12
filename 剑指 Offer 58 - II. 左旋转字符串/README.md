# 剑指 Offer 58 - II. 左旋转字符串
[https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/](https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/)
## 水题
```go
func reverseLeftWords(s string, n int) string {
	res := strings.Builder{}
	a := s[:n]
	b := s[n:]
	res.WriteString(b)
	res.WriteString(a)
	return res.String()
}

```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：3.3 MB, 在所有 Go 提交中击败了34.25%的用户

