# 659. 分割数组为连续子序列
[https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/](https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/) 
## 题目
给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个子序列，其中每个子序列都由连续整数组成且长度至少为 3 。

如果可以完成上述分割，则返回 true ；否则，返回 false 。


示例 1：
```
输入: [1,2,3,3,4,5]
输出: True
解释:
你可以分割出这样两个连续子序列 : 
1, 2, 3
3, 4, 5
```

示例 2：
```
输入: [1,2,3,3,4,4,5,5]
输出: True
解释:
你可以分割出这样两个连续子序列 : 
1, 2, 3, 4, 5
3, 4, 5
```
 

示例 3：
```
输入: [1,2,3,4,4,5]
输出: False
```

提示：

- 输入的数组长度范围为 [1, 10000]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
## 贪心
```go
func isPossible(nums []int) bool {
	// countMap 用于统计每个还未被划分的数的个数
	countMap := make(map[int]int, 0)
	for _, num := range nums {
		countMap[num]++
	}
	// endMap 用于统计该数字作为终点的子序列个数
	endMap := make(map[int]int, 0)
	for _, num := range nums {
		if countMap[num] > 0 {
			if endMap[num-1] > 0 {
				countMap[num]--
				endMap[num-1]--
				endMap[num]++
			} else {
				if countMap[num+1] > 0 && countMap[num+2] > 0 {
					countMap[num]--
					countMap[num+1]--
					countMap[num+2]--
					endMap[num+2]++
				} else {
					return false
				}
			}
		}
	}
	return true
}
```
>执行用时: 100 ms
内存消耗: 7 MB

简单来说就是一个数$x$，既可以插入到以 $x-1$结尾的子序列中，也可以作为新的子序列。因为要求最小长度为3，因此应该尽可能地添加到已有的子序列末尾而不是新建一个子序列。

因为如果是
- 尽可能多的生成子序列的话，就会遇到`1,2,3,4,5,6,7`这样数列。
- 尽可能生成长的子序列的话，就会遇到`1,2,3,3,4,5`这样的数列。

于是我们就可以用两个 `Map` 来记录。一个`countMap`记录这个数还有多少不在子序列中，一个`endMap`记录这个数作为子序列的终点的序列有几个。每次在第一个Map中找到不为0的值，然后把它插入到合适的子序列中

- 遍历 `nums` 数组
    - 如果 `num` 在 `countMap` 中不为 0 ，那么就把它插入到子序列中
        - 如果 `endMap[num-1]` 不为 0 ，那么就接在这个序列的后面，并且更新两个Map
        - 如果找不到能够插入的子序列，那么就让它自己生成一个长度为3的子序列
            - 如果不能生成，那么就返回 `false`
            - 如果可以，那么就更新两个`Map`
- 之前没有不能生成的情况，那么每个节点都在一个子序列中。返回 `true`

这个自己想的时候就差了一步，就是生成序列的方式错了。我是想了我上面说的那两个错误的方式，结果都能很轻松地找到反例。归根结底就是题刷得还不够多，不能直接找到合适的方法。