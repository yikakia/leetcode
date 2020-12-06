# 5617. 设计 Goal 解析器
[https://leetcode-cn.com/problems/goal-parser-interpretation/](https://leetcode-cn.com/problems/goal-parser-interpretation/) 
## 模式匹配
```go
func interpret(command string) string {
	res := strings.Builder{}
	needAl := false
	for _, char := range command {
		if char == 'G' {
			res.WriteRune(char)
		} else if char == '(' {
			continue
		} else {
			if char == 'a' || char == 'l' {
				needAl = true
			} else {
				if needAl {
					needAl = false
					res.WriteString("al")
				} else {
					res.WriteRune('o')
				}
			}
		}
	}
	return res.String()
}
```
>执行用时: 0 ms
内存消耗: 2 MB

就相当于简单的模式匹配。

