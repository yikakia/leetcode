# 561.数组拆分 I
[https://leetcode-cn.com/problems/array-partition-i/](https://leetcode-cn.com/problems/array-partition-i/) 
## 原题
给定长度为  `2n`  的整数数组 `nums` ，你的任务是将这些数分成  `n`  对, 例如 `(a<sub>1</sub>, b<sub>1</sub>), (a<sub>2</sub>, b<sub>2</sub>), ..., (a<sub>n</sub>, b<sub>n</sub>)` ，使得从 `1` 到  `n` 的 `min(a<sub>i</sub>, b<sub>i</sub>)` 总和最大。

返回该 **最大总和** 。

 

 **示例 1：** 

```

输入：nums = [1,4,3,2]
输出：4
解释：所有可能的分法（忽略元素顺序）为：
1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
所以最大总和为 4
```
 **示例 2：** 

```

输入：nums = [6,2,6,5,1,2]
输出：9
解释：最优的分法为 (2, 1), (2, 5), (6, 6). min(2, 1) + min(2, 5) + min(6, 6) = 1 + 2 + 6 = 9

```
 

 **提示：** 
-  `1 <= n <= 10^4` 
-  `nums.length == 2 * n` 
-  `-10^4 <= nums[i] <= 10^4` 
 
**标签**
`数组` 


## 先排序，后取值
```go
func arrayPairSum(nums []int) (res int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return
}
```
>执行用时: 80 ms
内存消耗: 6.9 MB

题目要求取两个数来组合，每次都是选到的较小的那一个，那么最大的数一定不可能取到，而最小的数一定会被取到。

最后的结果总和要最大，利用之前的推论从 `2` 个的情况开始讨论，要让总和最大，那么就让它们从小开始变化，取一对中小的那一个就可以了。

那么 `4` 个的情况呢？还是排序，最大的和最小的和哪个组合都不影响，那么影响结果的就是中间的两个了。因为最小的和多大的组合取得的结果都是最小的那个，所以为了让结果最大，就应该让剩下的两个中的小的那一个和最小的组合，这样就能够让结果整个最大化。

继续推广就可以发现，只要先排序，然后隔值取值即可。