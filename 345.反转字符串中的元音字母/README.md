# 345.反转字符串中的元音字母
[https://leetcode-cn.com/problems/reverse-vowels-of-a-string](https://leetcode-cn.com/problems/reverse-vowels-of-a-string) 
## 原题
给你一个字符串 `s` ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 `'a'` 、 `'e'` 、 `'i'` 、 `'o'` 、 `'u'` ，且可能以大小写两种形式出现。

 

 **示例 1：** 

```

输入：s = "hello"
输出："holle"

```
 **示例 2：** 

```

输入：s = "leetcode"
输出："leotcede"
```
 

 **提示：** 
-  `1 <= s.length <= 3 * 10^5` 
-  `s` 由 **可打印的 ASCII** 字符组成
 
**标签**
`双指针` `字符串` 


## 双指针
```go
func reverseVowels(s string) string {
	t := []byte(s)
	n := len(t)
	i, j := 0, n-1
	for i < j {
		for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
			i++
		}
		for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}
	}
	return string(t)
}
```
>执行用时：0 ms
内存消耗：3.9 MB

双指针，前后找然后交换找到的。