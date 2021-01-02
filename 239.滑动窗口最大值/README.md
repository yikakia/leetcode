# 239. 滑动窗口最大值
[https://leetcode-cn.com/problems/sliding-window-maximum/](https://leetcode-cn.com/problems/sliding-window-maximum/) 
## 原题
给你一个整数数组 `nums`，有一个大小为 `k` 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 `k` 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。
 
**示例 1：** 
```
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
```
**示例 2：** 
```
输入：nums = [1], k = 1
输出：[1]
```
**示例 3：** 
```
输入：nums = [1,-1], k = 1
输出：[1,-1]
```
**示例 4：** 
```
输入：nums = [9,11], k = 2
输出：[11]
```
**示例 5：** 
```
输入：nums = [4,-2], k = 2
输出：[4]
```
 
<b>提示：</b>
- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `1 <= k <= nums.length`


## 单调队列
```go
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n-k+1)
	queue := make([]int, 1)
	queue[0] = math.MinInt32
	for i := 0; i < k; i++ {
		queue = adjQueue(queue, nums[i])
	}
	for i := k; i < n; i++ {
		res[i-k] = queue[0]
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		queue = adjQueue(queue, nums[i])
	}
	res[n-k] = queue[0]
	return res
}

func adjQueue(queue []int, x int) []int {
	n := len(queue)
	res := make([]int, 0)
	if n == 0 {
		res = append(res, x)
	}
	for i := n - 1; i >= 0; i-- {
		if x <= queue[i] {
			res = append(queue[:i+1], x)
			break
		} else {
			res = append(queue[:i], x)
		}
	}

	return res
}
```
>执行用时: 232 ms
内存消耗: 10 MB

很单纯的单调队列。不过还是卡了很久。可能是迟钝了。

另外附上别人的一个解法，这是我最开始想写的版本，但是没写出来
```go
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 {
		return []int{}
	}
	//用切片模拟一个双端队列
	queue := []int{}
	result := []int{}
	for i := range nums {
		for i > 0 && (len(queue) > 0) && nums[i] > queue[len(queue)-1] {
            //将比当前元素小的元素祭天
			queue = queue[:len(queue)-1]
		}
        //将当前元素放入queue中
		queue = append(queue, nums[i])
		if i >= k && nums[i-k] == queue[0] {
            //维护队列，保证其头元素为当前窗口最大值
			queue = queue[1:]
		}
		if i >= k-1 {
            //放入结果数组
			result = append(result, queue[0])
		}
	}
	return result
}
```

现在想来应该是和另外一道题搞混了。另外一个题要求一直维护一个定长的队列。
我就是在维护队列和放入结果数组的时候搞错了。放入结果数组和维护队列的 i 的起始值并不一样。