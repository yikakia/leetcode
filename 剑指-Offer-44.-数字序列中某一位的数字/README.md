# 剑指 Offer 44. 数字序列中某一位的数字
[https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/](https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/)

## 逐步找到要输出的数
```go
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	n -= 10
	cur := 90
	digit := 1
	width := digit + 1
	for n-cur*width > 0 {
		n -= cur * width
		cur *= 10
		digit++
		width++
	}

	quo := n/width + cur/9
	rem := n % width

	return nInNum(quo, rem)
}
func nInNum(num, n int) int {
	s := strconv.Itoa(num)
	r, _ := strconv.Atoi(string(s[n]))
	return r
}


```

>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了65.32%的用户

简单地说就是有两个依据
- 十位数的字符宽度为2，百位数的字符宽度为3

- 十位数的时候一共有 90 个数字，百位数的时候一共有 900 个数字。

```go
//0 ~ 9 10
//10 ~ 99 90*2
//100 ~ 999 900*3
//1000 ~ 9999 9000*4
```
以此为根据，逐步找到应该输出的是哪个数字。以及是这个数字的第几位数字。

主要就是根据字符宽度和这个数位总共有多少个数来确定到底应该输出哪一个数字的第几位。