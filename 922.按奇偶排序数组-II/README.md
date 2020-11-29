# 922. 按奇偶排序数组 II
[https://leetcode-cn.com/problems/sort-array-by-parity-ii/](https://leetcode-cn.com/problems/sort-array-by-parity-ii/)
## 双指针
```go
func sortArrayByParityII(A []int) []int {
	a, b := 0, 1
	n := len(A)
	for a < n && b < n {
		for a < n && A[a]%2 == 0 {
			a += 2
		}
		for b < n && A[b]%2 != 0 {
			b += 2
		}
		if a < n && b < n {
			A[a], A[b] = A[b], A[a]
		}
		a += 2
		b += 2
	}
	return A
}
```
>执行用时：28 ms, 在所有 Go 提交中击败了53.30%的用户
内存消耗：6.3 MB, 在所有 Go 提交中击败了70.17%的用户

双指针，一个指向奇数位一个指向偶数位。指向奇数位的指针要满足这个指针指向的值为奇数，指向偶数位的指针要满足这个指针指向的值为偶数。

因此就不断地遍历指针，交换需要交换的位置。