# 228.汇总区间
[https://leetcode-cn.com/problems/summary-ranges/](https://leetcode-cn.com/problems/summary-ranges/) 
## 原题
给定一个无重复元素的有序整数数组 `nums` 。

返回 **恰好覆盖数组中所有数字**  的 **最小有序**  区间范围列表。也就是说，`nums` 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 `nums` 的数字 `x` 。

列表中的每个区间范围 `[a,b]` 应该按如下格式输出：
- `"a->b"` ，如果 `a != b`
- `"a"` ，如果 `a == b`
 

**示例 1：** 

```

输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"

```
**示例 2：** 

```

输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
解释：区间范围是：
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

```
**示例 3：** 

```

输入：nums = []
输出：[]

```
**示例 4：** 

```

输入：nums = [-1]
输出：["-1"]

```
**示例 5：** 

```

输入：nums = [0]
输出：["0"]

```
 

**提示：** 
- `0 <= nums.length <= 20`
- `-2^31 <= nums[i] <= 2^31 - 1`
- `nums` 中的所有值都 **互不相同** 
- `nums` 按升序排列
 
**标签**
`数组` 


## 
```go
func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return []string{}
	}
	ranges := make([][]int, 0)
	ranges = append(ranges, []int{})
	ranges[0] = append(ranges[0], nums[0])

	for i := 1; i < n; i++ {
		num := nums[i]
		m1 := len(ranges) - 1
		m2 := len(ranges[m1]) - 1
		if num-ranges[m1][m2] == 1 {
			if m2 == 0 {
				ranges[m1] = append(ranges[m1], num)
			} else {
				ranges[m1][m2] = num
			}
			continue
		}
		ranges = append(ranges, []int{num})
	}

	res := []string{}
	for i := range ranges {
		if len(ranges[i]) == 2 {
			res = append(res, strconv.Itoa(ranges[i][0])+"->"+strconv.Itoa(ranges[i][1]))
		} else {
			res = append(res, strconv.Itoa(ranges[i][0]))
		}
	}
	return res
}
```
>执行用时: 0 ms
内存消耗: 1.9 MB

纯粹的模拟，如果差值超过`1`就说明这个数是新的节点。
