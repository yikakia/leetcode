# 87.扰乱字符串
[https://leetcode-cn.com/problems/scramble-string](https://leetcode-cn.com/problems/scramble-string) 
## 原题
使用下面描述的算法可以扰乱字符串 `s` 得到字符串 `t` ：

- 如果字符串的长度为 1 ，算法停止
- 如果字符串的长度 > 1 ，执行下述步骤：
	
- 在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 `s` ，则可以将其分成两个子字符串 `x` 和 `y` ，且满足 `s = x + y` 。
-  **随机** 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后， `s` 可能是 `s = x + y` 或者 `s = y + x` 。
- 在 `x` 和 `y` 这两个子字符串上继续从步骤 1 开始递归执行此算法。
	
	
给你两个 **长度相等** 的字符串 `s1` 和  `s2` ，判断  `s2` * * 是否是  `s1` * * 的扰乱字符串。如果是，返回 `true` ；否则，返回 `false` 。

 

 **示例 1：** 

```

输入：s1 = "great", s2 = "rgeat"
输出：true
解释：s1 上可能发生的一种情形是：
"great" --> "gr/eat" // 在一个随机下标处分割得到两个子字符串
"gr/eat" --> "gr/eat" // 随机决定：「保持这两个子字符串的顺序不变」
"gr/eat" --> "g/r / e/at" // 在子字符串上递归执行此算法。两个子字符串分别在随机下标处进行一轮分割
"g/r / e/at" --> "r/g / e/at" // 随机决定：第一组「交换两个子字符串」，第二组「保持这两个子字符串的顺序不变」
"r/g / e/at" --> "r/g / e/ a/t" // 继续递归执行此算法，将 "at" 分割得到 "a/t"
"r/g / e/ a/t" --> "r/g / e/ a/t" // 随机决定：「保持这两个子字符串的顺序不变」
算法终止，结果字符串和 s2 相同，都是 "rgeat"
这是一种能够扰乱 s1 得到 s2 的情形，可以认为 s2 是 s1 的扰乱字符串，返回 true

```
 **示例 2：** 

```

输入：s1 = "abcde", s2 = "caebd"
输出：false

```
 **示例 3：** 

```

输入：s1 = "a", s2 = "a"
输出：true

```
 

 **提示：** 
-  `s1.length == s2.length` 
-  `1 <= s1.length <= 30` 
-  `s1` 和 `s2` 由小写英文字母组成
 
**标签**
`字符串` `动态规划` 


## 动态规划
```go
func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(offset1, offset2, length int) int8
	dfs = func(offset1, offset2, length int) (res int8) {
		d := &dp[offset1][offset2][length]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()
		sub1 := s1[offset1 : offset1+length]
		sub2 := s2[offset2 : offset2+length]

		if sub1 == sub2 {
			return 1
		}

		if !sameFreq(sub1, sub2) {
			return 0
		}

		for i := 1; i < length; i++ {
			if dfs(offset1, offset2, i) == 1 && dfs(offset1+i, offset2+i, length-i) == 1 {
				return 1
			}
			if dfs(offset1, offset2+length-i, i) == 1 && dfs(offset1+i, offset2, length-i) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(0, 0, n) == 1
}

func sameFreq(s1, s2 string) bool {
	freq := map[byte]int{}

	for i := range s1 {
		freq[s1[i]]++
	}
	for i := range s2 {
		freq[s2[i]]--
		if freq[s2[i]] < 0 {
			return false
		}
	}
	return true
}
```
>执行用时: 4 ms
内存消耗: 4.5 MB

整体的思路就是判断 `s1` 的某一段和 `s2` 的某一段是不是相等的，并且记录下来。而之后的操作就是判断另外一半是不是相等的，如果都相等，那么说明当前所遍布的范围是可以通过规定的操作得到的。把当前的结果存入 dp 数组即可。

# 暴力但是 TLE 了
```go
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	n := len(s1)

	if n == 2 {
		return s1[0] == s2[1] && s1[1] == s2[0]
	}

	if !sameFreq(s1, s2) {
		return false
	}

	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
	}
	return false
}

func sameFreq(s1, s2 string) bool {
	freq := map[byte]int{}

	for i := range s1 {
		freq[s1[i]]++
	}
	for i := range s2 {
		freq[s2[i]]--
		if freq[s2[i]] < 0 {
			return false
		}
	}
	return true
}
```
>286 / 288 个通过测试用例

这里是最原始的方法，就是把字符串拆分，然后判断满足哪种情况之后返回即可。