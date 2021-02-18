# 995.K 连续位的最小翻转次数
[https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/](https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/) 
## 原题
在仅包含 `0` 和 `1` 的数组 `A` 中，一次 *`K` 位翻转* 包括选择一个长度为 `K` 的（连续）子数组，同时将子数组中的每个 `0` 更改为 `1` ，而每个 `1` 更改为 `0` 。

返回所需的 `K` 位翻转的最小次数，以便数组没有值为 `0` 的元素。如果不可能，返回 `-1` 。

 

 **示例 1：** 

```

输入：A = [0,1,0], K = 1
输出：2
解释：先翻转 A[0]，然后翻转 A[2]。

```
 **示例 2：** 

```

输入：A = [1,1,0], K = 2
输出：-1
解释：无论我们怎样翻转大小为 2 的子数组，我们都不能使数组变为 [1,1,1]。

```
 **示例 3：** 

```

输入：A = [0,0,0,1,0,1,1,0], K = 3
输出：3
解释：
翻转 A[0],A[1],A[2]: A变成 [1,1,1,1,0,1,1,0]
翻转 A[4],A[5],A[6]: A变成 [1,1,1,1,1,0,0,0]
翻转 A[5],A[6],A[7]: A变成 [1,1,1,1,1,1,1,1]

```
 

 **提示：** 
-  `1 <= A.length <= 30000` 
-  `1 <= K <= A.length` 
 
**标签**
`贪心算法` `Sliding Window` 


## 暴力模拟
```go
func flipA2B(nums []int, A, B int) {
	for A < B {
		nums[A] = nums[A] ^ 1
		A++
	}
}
func firstX(nums []int, st, x int) int {
	for st < len(nums) && nums[st] != x {
		st++
	}
	return st
}
func minKBitFlips(A []int, K int) int {
	cnt := 0
	for st := firstX(A, 0, 0); st+K-1 < len(A); st = firstX(A, st+1, 0) {
		flipA2B(A, st, st+K)
		cnt++
	}
	if st := firstX(A, 0, 0); st < len(A) {
		return -1
	}
	return cnt
}
```
>执行用时: 1168 ms
内存消耗: 7.1 MB

一千毫秒，基本就是踩着`TLE`的边了。思路很简单，从左往右开始找第一个为 `0` 的，然后就把以它为起点的开始翻转。统计翻转的次数即可。

因为对于位置 i 上的数而言，只有它之前的才能影响它，而它后面的数不论怎么反转都对它不影响。因此就从左往右翻转即可。

## 差分数组
```go
func minKBitFlips(A []int, K int) (ans int) {
	n := len(A)
	diff := make([]int, n+1)
	revCnt := 0
	for i, v := range A {
		revCnt ^= diff[i]
		if v == revCnt { // v^revCnt == 0
			if i+K > n {
				return -1
			}
			ans++
			revCnt ^= 1
			diff[i+K] ^= 1
		}
	}
	return
}
```
>执行用时: 120 ms
内存消耗: 6.9 MB

这是对于差分数组的应用，这里摘录一段别人的解释
>**差分数组的主要适用场景是频繁对原始数组的某个区间的元素进行增减。**
>比如说，我给你输入一个数组 `nums`，然后又要求给区间 `nums[2..6]` 全部加 1，再给 `nums[3..9]` 全部减 3，再给 `nums[0..4]` 全部加 2，再给...
>
>一通操作猛如虎，然后问你，最后 `nums` 数组的值是什么？
>
>常规的思路很容易，你让我给区间 `nums[i..j]` 加上 `val`，那我就一个 for 循环给它们都加上呗，还能咋样？这种思路的时间复杂度是 O(N)，由于这个场景下对 `nums` 的修改非常频繁，所以效率会很低下。
>
>这里就需要差分数组的技巧，类似前缀和技巧构造的 `prefix` 数组，我们先对 `nums` 数组构造一个 `diff` 差分数组，`diff[i]` **就是** `nums[i]` **和** `nums[i-1]` **之差**
>
>来自： https://labuladong.github.io/algo/算法思维系列/差分技巧.html

简单地说就是diff记录的是两个相邻的数之差。下面举个例子
### 例子
记 `nums` 为原始数组
构造方法如下
```go
diff := make([]int,len(nums))
diff[0] = nums[0]
for i := 1; i < len(nums); i++ {
    diff[i] = nums[i] - nums[i-1]
}
```
那么由差分数组还原原始数组的话也很简单
```go
res := make([]int, len(diff))
res[0] = diff[0]
for i := 1; i < len(diff); i++ {
    res[i] = res[i-1] + diff[i]
}
```

关键的到了，此时如果要对区间进行加减的话就如下操作

```go
/* 
    比如说此时的 nums = {8,2,6,3,1}
    那么此时的 diff = {8,-6,4,-3,-2}
    要让 res = {8,5,9,6,1} 即 区间 nums[1:4] 中的元素都加 3 
    那么就让 diff[1] += 3 diff[4]-=3 即可
    此时的 diff = {8,-3,4,-3,-5}
 */
```

由这个例子可以得到，我们要记录翻转的次数就要对一个区间进行 `+1` ，此时用差分数组能够减小时间复杂度。把 $O(K)$ 的时间复杂度简化为 $O(1)$ 。
