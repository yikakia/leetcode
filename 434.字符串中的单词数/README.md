# 434.字符串中的单词数
[https://leetcode-cn.com/problems/number-of-segments-in-a-string](https://leetcode-cn.com/problems/number-of-segments-in-a-string) 
## 原题
统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。

请注意，你可以假定字符串里不包括任何不可打印的字符。

 **示例:** 

```
输入: "Hello, my name is John"
输出: 5
解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。

```
 
**标签**
`字符串` 


## 简单判断
```go
func countSegments(s string) int {
	var flag bool
	var res int
	for i := range s {
		if s[i] == ' ' {
			if flag {
				res++
			}
			flag = false
		} else {
			flag = true
		}
	}
	if flag {
		res++
	}
	return res
}
```
>执行用时：0 ms
内存消耗：1.9 MB

简单的进行判断即可。