# 剑指 Offer 50. 第一个只出现一次的字符
[https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/](https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/)

## 用 map 映射
```go
func firstUniqChar(s string) byte {
	n := 0
	ocur := make(map[int]byte)
	times := make(map[byte]int)
	for i := range s {
		char := s[i]
		_, exist := times[char]
		if !exist {
			ocur[n] = char
			times[char] = 1
			n++
		} else {
			times[char]++
		}
	}
	for i := 0; i < n; i++ {
		if times[ocur[i]] == 1 {
			return ocur[i]
		}
	}
	return ' '
}


```
>执行用时：56 ms, 在所有 Go 提交中击败了24.51%的用户
内存消耗：5.4 MB, 在所有 Go 提交中击败了64.23%的用户

最开始没注意到题意。题目里面说了只有小写字母。这样的话直接起一个`cap`为`26`的数组就好。然后出现就加一，这样的效率更高。而用 `map` 的话则是更加通用。而`ocur`可以优化也可以不优化。因为这个是记录了每个字符的出现顺序。相对而言这样的时间效率更高（对于用 `map` 而言）