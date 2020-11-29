# 10. 正则表达式匹配
[https://leetcode-cn.com/problems/regular-expression-matching/](https://leetcode-cn.com/problems/regular-expression-matching/) 
## 
```go

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	// 字符串没用完，模式用完了
	if n > 0 && m == 0 {
		return false
	}
	if n == 0 {
		// 字符串和模式都用完了
		if m == 0 {
			return true
		}
		// 字符串用完了，模式还剩字符没用
		// 只有一个字符的时候必然不可能
		if m == 1 {
			return false
		}
		// 不能空匹配
		if m > 1 && p[1] != '*' {
			return false
		}
		// 可以空匹配
		if m > 1 && p[1] == '*' {
			return isMatch(s[:], p[2:])
		}
	}
	// 首字母不匹配
	if !equal(s[0], p[0]) {
		// 首字母不匹配，但是可以空匹配
		if m > 1 && p[1] == '*' {
			return isMatch(s[:], p[2:])
		}
		// 首字母不匹配，且不能空匹配
		return false
	}
	// 首字母匹配，按照有没有 * 分成两种情况
	if m > 1 && p[1] == '*' {
		// 有 * 则可以匹配 0 个，也可以匹配 1 个，也可以匹配多个
		return isMatch(s[:], p[2:]) || isMatch(s[1:], p[2:]) || isMatch(s[1:], p[:])
	}
	// 没有 * 则返回下一个匹配
	return isMatch(s[1:], p[1:])
}
func equal(a, b byte) bool {
	if a == '.' || b == '.' || a == b {
		return true
	}
	return false
}
```
>执行用时：752 ms, 在所有 Go 提交中击败了5.05%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了77.00%的用户

和[剑指 Offer 19. 正则表达式匹配](..\剑指-Offer-19.-正则表达式匹配\README.md)是同一道题。不过这个时间有点惨烈啊，我对比看看是什么地方出错了。

找到了，还是那个多个匹配的时候的问题。
```go
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	// 字符串没用完，模式用完了
	if n > 0 && m == 0 {
		return false
	}
	if n == 0 {
		// 字符串和模式都用完了
		if m == 0 {
			return true
		}
		// 字符串用完了，模式还剩字符没用
		// 只有一个字符的时候必然不可能
		if m == 1 {
			return false
		}
		// 不能空匹配
		if m > 1 && p[1] != '*' {
			return false
		}
		// 可以空匹配
		if m > 1 && p[1] == '*' {
			return isMatch(s[:], p[2:])
		}
	}
	// 首字母不匹配
	if !equal(s[0], p[0]) {
		// 首字母不匹配，但是可以空匹配
		if m > 1 && p[1] == '*' {
			return isMatch(s[:], p[2:])
		}
		// 首字母不匹配，且不能空匹配
		return false
	}
	// 首字母匹配，按照有没有 * 分成两种情况
	if m > 1 && p[1] == '*' {
		// 有 * 则可以匹配 0 个，也可以匹配多个
		return isMatch(s[:], p[2:]) || isMatch(s[1:], p[:])
	}
	// 没有 * 则返回下一个匹配
	return isMatch(s[1:], p[1:])
}
func equal(a, b byte) bool {
	if a == '.' || b == '.' || a == b {
		return true
	}
	return false
}
```
>执行用时: 8 ms
内存消耗: 2.1 MB


具体来说就是这一段

```go
	// 首字母匹配，按照有没有 * 分成两种情况
	if m > 1 && p[1] == '*' {
		// 有 * 则可以匹配 0 个，也可以匹配多个
		return isMatch(s[:], p[2:]) || isMatch(s[1:], p[:])
    }
```

这里应该直接选择停止匹配多个，或是继续匹配更多，不用专门特意匹配一个的情况。因为会在后面的遍历中自动处理掉。

于是结合之前的代码，再优化一次逻辑
```go
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	// 字符串没用完，模式用完了
	if n > 0 && m == 0 {
		return false
	}
	// 匹配完成
	if n == 0 && m == 0 {
		return true
	}
	// 字符串用完
	if n == 0 {
		// 模式可以匹配 0 个字符
		if m > 1 && p[1] == '*' {
			return isMatch(s[:], p[2:])
		}
		// 模式不能继续匹配
		return false
	}
	// 首字母不匹配
	if !equal(s[0], p[0]) {
		// 首字母不匹配，但是可以空匹配
		if m > 1 && p[1] == '*' {
			return isMatch(s[:], p[2:])
		}
		// 首字母不匹配，且不能空匹配
		return false
	}
	// 首字母匹配，按照有没有 * 分成两种情况
	if m > 1 && p[1] == '*' {
		// 有 * 则可以匹配 0 个，也可以匹配多个
		return isMatch(s[:], p[2:]) || isMatch(s[1:], p[:])
	}
	// 没有 * 则返回下一个匹配
	return isMatch(s[1:], p[1:])
}
func equal(a, b byte) bool {
	if a == '.' || b == '.' || a == b {
		return true
	}
	return false
}
```
>执行用时：16 ms, 在所有 Go 提交中击败了24.27%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了79.17%的用户

竟然消耗比优化逻辑之前还大。这个也太奇怪了。