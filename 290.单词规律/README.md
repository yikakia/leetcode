# 290. 单词规律
[https://leetcode-cn.com/problems/word-pattern/](https://leetcode-cn.com/problems/word-pattern/) 
## 原题
给定一种规律 `pattern` 和一个字符串 `str` ，判断 `str` 是否遵循相同的规律。
这里的 **遵循** 指完全匹配，例如， `pattern` 里的每个字母和字符串 `str`中的每个非空单词之间存在着双向连接的对应规律。
**示例1:** 
```
输入: pattern = "abba", str = "dog cat cat dog"
输出: true
```
**示例 2:** 
```
输入:pattern = "abba", str = "dog cat cat fish"
输出: false
```
**示例 3:** 
```
输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
```
**示例 4:** 
```
输入: pattern = "abba", str = "dog dog dog dog"
输出: false
```
**说明:** <br>
你可以假设 `pattern` 只包含小写字母， `str` 包含了由单个空格分隔的小写字母。    


## 双哈希表反查
```go
func wordPattern(pattern string, s string) bool {
	ss := strings.Fields(s)
	runePattern := []rune(pattern)
	if len(ss) != len(runePattern) || len(ss) == 0 || len(runePattern) == 0 {
		return false
	}

	matchMap := make(map[rune]string)
	reverseMap := make(map[string]rune)

	for i := range ss {
		if match, ok := matchMap[runePattern[i]]; ok {
			if match != ss[i] || reverseMap[ss[i]] != runePattern[i] {
				return false
			}
		} else {
			if _, ok = reverseMap[ss[i]]; !ok {
				matchMap[runePattern[i]] = ss[i]
				reverseMap[ss[i]] = runePattern[i]
			} else {
				return false
			}
		}
	}
	return true
}
```
>执行用时: 0 ms
内存消耗: 1.9 MB

思路很简单，`map1` 记录模式`a` `(rune)`对应单词`word1` `(string)`,`map2`记录单词对应什么模式。

因为 golang 是采用 `rune` 类型来记录的，所以当用 for 循环取下标的时候可能会出错。不过其实题目的范围在 `ascii` 范围之内，直接用下标索引没问题。

而像我这样提前转化成 `[]rune` 则是个更加通用的做法。不然就得在 `for range `循环之外再加一个变量记录这是第几个 `rune` 。