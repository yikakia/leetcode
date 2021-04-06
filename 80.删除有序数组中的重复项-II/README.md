# 80.删除有序数组中的重复项 II
[https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii) 
## 原题
给你一个有序数组 `nums` ，请你 **<a href="http://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95" target="_blank"> 原地</a>** 删除重复出现的元素，使每个元素 **最多出现两次** ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 **<a href="https://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95" target="_blank">原地 </a>修改输入数组** 并在使用 O(1) 额外空间的条件下完成。

 

 **说明：** 

为什么返回数值是整数，但输出的答案是数组呢？

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

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 不需要考虑数组中超出新长度后面的元素。

```
 **示例 2：** 

```

输入：nums = [0,0,1,1,1,1,2,3,3]
输出：7, nums = [0,0,1,1,2,3,3]
解释：函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。 不需要考虑数组中超出新长度后面的元素。

```
 

 **提示：** 
-  `1 <= nums.length <= 3 * 10^4` 
-  `-10^4 <= nums[i] <= 10^4` 
-  `nums` 已按升序排列
 
**标签**
`数组` `双指针` 


## 移动数组
```go
func removeDuplicates(nums []int) (newlen int) {
	st := 0
	moved := 0
	n := len(nums)
	if n < 3 {
		return n
	}
	for st < n-moved {
		end := st + 1
		for end < n-moved && nums[end] == nums[st] {
			end++
		}
		if end == n-moved {
			if st == n-moved-1 {
				return st + 1
			}
			return st + 2
		}
		if end-st > 2 {
			copy(nums[st+2:], nums[end:])
			moved += end - st - 2
			nums[n-moved]--
			st += 2
		} else {
			st = end
		}
		newlen = end
	}
	return
}
```
>执行用时: 4 ms
内存消耗: 3 MB

直接就移动数组即可。关键是记录偏移量，即之前移动了多少个元素，遍历到 `n-moved` 即可。

主要思路就是先记录区间，然后根据区间来进行移动或者返回结果。

## 双指针

```go
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}
	fast, slow := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
```
>执行用时: 4 ms
内存消耗: 3 MB

这里用 `slow` 代表待填入的位置，用 `fast` 不断地往后遍历，如果满足可以填入的条件，则进行进行填入，并且把 `slow` 向后移一，不然的话 `slow` 就继续等待填充合适的元素。

可以填入的条件是 `nums[slow-2]!=nums[fast]` 即待填入的位置的上上个元素和当前的元素不同。这里的思路是数组中的每个元素都和它的上上个元素不同，因为是升序的，所以最多能有两个相同的存在。所以当遍历到的元素满足 `nums[slow-2]!=nums[fast]` 的时候，就说明这个是适合填入的元素，直接填入即可。

这里的方法有点类似于持果索因。通过最后得到的数列的性质来决定创建时候的方法。