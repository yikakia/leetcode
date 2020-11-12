# 剑指 Offer 58 - I. 翻转单词顺序
[https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/](https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/)
## 水题
```go
func reverseWords(s string) string {
	res := strings.Builder{}
	n := len(s)
	i := n - 1
	first := true
	for i >= 0 {
		i = trim(s, i)
		if i < 0 {
			break
		}
		j := i
		for j >= 0 && s[j] != ' ' {
			j--
		}
		if !first {
			res.WriteByte(' ')
		}
		res.WriteString(s[j+1 : i+1])
		first = false
		i = j
	}
	return res.String()
}

// 返回一个下标，下标指向 index 之前第一个非空格的位置
func trim(s string, index int) int {
	i := index
	for i >= 0 && s[i] == ' ' {
		i--
	}
	return i
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：3.7 MB, 在所有 Go 提交中击败了55.76%的用户

水题，就是从后往前，找到第一个非' ' 的元素，然后再找到这之前第一个为' '的元素。两个之间的夹的就是一个单词。然后加入到结果就行了。

不过我看一堆人题解用的是标准库，这真的好么？