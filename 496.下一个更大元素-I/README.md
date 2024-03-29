# 496.下一个更大元素 I
[https://leetcode-cn.com/problems/next-greater-element-i](https://leetcode-cn.com/problems/next-greater-element-i) 
## 原题
给你两个 **没有重复元素** 的数组  `nums1` 和  `nums2`  ，其中 `nums1`  是  `nums2`  的子集。

请你找出 `nums1`  中每个元素在  `nums2`  中的下一个比其大的值。

 `nums1`  中数字  `x`  的下一个更大元素是指  `x`  在  `nums2`  中对应位置的右边的第一个比  `x` 大的元素。如果不存在，对应位置输出 `-1` 。

 

 **示例 1:** 

```

输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
    对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
    对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
```
 **示例 2:** 

```

输入: nums1 = [2,4], nums2 = [1,2,3,4].
输出: [3,-1]
解释:
    对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
    对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。

```
 

 **提示：** 
-  `1 <= nums1.length <= nums2.length <= 1000` 
-  `0 <= nums1[i], nums2[i] <= 10^4` 
-  `nums1` 和 `nums2` 中所有整数 **互不相同** 
-  `nums1` 中的所有整数同样出现在 `nums2` 中
 

 **进阶：** 你可以设计一个时间复杂度为 `O(nums1.length + nums2.length)` 的解决方案吗？

 
**标签**
`栈` `数组` `哈希表` `单调栈` 


## 单调栈
```go
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextMap := map[int]int{}
	stack := []int{nums2[0]}

	for _, num := range nums2[1:] {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextMap[pre] = num
		}
		stack = append(stack, num)
	}

	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		if v, ok := nextMap[num]; ok {
			ans[i] = v
		} else {
			ans[i] = -1
		}
	}
	return ans
}
```
>执行用时：4 ms
内存消耗：3.4 MB

简单的做法就是从前往后每个搜一遍，要时间复杂度为 `O(nums1.length + nums2.length)` 的解决方案就需要用单调栈了，每次都把前面待匹配的解决掉。