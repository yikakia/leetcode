# 3. 无重复字符的最长子串
[https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/) 
##  滑动窗口
```go
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 1
	res := end - start
	visited := map[byte]int{}
	for i := range s {
		if index, ok := visited[s[i]]; !ok {
			visited[s[i]] = i
			end++
		} else if index < start {
			visited[s[i]] = i
			end++
		} else if index >= start {
			start = index + 1
			visited[s[i]] = i
			end++
		}
		if tmp := end - start; tmp > res {
			res = tmp
		}
	}
	return res - 1
}
```
>执行用时: 12 ms
内存消耗: 3.1 MB

和之前的题[剑指 Offer 48. 最长不含重复字符的子字符串](..\剑指-Offer-48.-最长不含重复字符的子字符串\README.md) 是一样的，简单地说就是用一个 map  来记录这个字符之前出现的位置，同时维护一个首尾指针。首位指针指向的区间就是我们当前所包括的下标。

我们找的时候就是判断当前元素之前出现过没有
1. 出现过，那么再分两种情况判断
    1. 出现过，但是在这个区间之外，也就是说在当前区间是第一次出现，那么就更新这个元素出现的位置在当前位置，尾指针向后挪一位。
    2. 出现过，并且在这个区间之内。由题意要求可知区间内不能有相同的元素，因此就得更新区间的左部到这个元素之前出现的位置的下一个位置。因为之前这个元素出现时所有的可能性都被考虑过了，因此就可以舍弃掉这个元素之前出现的位置的那一部分。
    同时还得更新这个元素的出现位置为当前位置。尾指针向后挪一位。
2. 没出现过
    那么就可以直接更新这个元素的出现位置，并且把尾指针向后挪一位。