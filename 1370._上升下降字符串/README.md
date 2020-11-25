# 1370. 上升下降字符串
[https://leetcode-cn.com/problems/increasing-decreasing-string/](https://leetcode-cn.com/problems/increasing-decreasing-string/) 
## 
```go
// 1370.上升下降字符串
// https://leetcode-cn.com/problems/increasing-decreasing-string/
func sortString(s string) string {
	res := strings.Builder{}
	lenOfS := len(s)
	alphabat := make(map[int]int, 0)
	for _, char := range s {
		alphabat[int(char-'a')]++
	}
	for lenOfS > 0 {
		for i := 0; i < 26; i++ {
			if nums, ok := alphabat[i]; ok && nums > 0 {
				alphabat[i]--
				lenOfS--
				res.WriteRune(rune('a' + i))
			}
		}
		for i := 25; i >= 0; i-- {
			if nums, ok := alphabat[i]; ok && nums > 0 {
				alphabat[i]--
				lenOfS--
				res.WriteRune(rune('a' + i))
			}
		}
	}
	return res.String()
}
```
>执行用时: 8 ms
内存消耗: 3.4 MB

直接用`map`做了，不过都这样了，不如直接用数组做？

```go
// 1370.上升下降字符串
// https://leetcode-cn.com/problems/increasing-decreasing-string/
func sortString(s string) string {
	res := strings.Builder{}
	lenOfS := len(s)
	alphabat := make([]int, 26)
	for _, char := range s {
		alphabat[int(char-'a')]++
	}
	for lenOfS > 0 {
		for i := 0; i < 26; i++ {
			if nums := alphabat[i]; nums > 0 {
				alphabat[i]--
				lenOfS--
				res.WriteRune(rune('a' + i))
			}
		}
		for i := 25; i >= 0; i-- {
			if nums := alphabat[i]; nums > 0 {
				alphabat[i]--
				lenOfS--
				res.WriteRune(rune('a' + i))
			}
		}
	}
	return res.String()
}
```
>执行用时: 4 ms
内存消耗: 3.1 MB

时间直接砍了一半。看来就算是$O(1)$的时间复杂度也是比不了直接计算偏移量的速度啊。