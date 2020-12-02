# 321. 拼接最大数
[https://leetcode-cn.com/problems/create-maximum-number/](https://leetcode-cn.com/problems/create-maximum-number/) 
## 单调栈与归并排序
```go
/* 输入的数据为
- 源数组 **a**

- 需要切分的长度**k**  要求 k < len(a)

- 返回数组 **s** 长度为 **k**
*/
func maxSubsequence(a []int, k int) (s []int) {
	for i, v := range a {
		// 当数组剩下的元素与栈内元素数量之和大于等于要求
		// 的长度 k 的时候就可以判断是不是大于末尾的元素，
		// 如果大于就可以不断地将最后的部分出栈。否则就把
		// 后面的加入进来就可以了。
		for len(s) > 0 && len(s)+len(a)-1-i >= k && v > s[len(s)-1] {
			s = s[:len(s)-1]
		}
		if len(s) < k {
			s = append(s, v)
		}
	}
	return
}

/*lexicographicalLess

返回 a,b 数组是否为字典序意义上的 a<b
*/
func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

/* merge 用于归并排序

按照字典序的顺序进行比较
*/
func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for i := range merged {
		if lexicographicalLess(a, b) {
			merged[i], b = b[0], b[1:]
		} else {
			merged[i], a = a[0], a[1:]
		}
	}
	return merged
}

func maxNumber(nums1, nums2 []int, k int) (res []int) {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := maxSubsequence(nums1, i)
		s2 := maxSubsequence(nums2, k-i)
		merged := merge(s1, s2)
		if lexicographicalLess(res, merged) {
			res = merged
		}
	}
	return
}
```
>执行用时: 16 ms
内存消耗: 6.9 MB

明明想到了单调栈的，但是没想出来怎么实现两个的合并。于是就用了官方的题解了。

简单地说就是分别在两个数组里面取 **i** 个数 和 **k-i** 个数，这一步就是用的单调栈。

然后再把它们合并起来，合并的时候是看它们剩下未合并部分的数组的字典序，因为都是单个数字，所以也可以看作是剩下部分组成的数字的大小。

最后就是判断结果数组和合并之后的数组的字典序大小了，维护为最大的结果就行。