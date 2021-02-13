# 448.找到所有数组中消失的数字
[https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/) 
## 原题
给定一个范围在  1 ≤ a[i] ≤ *n* ( *n* = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。

找到所有在 [1, *n*] 范围之间没有出现在数组中的数字。

您能在不使用额外空间且时间复杂度为 *O(n)* 的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。

 **示例:** 

```

输入:
[4,3,2,7,8,2,3,1]

输出:
[5,6]

```
 
**标签**
`数组` 


## 遍历修改
```go
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	res := []int{}
	for _, v := range nums {
		nums[(v-1)%n] += n
	}
	for i, v := range nums {
		if v <= n {
			res = append(res, i+1)
		}
	}

	return res
}
```
>执行用时: 60 ms
内存消耗: 7.4 MB

原理是将出现过的数的位置上的数加到`n`之上。遍历一遍后不大于`n`的就是没出现过的数。

## 
```go
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for _, vi := range nums {
		if vi != 0 {
			var cur int
			for j := vi - 1; nums[j] != 0; {
				cur = nums[j] - 1
				nums[j] = 0
				j = cur
			}
		}
	}
	for i := range nums {
		if nums[i] != 0 {
			res = append(res, i+1)
		}
	}
	return res
}
```
>执行用时: 60 ms
内存消耗: 7.4 MB

思路差不多，但是这次用了类似递归的思路，把出现过的数的位置上的数变为 `0` 。不过因为这个数指向的位置上可能还是有数，所以要用到类似于递归的操作，直到要修改的那一位上的数本来就是 `0` 为止。
这样操作之后判断该数位是不是 `0` 即可。