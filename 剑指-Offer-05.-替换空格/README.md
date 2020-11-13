# 剑指 Offer 05. 替换空格
[https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)

## 遍历并把结果存在一个新的字符串中
```go
func replaceSpace(s string) string {
	var r []rune
	for _, char := range s {
		if char == ' ' {
			r = append(r, '%', '2', '0')
		} else {
			r = append(r, char)
		}
	}
	return string(r)
}
```

>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了45.41%的用户

水题吧，没看原书不清楚考点是什么。需要注意的是`Golang`里面对 `string` 遍历的时候是得到的 `rune` 类型的变量。