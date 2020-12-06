# 118. 杨辉三角
[https://leetcode-cn.com/problems/pascals-triangle/](https://leetcode-cn.com/problems/pascals-triangle/) 
## 递归的方式
```go
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{[]int{1}}
	}
	res := generate(numRows - 1)
	lastRow := res[len(res)-1]
	curRow := make([]int, numRows)
	for i := 0; i < len(lastRow)-1; i++ {
		curRow[i+1] = lastRow[i] + lastRow[i+1]
	}
	curRow[0], curRow[numRows-1] = 1, 1
	res = append(res, curRow)
	return res
}
```
>执行用时: 0 ms
内存消耗: 2.1 MB

用了递归，所以内存占用大了点。
