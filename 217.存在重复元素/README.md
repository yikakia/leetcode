# 217. 存在重复元素
[https://leetcode-cn.com/problems/contains-duplicate/](https://leetcode-cn.com/problems/contains-duplicate/) 
## 原题
给定一个整数数组，判断是否存在重复元素。
如果任意一值在数组中出现至少两次，函数返回 `true` 。如果数组中每个元素都不相同，则返回 `false` 。
 
**示例 1:**
```
输入: [1,2,3,1]
输出: true
```
**示例 2:** 
```
输入: [1,2,3,4]
输出: false
```
**示例 3:** 
```
输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
```


## 哈希表
```go
func containsDuplicate(nums []int) bool {
	countMap := make(map[int]int, 0)
	for _, num := range nums {
		if _, ok := countMap[num]; ok {
			return true
		}
		countMap[num]++
	}
	return false
}
```
>执行用时: 28 ms
内存消耗: 7.7 MB

## 排序
```go
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
```
>执行用时: 24 ms
内存消耗: 6 MB

可以排序的时间复杂度和空间复杂度都要好一点。时间复杂度的话应该主要是哈希表开辟空间，以及索引可能会要耗时间。

