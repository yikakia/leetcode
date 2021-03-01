# 645.错误的集合
[https://leetcode-cn.com/problems/set-mismatch](https://leetcode-cn.com/problems/set-mismatch) 
## 原题
集合 `s` 包含从 `1` 到  `n`  的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 **丢失了一个数字** 并且 **有一个数字重复** 。

给定一个数组 `nums` 代表了集合 `S` 发生错误后的结果。

请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

 

 **示例 1：** 

```

输入：nums = [1,2,2,4]
输出：[2,3]

```
 **示例 2：** 

```

输入：nums = [1,1]
输出：[1,2]

```
 

 **提示：** 
-  `2 <= nums.length <= 10^4` 
-  `1 <= nums[i] <= 10^4` 
 
**标签**
`哈希表` `数学` 


## 用数值的正负代表是否出现过
```go
func findErrorNums(nums []int) []int {
	var repect int
	var disapper int
	for _, num := range nums {
		num = abs(num)
		if nums[num-1] < 0 {
			repect = num
			continue
		}
		nums[num-1] *= -1
	}

	for i, x := range nums {
		if x > 0 {
			disapper = i + 1
			break
		}
	}

	return []int{repect, disapper}
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
```
>执行用时: 28 ms
内存消耗: 6.4 MB

## 用位运算查找
```go
func findErrorNums(nums []int) []int {
	var ab int
	for i, num := range nums {
		ab ^= num ^ (i + 1)
	}
	mask := 1
	for mask&ab == 0 {
		mask <<= 1
	}
	var a, b int
	for i, num := range nums {
		if mask&num == 0 {
			a ^= num
		} else {
			b ^= num
		}
		if mask&(i+1) == 0 {
			a ^= i + 1
		} else {
			b ^= i + 1
		}
	}
	for _, num := range nums {
		if num == a {
			return []int{a, b}
		}
	}

	return []int{b, a}
}
```
>执行用时: 32 ms
内存消耗: 6.4 MB

这里涉及到了[另外一道题](../260.只出现一次的数字-III/README.md)的解法。可以先做了它再来做这道题。

从题目里面可以知道，它是先生成了 `[1...n]` **n** 个数，这里把它叫做 `nums1`，然后把其中一个数换成了`[1...n]` 里面的另外一个数，这里把它叫做 `nums2`。

输入进来的就是 `nums2` 。这里我们把 `nums1` 的每一个元素进行异或，然后再和 `nums2` 中的每一个元素进行异或，这样就得到了一个值来表示重复值 `a` 与缺失值 `b` 的异或 `ab = a^b`。

然后我们要找到一个办法来区分 `a` 和 `b`。这里就用到了[260.只出现一次的数字-III](../260.只出现一次的数字-III/README.md)中的解法，用最右边的`1`来表示这两个数的区别，即 `mask = (10000……)`。

于是再一次进行循环，按照 `nums1` 和 `nums2` 与 `mask` 相与是否为 `0` 来判断应该与 `a` 相与还是与 `b` 相与。

这里与哪个相与自己定即可。不过上下一定要相同。即 `mask & nums1 == 0 ` 的话 `a^=nums1` 那么 `mask & nums2 == 0` 的话，也应该 `a^=nums2`。

原理就是经过分类以后重复的那个数一共会重复出现三次(`nums1` 中出现一次，`nums2` 中出现两次)

而和重复的那个数一组的数会出现两次(`nums1` 中出现一次，`nums2` 中出现一次)。

同理，缺失的那个数一共会出现一次(`nums1` 中出现一次，`nums2` 中出现零次，因为缺失了)

而和重复的那个数一组的数会出现两次(`nums1` 中出现一次，`nums2` 中出现一次)

最后输出的时候要求先输出重复的数再输出缺失的数。这里就循环一次 `nums2` 看是否出现即可。