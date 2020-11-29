# 11. 盛最多水的容器
[https://leetcode-cn.com/problems/container-with-most-water/](https://leetcode-cn.com/problems/container-with-most-water/) 
## 双指针
```go
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxArea(height []int) int {
	l, r := 0, len(height)
	res := 0
	for l < r {
		tmp := (r - l) * min(height[l], height[r])
		if tmp > res {
			res = tmp
		}
		if height[l] < height[r] {
			l++
		} else if height[l] == height[r] {
			l++
			r--
		} else if height[l] > height[r] {
			r--
		} else {
			panic("unreachable")
		}
	}
	return res
}
```
>执行用时: 20 ms
内存消耗: 6.3 MB


从两边逼近，总是移动最小的那一个指针，
