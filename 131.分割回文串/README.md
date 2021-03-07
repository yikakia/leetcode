# 131.分割回文串
[https://leetcode-cn.com/problems/palindrome-partitioning](https://leetcode-cn.com/problems/palindrome-partitioning) 
## 原题
给定一个字符串 *s* ，将 *s* 分割成一些子串，使每个子串都是回文串。

返回 *s* 所有可能的分割方案。

 **示例:** 

```
输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
```
 
**标签**
`深度优先搜索` `动态规划` `回溯算法` 


## dfs
```go
func isPa(s string) bool {
	n := len(s)
	if n < 2 {
		return true
	}

	half := n / 2
	for i := 0; i < half; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}

	return true

}

func partition(s string) [][]string {

	n := len(s)
	if n == 0 {
		return [][]string{{}}
	} else if n == 1 {
		return [][]string{{s}}
	}

	res := [][]string{}
	for i := 0; i < n; i++ {
		if isPa(s[:i+1]) {
			cur := s[:i+1]
			next := partition(s[i+1:])
			for _, row := range next {
				res = append(res, append([]string{cur}, row...))
			}
		}
	}
	return res
}
```
>执行用时: 452 ms
内存消耗: 22.7 MB

就是 dfs 。当现在遍历到的  `s[:i+1]` 是回文串的时候，那么它后面的部分也可能被分割为回文串。因此就简化成了当找到了新的回文串的时候，再找后面的回文串组合。

要注意的是当 s 本身就是一个回文串的话，那么它应该和 [][]string{{}} 相组合。也就是 len(s) == 0 的时候，就说明应该返回含有一个空的子串的组合。

不过时间复杂度高了，还是要简化复杂度。看题解才发现能够简化判断回文串的方式。

## 简化回文串的判断

上面我的做法是把回文串当作了不相干的数据。实际上对于之前判断过的数据可以进行利用。

比如对于子串 `s[i:j]` 如果它是回文串，那么 `s[i-1:j+1]` 在 `s[i-1]==s[j]` 且 `s[i:j]` 为回文串的时候也为回文串。

于是用一个 map 来储存
### map 储存
```go
var paMap map[string]bool

func isPa(s string) bool {
	if paMap == nil {
		paMap = map[string]bool{}
	}
	if t, ok := paMap[s]; ok {
		return t
	}
	n := len(s)
	if n < 2 {
		return true
	}
	t := s[0] == s[n-1] && isPa(s[1:n-1])
	paMap[s] = t
	return t
}

func partition(s string) [][]string {

	n := len(s)
	if n == 0 {
		return [][]string{{}}
	} else if n == 1 {
		return [][]string{{s}}
	}

	res := [][]string{}
	for i := 0; i < n; i++ {
		if isPa(s[:i+1]) {
			cur := s[:i+1]
			next := partition(s[i+1:])
			for _, row := range next {
				res = append(res, append([]string{cur}, row...))
			}
		}
	}
	return res
}
```
>执行用时: 452 ms
内存消耗: 23.9 MB

和之前没有区别啊。应该是优化太少了？用题解的方法来优化下试试

### 官方题解
```go
func partition(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	res := [][]string{}
	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return res
}
```
>执行用时: 300 ms
内存消耗: 23.6 MB

30% 左右的差距。思路是先整理出所有排列的是否是回文串的可能性，然后用 dfs 逐步搜索可能的回文子串，然后匹配之后的结果。

后面的子串搜索完之后就把当前找到的这个子串剔除，继续寻找下一个符合的子串。

### 再优化
```go
func partition(s string) (ans [][]string) {
	n := len(s)
	f := make([][]int8, n)
	for i := range f {
		f[i] = make([]int8, n)
	}

	// 0 表示尚未搜索，1 表示是回文串，-1 表示不是回文串
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if f[i][j] != 0 {
			return f[i][j]
		}
		f[i][j] = -1
		if s[i] == s[j] {
			f[i][j] = isPalindrome(i+1, j-1)
		}
		return f[i][j]
	}

	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(i, j) > 0 {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}

```
>执行用时: 412 ms
内存消耗: 26 MB

比起之前的全部预处理慢了 30% 

这里的思路是不是所有子串都需要搜索，只搜索需要处理的那一部分就可以了。