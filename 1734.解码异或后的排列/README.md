# 1734.解码异或后的排列
[https://leetcode-cn.com/problems/decode-xored-permutation](https://leetcode-cn.com/problems/decode-xored-permutation) 
## 原题
给你一个整数数组  `perm`  ，它是前  `n`  个正整数的排列，且  `n`  是个 **奇数**  。

它被加密成另一个长度为 `n - 1`  的整数数组  `encoded`  ，满足  `encoded[i] = perm[i] XOR perm[i + 1]`  。比方说，如果  `perm = [1,3,2]`  ，那么  `encoded = [2,1]`  。

给你  `encoded`  数组，请你返回原始数组  `perm`  。题目保证答案存在且唯一。

 

 **示例 1：** 

```
输入：encoded = [3,1]
输出：[1,2,3]
解释：如果 perm = [1,2,3] ，那么 encoded = [1 XOR 2,2 XOR 3] = [3,1]

```
 **示例 2：** 

```
输入：encoded = [6,5,4,6]
输出：[2,4,1,5,3]

```
 

 **提示：** 
-  `3 <= n < 10^5` 
-  `n`  是奇数。
-  `encoded.length == n - 1` 
 
**标签**
`位运算` 


## 
```go
func decode(encoded []int) []int {
	n := len(encoded)

	totalXor := 0

	for i := 1; i <= n+1; i++ {
		totalXor ^= i
	}
	mask := 0
	for i := 1; i < n; i += 2 {
		mask ^= encoded[i]
	}
	res := make([]int, n+1)
	res[0] = totalXor ^ mask

	for i := 1; i < n+1; i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}

	return res
}
```
>执行用时: 188 ms
内存消耗: 9.8 MB

题目有个条件是 `decode` 的数组是 `1~n` 的全排列，这个时候就说明可以先处理得到所有的数字异或后的值 `totalXor` 。

然后就是怎么得到一个 `decode` 中的字了。只要知道其中的一个结果就可以得到其他的结果了。我们用取偶数位 `encode` 的异或后的结果，这个时候组合得到的就是 `decode` 的下标 `[1:n]` 异或的值。然后与 `totalXor` 相异或就能得到`decode` 的第一个值了。这个时候就和 [1720.解码异或后的数组](../1720.解码异或后的数组/README.md) 一样了。
