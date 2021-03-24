# 456.132模式
[https://leetcode-cn.com/problems/132-pattern](https://leetcode-cn.com/problems/132-pattern) 
## 原题
给定一个整数序列：a<sub>1</sub>, a<sub>2</sub>, ..., a<sub>n</sub>，一个132模式的子序列 a<sub> **i** </sub>, a<sub> **j** </sub>, a<sub> **k** </sub> 被定义为：当 **i** < **j** < **k** 时，a<sub> **i** </sub> < a<sub> **k** </sub> < a<sub> **j** </sub>。设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有132模式的子序列。

 **注意：** n 的值小于15000。

 **示例1:** 

```

输入: [1, 2, 3, 4]

输出: False

解释: 序列中不存在132模式的子序列。

```
 **示例 2:** 

```

输入: [3, 1, 4, 2]

输出: True

解释: 序列中有 1 个132模式的子序列： [1, 4, 2].

```
 **示例 3:** 

```

输入: [-1, 3, 2, 0]

输出: True

解释: 序列中有 3 个132模式的的子序列: [-1, 3, 2], [-1, 3, 0] 和 [-1, 2, 0].

```
 
**标签**
`栈` 


## 单调栈
```go
func find132pattern(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64

	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}

	return false
}
```
>执行用时: 12 ms
内存消耗: 5.3 MB

原理就是用一个单调栈维护 **最大** 的 **次小值** ，换言之就是维护 `132模式` 中的 `2模式` 。让 `2模式` 尽可能地大，就能更可能找到合适的 `1模式` 

这里是从右往左遍历比较，这样遍历到的 `2模式` 即 `k` 都可以被左边的元素进行使用。

于是当 `2模式` 的最大值大于了当前遍历到的元素的时候，就说明 
```math
i<k 且 nums[i]<nums[j]
```

而全部遍历了都找不到的话就说明不存在 直接返回 `false` 即可。

