# 421.数组中两个数的最大异或值
[https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array](https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array) 
## 原题
给你一个整数数组 `nums` ，返回 `nums[i] XOR nums[j]` 的最大运算结果，其中 `0 ≤ i ≤ j < n` 。

 **进阶：** 你可以在 `O(n)` 的时间解决这个问题吗？

 
 **示例 1：** 

```

输入：nums = [3,10,5,25,2,8]
输出：28
解释：最大运算结果是 5 XOR 25 = 28.
```
 **示例 2：** 

```

输入：nums = [0]
输出：0

```
 **示例 3：** 

```

输入：nums = [2,4]
输出：6

```
 **示例 4：** 

```

输入：nums = [8,10,2]
输出：10

```
 **示例 5：** 

```

输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
输出：127

```
 

 **提示：** 
-  `1 <= nums.length <= 2 * 10^4` 
-  `0 <= nums[i] <= 2^31 - 1` 
 
**标签**
`位运算` `字典树` 


## 
```go
func findMaximumXOR(nums []int) int {
	t := trie{}
	for _, num := range nums {
		t.Add(num)
	}
	res := 0
	for _, num := range nums {
		if t := t.MaxXor(num); t > res {
			res = t
		}
	}
	return res
}

type trie struct {
	zero, one *trie
}

const highBin = 30

func (t *trie) Add(num int) {
	cur := t
	for i := highBin; i >= 0; i-- {
		if bit := num >> i & 1; bit == 0 {
			if cur.zero == nil {
				cur.zero = &trie{}
			}
			cur = cur.zero
		} else {
			if cur.one == nil {
				cur.one = &trie{}
			}
			cur = cur.one
		}
	}
}

func (t *trie) MaxXor(num int) (res int) {
	cur := t
	for i := highBin; i >= 0; i-- {
		if bit := num >> i & 1; bit == 0 {
			if cur.one != nil {
				cur = cur.one
				res = res*2 + 1
			} else {
				cur = cur.zero
				res *= 2
			}
		} else {
			if cur.zero != nil {
				cur = cur.zero
				res = res*2 + 1
			} else {
				cur = cur.one
				res *= 2
			}
		}
	}
	return
}
```
>执行用时: 48 ms
内存消耗: 7.7 MB

本质是用字典树记录出现过的数的二进制位，而求异或就是求这个数的当前位为0的时候是否存在一个数这个位是1.