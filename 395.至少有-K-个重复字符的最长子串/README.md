# 395.至少有 K 个重复字符的最长子串
[https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters](https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters) 
## 原题
给你一个字符串 `s` 和一个整数 `k` ，请你找出 `s` 中的最长子串， 要求该子串中的每一字符出现次数都不少于 `k` 。返回这一子串的长度。

 

 **示例 1：** 

```

输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。

```
 **示例 2：** 

```

输入：s = "ababbc", k = 2
输出：5
解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
```
 

 **提示：** 
-  `1 <= s.length <= 10^4` 
-  `s` 仅由小写英文字母组成
-  `1 <= k <= 10^5` 
 
**标签**
`递归` `分治算法` `Sliding Window` 


## 递归调用
```go
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestSubstring(s string, k int) (res int) {
	if s == "" {
		return
	}
	cnt := [26]int{}
	for i := range s {
		cnt[s[i]-'a']++
	}
	var split int
	for i := range cnt[:] {
		if cnt[i] > 0 && cnt[i] < k {
			split = 'a' + i
			break
		}
	}
	if split == 0 {
		return len(s)
	}
	for _, newS := range strings.Split(s, string(byte(split))) {
		res = max(res, longestSubstring(newS, k))
	}
	return
}
```
>执行用时: 0 ms
内存消耗: 2.3 MB

题目要求子串的每个字符出现次数都要大于等于 `k` 那么就相当于说字符串中的字符出现次数小于 `k` 的不会出现在结果里。

同时因为要求的是子串而不是子序列，那么就要切分掉所有出现次数小于 `k` 的字符，即把字符串切分成多个子串，在这些子串中寻找满足条件的子串，并且更新最大值。

题目本身不难，难的是逆向思维。

## 滑动窗口
```go

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestSubstring(s string, k int) (res int) {
	for t := 1; t <= 26; t++ {
		cnt := [26]int{}
		l := 0
		total := 0
		lessK := 0
		for r, ch := range s {
			if cnt[ch-'a'] == 0 {
				total++
				lessK++
			}
			cnt[ch-'a']++
			if cnt[ch-'a'] == k {
				lessK--
			}
			for total > t {
				if cnt[s[l]-'a'] == k {
					lessK++
				}
				cnt[s[l]-'a']--
				if cnt[s[l]-'a'] == 0 {
					total--
					lessK--
				}
				l++
			}
			if lessK <= 0 {
				res = max(res, r-l+1)
			}
		}
	}
	return
}
```
>执行用时: 0 ms
内存消耗: 2 MB

滑动窗口的思路是开一个窗口，维护窗口里面只有 `t` 种字符，看只有 `t` 种字符的时候是否能满足条件，更新结果。`t` 从 `1` 到 `26` 即遍历所有小写字符的情况。 

于是用一个数组来记录窗口中每种字符的数量，这里用 `total` 来记录，在更新的时候看窗口内是否所有字符都满足大于等于 `k` 。这里用一个 `lessK` 来记录。因为在更新的窗口的时候也能顺便判断进出窗口的这个字符的数量是否大于等于 `k` 或者等于 `0` 

感觉滑动窗口还是对窗口的设计最重要。其实实现起来比较简单。