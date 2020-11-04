# 剑指 Offer 48. 最长不含重复字符的子字符串
[https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/](https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/)
## 滑动窗口
```go
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	r := 1
	visited := make(map[byte]int)
	head := 0
	tail := 0
	visited[s[0]] = 0
	for i := 1; i < len(s); i++ {
		preLocal, exist := visited[s[i]]
		// 没出现过
		if !exist {
			visited[s[i]] = i
			tail = i
			// 出现过但是在头指针之前，这个时候说明没在当前的
			// 字符串中出现过，那么尾指针可以向后移动
		} else if preLocal < head {
			visited[s[i]] = i
			tail = i
			// 出现过但是在头指针之后，这个时候说明在当前的字
			// 符串中出现过，那么头指针向后移动一位，跳过。
		} else {
			head = preLocal + 1
			visited[s[i]] = i
			tail = i
			continue
		}
		if lenth := tail - head + 1; lenth > r {
			r = lenth
		}
	}
	return r
}

```

>执行用时：8 ms, 在所有 Go 提交中击败了73.00%的用户
内存消耗：3.1 MB, 在所有 Go 提交中击败了29.80%的用户

简单地说就是通过一个头指针和尾指针来表示当前的字符串的起始和结束位置。同时当遇到相同元素时，尽可能大的维持滑动窗口（跳到这个元素之前出现位置的下一位）。不过也有个要判断的是这个元素之前出现过没有，如果没有那么直接向后滑动就可以了。如果出现了，那么就要判断这个元素是在当前的头指针之前还是之后。在头指针之前说明不在这个字符串中出现过，那么直接更新到现在这个位置；如果在头指针之后，就说明之前在这个字符串中出现过，那么头指针就得直接跳到这个元素之前出现过的位置，同时也要更新这个字符最后出现的位置到现在这个位置。