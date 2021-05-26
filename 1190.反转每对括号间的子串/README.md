# 1190.反转每对括号间的子串
[https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses](https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses) 
## 原题
给出一个字符串 `s` （仅含有小写英文字母和括号）。

请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。

注意，您的结果中 **不应** 包含任何括号。

 

 **示例 1：** 

```
输入：s = "(abcd)"
输出："dcba"

```
 **示例 2：** 

```
输入：s = "(u(love)i)"
输出："iloveu"

```
 **示例 3：** 

```
输入：s = "(ed(et(oc))el)"
输出："leetcode"

```
 **示例 4：** 

```
输入：s = "a(bcdefghijkl(mno)p)q"
输出："apmnolkjihgfedcbq"

```
 

 **提示：** 
-  `0 <= s.length <= 2000` 
-  `s` 中只有小写英文字母和括号
- 我们确保所有括号都是成对出现的
 
**标签**
`栈` 


## 栈
```go
func reverseParentheses(s string) string {
	q := []string{""}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			q = append(q, "")
		} else if s[i] == ')' {
			str := reverseStr(q[len(q)-1])
			q = q[:len(q)-1]
			q[len(q)-1] += str
		} else {
			q[len(q)-1] = q[len(q)-1] + string(s[i])
		}
	}

	return q[0]
}

func reverseStr(s string) string {
	bStr := []byte(s)
	for i := 0; i < len(bStr)/2; i++ {
		bStr[i], bStr[len(bStr)-i-1] = bStr[len(bStr)-i-1], bStr[i]
	}
	return string(bStr)
}
```
>执行用时: 0 ms
内存消耗: 2.2 MB

简单来说就是通过栈来暂存需要处理的当前部分的数据。
- 遇到了`'('` 说明是新的一层，于是压`""`进去，等待处理。
- 遇到了字母  说明是这一层的数据，直接压到栈顶的尾部即可
- 遇到了`')'` 说明这一层结束了，需要逆序然后出栈返回到上一层的尾部。

按照这个流程处理完之后，栈底就是对应的结果。

需要注意的是，如果字符串的开头就是`(`，那么直接用这个流程的话会出现栈底也被弹出的问题，这个时候初始化的时候先压入一个`""`进去即可。