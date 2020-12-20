# 5631. 跳跃游戏 VI
[https://leetcode-cn.com/problems/jump-game-vi/](https://leetcode-cn.com/problems/jump-game-vi/) 
## 原题
给你一个下标从 **0**  开始的整数数组 `nums` 和一个整数 `k` 。
一开始你在下标 `0` 处。每一步，你最多可以往前跳 `k` 步，但你不能跳出数组的边界。也就是说，你可以从下标 `i` 跳到 `[i + 1， min(n - 1, i + k)]` **包含**  两个端点的任意位置。
你的目标是到达数组最后一个位置（下标为 `n - 1` ），你的 **得分**  为经过的所有数字之和。
请你返回你能得到的 **最大得分**  。
 
**示例 1：** 
```
输入：nums = [1,-1,-2,4,-7,3], k = 2
输出：7
解释：你可以选择子序列 [1,-1,4,3] （上面加粗的数字），和为 7 。
```
**示例 2：** 
```
输入：nums = [10,-5,-2,4,0,3], k = 3
输出：17
解释：你可以选择子序列 [10,4,3] （上面加粗数字），和为 17 。
```
**示例 3：** 
```
输入：nums = [1,-5,-20,4,-1,3,-6,-3], k = 2
输出：0
```
 
**提示：** 
-  `1 <= nums.length, k <= 10^5`
- `-10^4 <= nums[i] <= 10^4`


## 
```go

func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums))
	if k > len(nums) {
		k = len(nums)
	}
	dp[0] = nums[0]
	preMax := dp[0]
	for i := 1; i < len(nums); i++ {
		if i <= k {
			dp[i] = preMax + nums[i]
			if dp[i] > preMax {
				preMax = dp[i]
			}
		} else {
			if dp[i-k-1] == preMax {
				preMax = max(dp[i-k : i]...)
			}
			dp[i] = preMax + nums[i]
			if dp[i] > preMax {
				preMax = dp[i]
			}
		}
	}
	return dp[len(nums)-1]
}

func max(nums ...int) (res int) {
	res = nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return
}
func test() {
	type TestType struct {
		nums []int
		k    int
		want int
	}
	ts := []TestType{
		TestType{
			nums: []int{1, -1, -2, 4, -7, 3},
			k:    2,
			want: 7,
		},
		TestType{
			nums: []int{10, -5, -2, 4, 0, 3},
			k:    3,
			want: 17,
		},
		TestType{
			nums: []int{1, -5, -20, 4, -1, 3, -6, -3},
			k:    2,
			want: 0,
		},
	}
	for i, t := range ts {
		get := maxResult(t.nums, t.k)
		if t.want != get {
			// 填充输出格式
			fmt.Printf("i:%d %+v get:%v\n", i, t, get)
		}
	}
	fmt.Println("Finished Test!")

}

```
>执行用时: 100 ms
内存消耗: 9.1 MB

不容易，终于做出了周赛的第三题（虽然过了十二点才做出来），💪。

直接用动规超时了，于是记录当前区间的最大值，当出去的值也为最大值时就更新最大值。

不过用单调栈更好。因为按我的这个方法的话对于`1,1,1,1,1,1,1`这样的数据和直接动规没有区别。