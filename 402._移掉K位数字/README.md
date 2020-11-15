# 402. 移掉K位数字
[https://leetcode-cn.com/problems/remove-k-digits/description/](https://leetcode-cn.com/problems/remove-k-digits/description/) 
## 单调栈
```go
func removeKdigits(num string, k int) string {
	n := len(num)
	if n == k {
		return "0"
	}
	queue := make([]byte, 1)
	queue[0] = num[0]

	for i := 1; i < n; i++ {
		nqueue := len(queue)
		// 调整单调栈
		hasAdjust := false
		for j := nqueue - 1; !hasAdjust && j >= 0; j-- {
			for k > 0 && j >= 0 && num[i] < queue[j] {
				queue = queue[:j]
				j--
				k--
				hasAdjust = true
			}
		}
		queue = append(queue, num[i])
	}
	res := "0"
	for i := range queue {
		if queue[i] != '0' {
			if k == 0 {
				res = string(queue[i:])
			} else {
				res = string(queue[i : len(queue)-k])
			}
			break
		}
	}
	return res
}
```
>执行用时：120 ms, 在所有 Go 提交中击败了5.48%的用户
内存消耗：2.7 MB, 在所有 Go 提交中击败了53.47%的用户

简单地说就是要得到尽可能小的数，就要让所有数位尽可能地呈现单调递增的情形。

为了能够做到单调递增，那我们就需要尽可能地去掉该数位之前的比它大的数，并且是越大的就得越先去掉。所以我们需要知道左边的数列的最大值，这个时候就可以用单调栈了。单调栈的应用还可以看另一道题[剑指 Offer 59 - I. 滑动窗口的最大值](..\剑指-Offer-59---I.-滑动窗口的最大值\README.md)。

简单地说就是单调栈可以做到在线性时间内查找之前的最大或最小元素，同时还不影响相对位置关系。

这题虽然做出来了，但是效率太慢了，应该是哪里的循环出了问题。


## 优化后的
```go
func removeKdigits(num string, k int) string {
	n := len(num)
	if n == k {
		return "0"
	}
	res := ""
	queue := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(queue) > 0 && digit < queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
			k--
		}
		queue = append(queue, digit)
	}
	queue = queue[:len(queue)-k]
	res = strings.TrimLeft(string(queue), "0")
	if res == "" {
		res = "0"
	}
	return res
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.6 MB, 在所有 Go 提交中击败了76.39%的用户

简单地说就是我之前的循环里面，比它大的应该直接插入，然后遍历到下一个元素。因此流程应该是

- 遍历元素
    1. 维护单调栈，把影响单调关系的元素去掉
    2. 插入当前元素。

因为为了方便实现取出结果，不用把结果再翻转一次（栈的输出顺序），所以用双向队列比较好（能够两边进出）。而 Go 语言的切片比较灵活，可以支持这样的操作，因此我们用切片就能比较好地实现操作了。