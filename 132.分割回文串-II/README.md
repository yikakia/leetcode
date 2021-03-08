# 132.分割回文串 II
[https://leetcode-cn.com/problems/palindrome-partitioning-ii](https://leetcode-cn.com/problems/palindrome-partitioning-ii) 
## 原题
给你一个字符串 `s` ，请你将 `s` 分割成一些子串，使每个子串都是回文。

返回符合要求的 **最少分割次数** 。
 

 **示例 1：** 

```

输入：s = "aab"
输出：1
解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。

```
 **示例 2：** 

```

输入：s = "a"
输出：0

```
 **示例 3：** 

```

输入：s = "ab"
输出：1

```
 

 **提示：** 
-  `1 <= s.length <= 2000` 
-  `s` 仅由小写英文字母组成
 
**标签**
`动态规划` 


## 
```go

func minCut(s string) int {
	n := len(s)
	isPa := make([][]bool, n)
	for i := range isPa {
		isPa[i] = make([]bool, n)
		for j := range isPa {
			isPa[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			isPa[i][j] = s[i] == s[j] && isPa[i+1][j-1]
		}
	}

	f := make([]int, n)

	for i := range f {
		if isPa[0][i] {
			continue
		}
		f[i] = n
		for j := 0; j < i; j++ {
			if isPa[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}

	return f[n-1]
}
```
>执行用时: 12 ms
内存消耗: 5 MB

和 [131.分割回文串](../131.分割回文串/README.md) 一样，先求出哪些是回文串。不同的是这次是要求出最少分割，而字符串长度最长为 2000 ，这个时候用之前求出所有分割就会超时。

因为要求的是个数，而不是怎么切分，所以就用动态规划记录到当前位置为止，要达到题目要求需要切分多少个。那么下一个回文串就是当前位置再分割一次。

