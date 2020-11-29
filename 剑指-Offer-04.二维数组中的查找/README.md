# 剑指 Offer 04. 二维数组中的查找
[https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)

## 二分查找
```go
func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	if matrix[n-1][m-1] < target {
		return false
	}
	if matrix[0][0] > target {
		return false
	}
	getMedium := func(a, b int) int {
		return (a + b) / 2
	}
	tag := 0
	for y := m - 1; y >= 0; y-- {
		if matrix[0][y] > target || matrix[n-1][y] < target {
			continue
		}
		minx := tag
		maxx := n
		for {
			medium := getMedium(minx, maxx)
			if matrix[medium][y] > target {
				maxx = medium - 1
			} else if matrix[medium][y] == target {
				return true
			} else if matrix[medium][y] < target {
				minx = medium + 1
			}
			if minx > maxx {
				tag = minx
				break
			}
		}
	}

	return false
}
```

>执行用时：28 ms, 在所有 Go 提交中击败了88.81%的用户
内存消耗：6.6 MB, 在所有 Go 提交中击败了36.38%的用户


根据题意可知这个数组满足`nums[x][y]`大于所有 `nums[0:x-1][0:y-1]` 这个性质。因此我们可以根据这个来寻找需要的数。

对于优化而言，我们从最右边开始搜索，当找到了一个位置`(x,y)`使得`nums[x-1][y]<= target <= nums[x+1][y]` 。简单地说就是找到目标数的下界，之后的搜索可以通过这个来作为一个终止条件，换句话说就是进行了剪枝操作。

还有一个地方可以优化，那就是使用二分查找的方式。因为数组本来就是有序的了，那么我们可以通过二分的方式来找到那个目标数的下界来进行快速的剪枝操作。

对于在搜索目标数的时候，我们还有种优化的方式就是将目标数和数组的最小值与最大值进行判断。从题意可以得知数组的左上角是数组的最小值，数组的右下角是数组的最大值。我们可以先进行一个判断看目标数是否在这个数组里面，即将目标数和数组的左上角与右下角分别进行比较，看目标数是否在数组中。

当然了，其实 golang 本身是有搜索的函数 `sort.SearchInts(nums,target)` 的。用这个替换查找的部分可以大大减少工作量，毕竟是系统库提供的函数。虽然比起手写而言少了确定下界的剪枝，所以性能会多消耗一丢丢。

简单地说就是这样
```go
func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	if matrix[n-1][m-1] < target {
		return false
	}
	if matrix[0][0] > target {
		return false
	}
	for _, nums := range matrix {
		if nums[m-1] < target {
			continue
		}
		index := sort.SearchInts(nums, target)
		if index < len(nums) && target == nums[index] {
			return true
		}
	}
	return false
}
```

##  坑点
这些写完了，还有就是坑点了。主要就是对于`[]` 和 `[[]]`这样的空数组，应该即刻返回 `false`。还有就是二分时候应该怎么取左边和右边的值。