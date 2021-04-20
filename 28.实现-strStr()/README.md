# 28.实现 strStr()
[https://leetcode-cn.com/problems/implement-strstr](https://leetcode-cn.com/problems/implement-strstr) 
## 原题
实现 <a href="https://baike.baidu.com/item/strstr/811469" target="_blank">strStr()</a> 函数。

给你两个字符串  `haystack` 和 `needle` ，请你在 `haystack` 字符串中找出 `needle` 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  `-1` **** 。

 

 **说明：** 

当  `needle`  是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当  `needle`  是空字符串时我们应当返回 0 。这与 C 语言的 <a href="https://baike.baidu.com/item/strstr/811469" target="_blank">strstr()</a> 以及 Java 的 <a href="https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)" target="_blank">indexOf()</a> 定义相符。

 

 **示例 1：** 

```

输入：haystack = "hello", needle = "ll"
输出：2

```
 **示例 2：** 

```

输入：haystack = "aaaaa", needle = "bba"
输出：-1

```
 **示例 3：** 

```

输入：haystack = "", needle = ""
输出：0

```
 

 **提示：** 
-  `0 <= haystack.length, needle.length <= 5 * 10^4` 
-  `haystack` 和 `needle` 仅由小写英文字符组成
 
**标签**
`双指针` `字符串` 


## 简单的字符串匹配
```go
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for i := range haystack[:len(haystack)-len(needle)+1] {
		if haystack[i] != needle[0] {
			continue
		}
		exist := true
		for j := range needle {
			if haystack[i+j] != needle[j] {
				exist = false
				break
			}
		}
		if exist {
			return i
		}
	}

	return -1
}
```
>执行用时: 0 ms
内存消耗: 2.2 MB

就是用两个指针分别指向 `haystack` 和 `needle`，然后分别进行匹配即可。

不过这个只是朴素的查找算法，说到字符串匹配不得不提到的就是 `KMP` 算法了。

## KMP 算法

简单地说就是虽然匹配失败了，但是其中有可以利用的信息让我们可以略过部分循环。

比如在 "ABC ABCDAB ABCDABCDABDE" 里面找 "ABCDABD" 

当找到中间的 `ABCDAB` 的时候，我们可以发现 `AB` 是要找的字符串的开头，这个时候直接把字符串的前面两个记录为匹配即可。这样就略过了中间的 `BCD` 的循环。在要找的子串里面有大量重复的字符的时候对效率的提升最大。