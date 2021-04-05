# 88.合并两个有序数组
[https://leetcode-cn.com/problems/merge-sorted-array](https://leetcode-cn.com/problems/merge-sorted-array) 
## 原题
给你两个有序整数数组  `nums1` 和 `nums2` ，请你将 `nums2` 合并到  `nums1` * * 中 *，* 使 `nums1` 成为一个有序数组。

初始化  `nums1` 和 `nums2` 的元素数量分别为  `m` 和 `n` 。你可以假设  `nums1` 的空间大小等于  `m + n` ，这样它就有足够的空间保存来自 `nums2` 的元素。

 

 **示例 1：** 

```

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]

```
 **示例 2：** 

```

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]

```
 

 **提示：** 
-  `nums1.length == m + n` 
-  `nums2.length == n` 
-  `0 <= m, n <= 200` 
-  `1 <= m + n <= 200` 
-  `-10^9 <= nums1[i], nums2[i] <= 10^9` 
 
**标签**
`数组` `双指针` 


## 从后往前找
```go
func merge(nums1 []int, m int, nums2 []int, n int) {
	a, b := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		switch {
		case a == -1:
			nums1[i] = nums2[b]
			b--
		case b == -1:
			nums1[i] = nums1[a]
			a--
		case nums1[a] > nums2[b]:
			nums1[i] = nums1[a]
			a--
		default:
			nums1[i] = nums2[b]
			b--
		}
	}
}
```
>执行用时: 0 ms
内存消耗: 2.3 MB

看到题目里面专门说 `nums1` 能够放下 `m+n` 就直到要考的是原地更新。这里从后往前进行归并，把大的放到尾部即可。

当一边的指针到底了之后，就不断地把另外一边的结果赋进去即可。