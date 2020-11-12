# 剑指 Offer 61. 扑克牌中的顺子
[https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/](https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/)
## 排序后确定需要的个数
```go
func isStraight(nums []int) bool {
	n := len(nums)
	sort(nums)
	cnt0 := 0
	for _, v := range nums {
		if v == 0 {
			cnt0++
		} else {
			break
		}
	}
	need := 0
	for i := cnt0; i < n-1; i++ {
		if tmp := nums[i+1] - nums[i]; tmp > 1 {
			need += tmp - 1
		} else if tmp == 0 {
			return false
		}
	}
	if need > cnt0 {
		return false
	}
	return true
}
func sort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了54.55%的用户

看了下，我这个排序然后比较弱爆了。就是找最大和最小，如果相差小于 5 就可以了。同时还要排除相同的情况。