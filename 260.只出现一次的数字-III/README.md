# 260.只出现一次的数字 III
[https://leetcode-cn.com/problems/single-number-iii](https://leetcode-cn.com/problems/single-number-iii) 
## 原题
给定一个整数数组  `nums` ，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 **任意顺序** 返回答案。

 

 **进阶：** 你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

 

 **示例 1：** 

```

输入：nums = [1,2,1,3,2,5]
输出：[3,5]
解释：[5, 3] 也是有效的答案。

```
 **示例 2：** 

```

输入：nums = [-1,0]
输出：[-1,0]

```
 **示例 3：** 

```

输入：nums = [0,1]
输出：[1,0]

```
 **提示：** 
-  `2 <= nums.length <= 3 * 10^4` 
-  `-2^31 <= nums[i] <= 2^31 - 1` 
- 除两个只出现一次的整数外， `nums` 中的其他数字都出现两次
 
**标签**
`位运算` 


## 位运算
```go
func singleNumber(nums []int) []int {
	var ab, a, b int
	for _, num := range nums {
		ab ^= num
	}
	var mask int = 1
	for mask&ab == 0 {
		mask <<= 1
	}
	for _, num := range nums {
		if num&mask == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
```
>执行用时: 8 ms
内存消耗: 3.9 MB

这道题和 [剑指-Offer-56---I.数组中数字出现的次数](../剑指-Offer-56---I.数组中数字出现的次数/README.md) 是一样的。它的进阶版是 [645.错误的集合](../645.错误的集合/README.md)