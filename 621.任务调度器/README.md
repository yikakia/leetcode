# 621. 任务调度器
[https://leetcode-cn.com/problems/task-scheduler/](https://leetcode-cn.com/problems/task-scheduler/) 
## 模拟
```go
func leastInterval(tasks []byte, n int) (res int) {

	countMap := make(map[byte]int, 26)
	for _, task := range tasks {
		countMap[task]++
	}

	type Pair struct {
		char  byte
		count int
	}
	pairs := []Pair{}
	for char, count := range countMap {
		pairs = append(pairs, Pair{char: char, count: count})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	needWait := n + 1
	finalCount := 0
	for res < len(tasks)*(n+1) {
		needWait = n + 1
		hasChanged := false
		for i := range pairs {
			if pairs[i].count > 0 && needWait > 0 {
				if !hasChanged {
					hasChanged = true
				}
				pairs[i].count--
				needWait--
				res++
			}
		}
		if !hasChanged {
			return res - finalCount
		}
		if needWait > 0 {
			res += needWait
			finalCount = needWait
		}
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].count > pairs[j].count
		})
	}
	return res - finalCount
}
```
>执行用时: 24 ms
内存消耗: 6.4 MB

模拟时间片操作，按元素多少排序，不断地按大小放入。最后退出的时候要减去最后一轮时间片中的空的部分。

## 
```go
func leastInterval(tasks []byte, n int) int {

	countMap := make(map[byte]int, 26)
	for _, task := range tasks {
		countMap[task]++
	}
	maxCount, maxNum := 0, 0
	for _, cnt := range countMap {
		if cnt > maxCount {
			maxCount, maxNum = cnt, 1
		} else if cnt == maxCount {
			maxNum++
		}
	}
	if res := (maxCount-1)*(n+1) + maxNum; res > len(tasks) {
		return res
	}
	return len(tasks)
}
```
>执行用时: 16 ms
内存消耗: 6.2 MB

算出元素最多的次数`maxCount`和个数`maxNum`，结果就是`max((maxCount-1)*(n+1) + maxNum , len(tasks))`