# 220.存在重复元素 III
[https://leetcode-cn.com/problems/contains-duplicate-iii](https://leetcode-cn.com/problems/contains-duplicate-iii) 
## 原题
给你一个整数数组 `nums` 和两个整数  `k` 和 `t` 。请你判断是否存在两个下标 `i` 和 `j` ，使得  `abs(nums[i] - nums[j]) <= t` ，同时又满足 `abs(i - j) <= k` 。

如果存在则返回 `true` ，不存在返回 `false` 。

 

 **示例 1：** 

```

输入：nums = [1,2,3,1], k = 3, t = 0
输出：true
```
 **示例 2：** 

```

输入：nums = [1,0,1,1], k = 1, t = 2
输出：true
```
 **示例 3：** 

```

输入：nums = [1,5,9,1,5,9], k = 2, t = 3
输出：false
```
 

 **提示：** 
-  `0 <= nums.length <= 2 * 10^4` 
-  `-2^31 <= nums[i] <= 2^31 - 1` 
-  `0 <= k <= 10^4` 
-  `0 <= t <= 2^31 - 1` 
 
**标签**
`排序` `Ordered Map` 


## 暴力模拟
```go
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		l, r := i-k, i+k
		if l < 0 {
			l = 0
		}
		if r > n {
			r = n
		}
		for j := l; j < r; j++ {
			if i != j && abs(nums[i]-nums[j]) <= t {
				return true
			}

		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
>执行用时: 740 ms
内存消耗: 5.9 MB

就是暴力模拟而已。

## 桶
```go
func getID(x, w int) int {
    if x >= 0 {
        return x / w
    }
    return (x+1)/w - 1
}

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
    mp := map[int]int{}
    for i, x := range nums {
        id := getID(x, t+1)
        if _, has := mp[id]; has {
            return true
        }
        if y, has := mp[id-1]; has && abs(x-y) <= t {
            return true
        }
        if y, has := mp[id+1]; has && abs(x-y) <= t {
            return true
        }
        mp[id] = x
        if i >= k {
            delete(mp, getID(nums[i-k], t+1))
        }
    }
    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/contains-duplicate-iii/solution/cun-zai-zhong-fu-yuan-su-iii-by-leetcode-bbkt/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```

这里是用的官方题解的版本。因为求的是存在，所以找到一组即可。在做的时候就是开多个大小为 t 的桶，如果两个元素在同一个桶里，说明两个元素的距离肯定小于 t 。不然的话就检查相邻的桶，如果绝对值之差小于 t 说明也满足条件。

当然了，题目还有对于下标的控制，这个时候就要维护一个滑动窗口就行。不断移除 i-k 的桶，即可保证向前的距离不大于 k 。而匹配的元素在后面的话也会被后面的元素匹配到。