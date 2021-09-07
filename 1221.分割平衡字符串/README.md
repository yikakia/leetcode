# 1221.分割平衡字符串
[https://leetcode-cn.com/problems/split-a-string-in-balanced-strings](https://leetcode-cn.com/problems/split-a-string-in-balanced-strings) 
## 原题
在一个 **平衡字符串** 中， `'L'` 和 `'R'` 字符的数量是相同的。

给你一个平衡字符串  `s` ，请你将它分割成尽可能多的平衡字符串。

 **注意：** 分割得到的每个字符串都必须是平衡字符串。

返回可以通过分割得到的平衡字符串的 **最大数量** **。** 

 

 **示例 1：** 

```

输入：s = "RLRRLLRLRL"
输出：4
解释：s 可以分割为 "RL"、"RRLL"、"RL"、"RL" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。

```
 **示例 2：** 

```

输入：s = "RLLLLRRRLR"
输出：3
解释：s 可以分割为 "RL"、"LLLRRR"、"LR" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。

```
 **示例 3：** 

```

输入：s = "LLLLRRRR"
输出：1
解释：s 只能保持原样 "LLLLRRRR".

```
 **示例 4：** 

```

输入：s = "RLRRRLLRLL"
输出：2
解释：s 可以分割为 "RL"、"RRRLLRLL" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。

```
 

 **提示：** 
-  `1 <= s.length <= 1000` 
-  `s[i] = 'L' 或 'R'` 
-  `s` 是一个 **平衡** 字符串
 
**标签**
`贪心` `字符串` `计数` 


## 水题
```go
func balancedStringSplit(s string) (ans int) {
	lC, rC := 0, 0
	for _, ch := range s {
		if ch == 'L' {
			lC++
		}
		if ch == 'R' {
			rC++
		}
		if lC == rC {
			ans++
			lC = 0
			rC = 0
		}
	}
	return
}
```
>执行用时：0 ms
内存消耗：1.9 MB

为了求最大的解决方案，那么就让尽量让每一个不平衡的字符串尽快匹配。如果不尽快匹配的话，那么当前是平衡的，那么就要再找到下一个平衡的地方这样整体才能平衡。但是这样的话原本是两个平衡字符串，现在就变成一个了。