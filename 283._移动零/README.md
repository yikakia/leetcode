# 283. 移动零
[https://leetcode-cn.com/problems/move-zeroes/](https://leetcode-cn.com/problems/move-zeroes/) 
## 记录 0 的个数，并直接用非 0 值覆盖
```go
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	cnt := 0
	if nums[0] == 0 {
		cnt++
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt++
		} else {
			nums[i-cnt] = nums[i]
		}
	}
	for i := len(nums) - cnt; i < len(nums); i++ {
		nums[i] = 0
	}
}
```
>执行用时: 4 ms
内存消耗: 3.7 MB

和官方的题解其实差不多，不过我这个是直接覆盖。要操作数少的话其实还可以加个判断看需不需要交换（即cnt == 0  的时候不用交换）。

其实要说的话和双指针差不多，不过我这个是直接算出来的双指针。
