# 65.有效数字
[https://leetcode-cn.com/problems/valid-number](https://leetcode-cn.com/problems/valid-number) 
## 原题

**有效数字** （按顺序）可以分成以下几个部分：
- 一个 **小数** 或者 **整数** 
- （可选）一个 `'e'` 或 `'E'` ，后面跟着一个 **整数** 
 
**小数** （按顺序）可以分成以下几个部分：
- （可选）一个符号字符（ `'+'` 或 `'-'` ）
- 下述格式之一：
	
- 至少一位数字，后面跟着一个点 `'.'` 
- 至少一位数字，后面跟着一个点 `'.'` ，后面再跟着至少一位数字
- 一个点 `'.'` ，后面跟着至少一位数字
	
	
 **整数** （按顺序）可以分成以下几个部分：
- （可选）一个符号字符（ `'+'` 或 `'-'` ）
- 至少一位数字

部分有效数字列举如下：
-  `["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]` 

部分无效数字列举如下：
-  `["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]` 
给你一个字符串 `s` ，如果 `s` 是一个 **有效数字** ，请返回 `true` 。

 

 **示例 1：** 

```

输入：s = "0"
输出：true

```
 **示例 2：** 

```

输入：s = "e"
输出：false

```
 **示例 3：** 

```

输入：s = "."
输出：false

```
 **示例 4：** 

```

输入：s = ".1"
输出：true

```
 

 **提示：** 
-  `1 <= s.length <= 20` 
-  `s` 仅含英文字母（大写和小写），数字（ `0-9` ），加号 `'+'` ，减号 `'-'` ，或者点 `'.'` 。
 
**标签**
`数学` `字符串` 


## 
```go
func isNum(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	if len(s) == 0 {
		return false
	}
	return isNum(s)
}

func isDecimal(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	index := strings.Index(s, ".")
	if index == -1 {
		return false
	}
	if index == 0 {
		return isNum(s[1:])
	}
	if index == len(s)-1 {
		return isInt(s[:index])
	}

	return isInt(s[:index]) && isNum(s[index+1:])

}

func isNumber(s string) bool {
	s = strings.ReplaceAll(s, "E", "e")
	index := strings.Index(s, "e")
	if index != -1 {
		return (isDecimal(s[:index]) || isInt(s[:index])) &&
			isInt(s[index+1:])
	}

	return isInt(s) || isDecimal(s)
}
```
>执行用时: 0 ms
内存消耗: 2.3 MB

就是简单的模拟而已。另外写几个函数，验证是否是数字、整数、小数即可。