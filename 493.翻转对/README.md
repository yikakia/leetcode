# 493. 翻转对
[https://leetcode-cn.com/problems/reverse-pairs/](https://leetcode-cn.com/problems/reverse-pairs/) 
## 
```go
func reversePairs(nums []int) int {
	if len(nums) == 0 { // 没有元素，没有翻转对
		return 0
	}
	count := 0                              // 个数
	mergeSort(nums, &count, 0, len(nums)-1) // 归并的范围 0到末尾
	return count
}

// 对当前的序列（从start到end）进行归并排序
func mergeSort(nums []int, count *int, start, end int) {
	if start == end { // 不能再二分，返回。递归的出口
		return
	}
	mid := start + (end-start)>>1 // 当前序列的中点索引

	mergeSort(nums, count, start, mid) // 递归左序列
	mergeSort(nums, count, mid+1, end) // 递归右序列

	// 此时左右序列已升序，现在，合并前的统计，以及合并
	i := start                 // 左序列的开头
	j := mid + 1               // 右序列的开头
	for i <= mid && j <= end { // i j 都不越界
		if nums[i] > 2*nums[j] {
			*count += mid - i + 1 // i 到 mid，都ok
			j++                   // 考察下一个j，继续找 i
		} else { // 当前i不满足，考察下一个i
			i++
		}
	}
	i = start
	j = mid + 1 // 复原 i j 指针，现在合并左右序列

	temp := make([]int, end-start+1) // 辅助数组，存放合并排序的数
	index := 0                       // 从0开始
	for i <= mid && j <= end {       // 如果 i j 都没越界
		if nums[i] < nums[j] { // nums[i]更小
			temp[index] = nums[i] // 取nums[i]，确定了temp[index]
			index++               // 确定下一个
			i++                   // 考察下一个i，j不动
		} else {
			temp[index] = nums[j]
			index++
			j++
		}
	}
	for i <= mid { // 如果 i 没越界，j越界了
		temp[index] = nums[i] // i 和 i右边的都取过来
		index++               // 确定下一个数
		i++
	}
	for j <= end { // j 没越界，i越界了
		temp[index] = nums[j] // j 和 j右边的都取过来
		index++               // 确定下一个数
		j++
	}
	k := 0
	for i := start; i <= end; i++ { // 根据合并后的情况，更新nums
		nums[i] = temp[k]
		k++
	}
}
```
>
