# 剑指 Offer 29. 顺时针打印矩阵
[https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/](https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/)

## 模拟
```go
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
    }
	r := make([]int, len(matrix)*len(matrix[0]))
	x, y := 0, 0
	rx, ry := 0, 0
	for i := 0; i < len(matrix)*len(matrix[0]); {
		for y+ry < len(matrix[0]) {
			r[i] = matrix[x][y]
			y++
			i++
		}
		if i >= len(matrix)*len(matrix[0]) {
			break
		}
		y--
		x++
		for x+rx < len(matrix) {
			r[i] = matrix[x][y]
			x++
			i++
		}
		if i >= len(matrix)*len(matrix[0]) {
			break
		}
		x--
		rx++
		y--
		for y-ry > -1 {
			r[i] = matrix[x][y]
			y--
			i++
		}
		if i >= len(matrix)*len(matrix[0]) {
			break
		}
		y++
		x--
		ry++
		for x-rx > -1 {
			r[i] = matrix[x][y]
			x--
			i++
		}
		if i >= len(matrix)*len(matrix[0]) {
			break
		}
		x++
		y++
	}
	return r
}
```

>执行用时：16 ms, 在所有 Go 提交中击败了45.16%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了49.74%的用户

就是很简单的模拟不断地沿👉👇👈👆的方向走。注意边界条件就是了。

然后发现时间太长，于是优化了一下，换成了和常量比较
```go

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	lenth := len(matrix) * len(matrix[0])
	r := make([]int, lenth)
	x, y := 0, 0
	rx, ry := 0, 0
	for i := 0; i < lenth; {
		for y+ry < len(matrix[0]) {
			r[i] = matrix[x][y]
			y++
			i++
		}
		if i >= lenth {
			break
		}
		y--
		x++
		for x+rx < len(matrix) {
			r[i] = matrix[x][y]
			x++
			i++
		}
		if i >= lenth {
			break
		}
		x--
		rx++
		y--
		for y-ry > -1 {
			r[i] = matrix[x][y]
			y--
			i++
		}
		if i >= lenth {
			break
		}
		y++
		x--
		ry++
		for x-rx > -1 {
			r[i] = matrix[x][y]
			x--
			i++
		}
		if i >= lenth {
			break
		}
		x++
		y++
	}
	return r
}
```
>执行用时：12 ms, 在所有 Go 提交中击败了88.02%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了53.66%的用户

不过有些时候也会跳到 16 ms 所以这纯粹是娱乐的吧。