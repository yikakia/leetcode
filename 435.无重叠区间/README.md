# 435. 无重叠区间
[https://leetcode-cn.com/problems/non-overlapping-intervals/](https://leetcode-cn.com/problems/non-overlapping-intervals/) 
## 原题
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
**注意:** 
- 可以认为区间的终点总是大于它的起点。
- 区间 [1,2] 和 [2,3] 的边界相互&ldquo;接触&rdquo;，但没有相互重叠。
**示例 1:** 
```
输入: [ [1,2], [2,3], [3,4], [1,3] ]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。
```
**示例 2:** 
```
输入: [ [1,2], [1,2], [1,2] ]
输出: 2
解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
```
**示例 3:** 
```
输入: [ [1,2], [2,3] ]
输出: 0
解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
```


## 
```go
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	res := 1
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	maxRight := intervals[0][1]
	for _, interval := range intervals {
		if interval[0] >= maxRight {
			maxRight = interval[1]
			res++
		}
	}

	return len(intervals) - res
}

func test() {
	type TestType struct {
		intervals [][]int
		want      int
	}
	ts := []TestType{
		TestType{
			intervals: [][]int{
				[]int{1, 2},
				[]int{2, 3},
				[]int{3, 4},
				[]int{1, 3},
			},
			want: 1,
		},
		TestType{
			intervals: [][]int{
				[]int{1, 2},
				[]int{1, 2},
				[]int{1, 2},
			},
			want: 2,
		},
		TestType{
			intervals: [][]int{
				[]int{1, 2},
				[]int{2, 3},
			},
			want: 0,
		},
	}
	for _, t := range ts {
		get := eraseOverlapIntervals(t.intervals)
		if get != t.want {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")
}
```
>执行用时: 4 ms
内存消耗: 3.9 MB


这个和[452. 用最少数量的箭引爆气球](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/) 非常相似，唯一不同的是一个要求边界不能重叠（气球），一个边界可以重叠（本题）。
