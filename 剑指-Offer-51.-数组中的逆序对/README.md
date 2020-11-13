# 剑指 Offer 51. 数组中的逆序对
[https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)

## 归并排序
```go
func reversePairs(nums []int) int {
	n := len(nums)

	return mergeSort(nums, 0, n-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	count := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	i, j := start, mid+1
	tmp := []int{}
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			count += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	count += (end - (mid + 1) + 1) * (mid - i + 1)
	if i <= mid {
		tmp = append(tmp, nums[i:mid+1]...)
	}
	if j <= end {
		tmp = append(tmp, nums[j:end+1]...)
	}

	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}

	return count
}
```
>执行用时：164 ms, 在所有 Go 提交中击败了33.82%的用户
内存消耗：8.8 MB, 在所有 Go 提交中击败了20.23%的用户

不看题解根本想不到还能用归并排序来做这道题。归并排序的时候通过对比左右两边的大小来得到是否构成逆序对以及多少个逆序对。

为了实现归并排序所以还得有个临时数组来储存排序的结果。要优化的话就是提前初始化好这个临时数组，这样就能省掉递归过程中的内存占用。