# 424.替换后的最长重复字符
[https://leetcode-cn.com/problems/longest-repeating-character-replacement/](https://leetcode-cn.com/problems/longest-repeating-character-replacement/) 
## 原题
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 *k* 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

**注意：** 字符串长度 和 *k* 不会超过 10^4。

 

**示例 1：** 

```

输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。

```
**示例 2：** 

```

输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。

```
 
**标签**
`双指针` `Sliding Window` 


## 
```go
func characterReplacement(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range s {
		cnt[ch-'A']++
		maxCnt = max(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
>执行用时: 0 ms
内存消耗: 2.4 MB

开一个滑动窗口，里面保存了26个大写英文字母在窗口中的出现次数。每次向右拓展窗口，更新最大数量。如果最大数量与窗口大小之差大于了可以变更的数量，那么窗口的左边向右移动一位，缩小窗口，这样让这里面最大的数能够符合条件。一直到最右，长度减去窗口左边的位置即为最后的最长边。