# 面试题 17.21.直方图的水量
[https://leetcode-cn.com/problems/volume-of-histogram-lcci](https://leetcode-cn.com/problems/volume-of-histogram-lcci) 
## 原题
给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为 1。

<img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png" style="height: 161px; width: 412px;">

<small>上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的直方图，在这种情况下，可以接 6 个单位的水（蓝色部分表示水）。 **感谢 Marcos** 贡献此图。</small>

 **示例:** 

```
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
```
 
**标签**
`栈` `数组` `双指针` 


## 单调栈
```go
func trap(height []int) (ans int) {
	stk := []int{}
	for i, h := range height {
		for len(stk) > 0 && h > height[stk[len(stk)-1]] {
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if len(stk) == 0 {
				break
			}
			left := stk[len(stk)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stk = append(stk, i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
>执行用时: 4 ms
内存消耗: 2.8 MB

怎么才能兜住水呢？需要一个凹陷，而且这个凹陷的左右都需要能够拦住的。当遇到了凹陷的时候，就计算出当前这个凹陷能够覆盖的长度即可。
