# 5618. K 和数对的最大数目
[https://leetcode-cn.com/problems/max-number-of-k-sum-pairs/](https://leetcode-cn.com/problems/max-number-of-k-sum-pairs/) 
## 哈希
```go
func maxOperations(nums []int, k int) int {
	countMap := make(map[int]int, 0)
	res := 0
	for _, num := range nums {
		if count, ok := countMap[k-num]; ok && count > 0 {
			res++
			countMap[k-num]--
		} else {
			countMap[num]++
		}
	}
	return res
}
```
>执行用时: 136 ms
内存消耗: 10.3 MB

类似于两数之和。
