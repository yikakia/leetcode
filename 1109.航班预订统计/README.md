# 1109.航班预订统计
[https://leetcode-cn.com/problems/corporate-flight-bookings](https://leetcode-cn.com/problems/corporate-flight-bookings) 
## 原题
这里有  `n`  个航班，它们分别从 `1` 到 `n` 进行编号。

有一份航班预订表  `bookings` ，表中第  `i`  条预订记录  `bookings[i] = [first<sub>i</sub>, last<sub>i</sub>, seats<sub>i</sub>]`  意味着在从 `first<sub>i</sub>`  到 `last<sub>i</sub>` （ **包含** `first<sub>i</sub>` 和 `last<sub>i</sub>` ）的 **每个航班** 上预订了 `seats<sub>i</sub>`  个座位。

请你返回一个长度为 `n` 的数组  `answer` ，其中 `answer[i]` 是航班 `i` 上预订的座位总数。

 

 **示例 1：** 

```

输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]

```
 **示例 2：** 

```

输入：bookings = [[1,2,10],[2,2,15]], n = 2
输出：[10,25]
解释：
航班编号        1   2
预订记录 1 ：   10  10
预订记录 2 ：       15
总座位数：      10  25
因此，answer = [10,25]

```
 

 **提示：** 
-  `1 <= n <= 2 * 10^4` 
-  `1 <= bookings.length <= 2 * 10^4` 
-  `bookings[i].length == 3` 
-  `1 <= first<sub>i</sub> <= last<sub>i</sub> <= n` 
-  `1 <= seats<sub>i</sub> <= 10^4` 
 
**标签**
`数组` `前缀和` 


## 差分数组
```go
func corpFlightBookings(bookings [][]int, n int) []int {
	answer := make([]int, n+2)
	for _, booking := range bookings {
		from, to, orders := booking[0], booking[1], booking[2]
		answer[from] += orders
		answer[to+1] -= orders
	}
	for i := 1; i <= n; i++ {
		answer[i] += answer[i-1]
	}
	return answer[1 : len(answer)-1]
}
```
>执行用时：120 ms
内存消耗：9.6 MB

简单的差分数组应用于区间和问题。