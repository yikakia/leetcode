# 567.字符串的排列
[https://leetcode-cn.com/problems/permutation-in-string/](https://leetcode-cn.com/problems/permutation-in-string/) 
## 原题
给定两个字符串 **s1** 和 **s2**，写一个函数来判断 **s2** 是否包含 **s1**的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

**示例1:**

```

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

```
 

**示例2:**

```

输入: s1= "ab" s2 = "eidboaoo"
输出: False

```
 

**注意：**
- 输入的字符串只包含小写字母
- 两个字符串的长度都在 [1, 10,000] 之间
 
**标签**
`双指针` `Sliding Window` 


## 词频统计
```go
func isSame(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	n := len(s1)
	m := len(s2)

	if n > m {
		return false
	}
	cnt1 := [26]int{}
	cnt2 := [26]int{}

	for i := 0; i < n; i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}

	for i := n; i < m; i++ {
		if isSame(cnt1, cnt2) {
			return true
		}
		cnt2[s2[i-n]-'a']--
		cnt2[s2[i]-'a']++
	}

	return isSame(cnt1, cnt2)
}
```
>执行用时: 0 ms
内存消耗: 2.6 MB

一个字符串的排列就是把它所有字符打乱排序的集合。所以一个字符串*a*是另一个字符串*b*的排列，就是说它们两个的字母频率相同。

因此要求*s1*的排列之一是*s2*的子串的话，只用统计出现字母的个数即可。因为题目限制了是小写字母，因此直接用数组即可。而如果不限定小写字母的话就得用*map*了。

另外其实可以只用一个数组来表示，判断相同的条件是这个数组全为*0*。