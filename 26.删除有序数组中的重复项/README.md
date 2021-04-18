# 26.删除有序数组中的重复项
[https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array) 
## 原题
给你一个有序数组 `nums` ，请你 **<a href="http://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95" target="_blank"> 原地</a>** 删除重复出现的元素，使每个元素 **只出现一次** ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 **<a href="https://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95" target="_blank">原地 </a>修改输入数组** 并在使用 O(1) 额外空间的条件下完成。

 

 **说明:** 

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以 **「引用」** 方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

```

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

```

 

 **示例 1：** 

```

输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

```
 **示例 2：** 

```

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

```
 

 **提示：** 
-  `0 <= nums.length <= 3 * 10^4` 
-  `-10^4 <= nums[i] <= 10^4` 
-  `nums` 已按升序排列
 

 
**标签**
`数组` `双指针` 


## 双指针
```go
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	i := 1
	for j := 1; j < n; j++ {
		if nums[j] != nums[i-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
```
>执行用时: 8 ms
内存消耗: 4.5 MB

`i` 指向当前需要填充的下标，`j`代表遍历的元素。而`i-1`就表示的是之前那个更新的元素。只有和之前那个元素不同的时候，才需要更新当前的元素。而最后 `i` 就刚好代表有多少个不重复的元素。