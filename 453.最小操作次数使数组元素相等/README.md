# 453.最小操作次数使数组元素相等
[https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements) 
## 原题
给你一个长度为 `n` 的整数数组，每次操作将会使 `n - 1` 个元素增加 `1` 。返回让数组所有元素相等的最小操作次数。

 

 **示例 1：** 

```

输入：nums = [1,2,3]
输出：3
解释：
只需要3次操作（注意每次操作会增加两个元素的值）：
[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]

```
 **示例 2：** 

```

输入：nums = [1,1,1]
输出：0

```
 

 **提示：** 
-  `n == nums.length` 
-  `1 <= nums.length <= 10^5` 
-  `-10^9 <= nums[i] <= 10^9` 
- 答案保证符合 **32-bit** 整数
 
**标签**
`数组` `数学` 


## 找最小
```go
func minMoves(nums []int) int {
	minNum := min(nums)
	ans := 0

	for _, num := range nums {
		ans += num - minNum
	}

	return ans
}

func min(nums []int) int {
	ans := nums[0]

	for _, num := range nums[1:] {
		if ans > num {
			ans = num
		}
	}

	return ans
}
```
>执行用时：32 ms
内存消耗：6.6 MB

`n-1` 个元素增加 `1` 就是 `1` 个元素减 `1` 。因此找到最小元素，然后求出数组与最小元素的差即可。