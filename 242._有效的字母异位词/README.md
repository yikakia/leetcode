# 242. 有效的字母异位词
[https://leetcode-cn.com/problems/valid-anagram/](https://leetcode-cn.com/problems/valid-anagram/) 
## 伪鸽巢排序
```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countS := make(map[rune]int, 0)
	for _, v := range s {
		countS[v]++
	}
	for _, v := range t {
		if countS[v] == 0 {
			return false
		}
		countS[v]--
	}
	return true
}
```
>执行用时: 8 ms
内存消耗: 2.8 MB

什么是字母异位词？简单地说就是把原字符串排序之后生成的序列相同的字符串就是字母异位词。因此对于解法而言就有两种思路了，一个是排序后对比，一个是相当于鸽巢排序 s ，然后遍历 t 挨个减去，如果有数量不同的情况就返回 false ，只不过用了 map 来简单地实现。

这里我用的是第二种方法，不过虽然这题简单，还是有需要证明的地方，比如长度相等，然后这么相减会不会出现遗漏的情况？比如字符 a 在 s 里面出现了两次，在 t 里面出现了一次，最后这个地方会出现 countS['a']>0 的结果。虽然是这样没错，但是字符 a 少出现了一次的结果是字符 b 在 t 里面多出现了一次，于是会在多出现的那个地方出现 countS['b']==0 的情况，这个时候就已经不满足条件退出循环并返回 false 了，也就无所谓 countS['a'] > 0 了。