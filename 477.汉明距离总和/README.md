# 477.汉明距离总和
[https://leetcode-cn.com/problems/total-hamming-distance](https://leetcode-cn.com/problems/total-hamming-distance) 
## 原题
两个整数的 <a href="https://baike.baidu.com/item/%E6%B1%89%E6%98%8E%E8%B7%9D%E7%A6%BB/475174?fr=aladdin">汉明距离</a> 指的是这两个数字的二进制数对应位不同的数量。

计算一个数组中，任意两个数之间汉明距离的总和。

 **示例:** 

```

输入: 4, 14, 2

输出: 6

解释: 在二进制表示中，4表示为0100，14表示为1110，2表示为0010。（这样表示是为了体现后四位之间关系）
所以答案为：
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

```
 **注意:** 
- 数组中元素的范围为从 `0` 到 `10^9` 。
- 数组的长度不超过 `10^4` 。
 
**标签**
`位运算` 


## 暴力
```go
func totalHammingDistance(nums []int) int {
	res := 0
	for i := range nums[1:] {
		for j := i + 1; j < len(nums); j++ {
			res += bits.OnesCount(uint(nums[i] ^ nums[j]))
		}
	}
	return res
}
```
>执行用时: 644 ms
内存消耗: 6.4 MB

粗暴的$O(n^2)$

## 按位计算汉明距离的和
```go
func totalHammingDistance(nums []int) int {
	res := 0
	n := len(nums)
	for i := 0; i <= 30; i++ {
		oneCounts := 0
		for _, num := range nums {
			oneCounts += (num >> i) & 1
		}
		res += oneCounts * (n - oneCounts)
	}
	return res
}
```
>执行用时: 32 ms
内存消耗: 6.5 MB

简单来说就是统计每个位上的汉明距离，然后统计起来即可。