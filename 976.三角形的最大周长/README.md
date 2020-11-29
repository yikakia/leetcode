# 976. 三角形的最大周长
[https://leetcode-cn.com/problems/largest-perimeter-triangle/](https://leetcode-cn.com/problems/largest-perimeter-triangle/) 
## 降序排序后算最大
```go
func largestPerimeter(A []int) (res int) {
	if len(A) < 3 {
		return
	}
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	for i := 0; i < len(A)-2; i++ {
		if A[i] < A[i+1]+A[i+2] {
			res = A[i] + A[i+1] + A[i+2]
			break
		}
	}
	return
}
```
>执行用时: 40 ms
内存消耗: 6.6 MB

简单地说就是降序排序后从最长的边$c$开始找，然后找次长边$b$和次次长边$a$。因为这个时候的$a$和$b$都不能和$c$组成三角形的话，那么比$a$和$b$还小的边就更不可能和$c$组成三角形了