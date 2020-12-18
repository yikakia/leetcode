# 389. 找不同
[https://leetcode-cn.com/problems/find-the-difference/](https://leetcode-cn.com/problems/find-the-difference/) 
## 原题
给定两个字符串 ***s** * 和 ***t** *，它们只包含小写字母。
字符串 ***t***  由字符串 ***s***  随机重排，然后在随机位置添加一个字母。
请找出在 ***t** * 中被添加的字母。
 
**示例 1：** 
```
输入：s = "abcd", t = "abcde"
输出："e"
解释：&#39;e&#39; 是那个被添加的字母。
```
**示例 2：** 
```
输入：s = "", t = "y"
输出："y"
```
**示例 3：** 
```
输入：s = "a", t = "aa"
输出："a"
```
**示例 4：** 
```
输入：s = "ae", t = "aea"
输出："a"
```
 
**提示：** 
- `0 <= s.length <= 1000`
- `t.length == s.length + 1`
- `s` 和 `t` 只包含小写字母


## 
```go
func findTheDifference(s string, t string) byte {
	count := [26]int{}
	for _, char := range s {
		count[char-'a']++
	}
	for _, char := range t {
		count[char-'a']--
	}
	for index, v := range count {
		if v != 0 {
			return byte('a' + index)
		}
	}

	return 0
}
```
>执行用时: 0 ms
内存消耗: 2.1 MB

就很简单地记录每个字符出现的次数就行了。因为题目要求是只有小写字母，而函数签名的返回值也是 **byte**，所以就不用**rune**来判断了。
