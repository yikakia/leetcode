# 31. 下一个排列
[https://leetcode-cn.com/problems/next-permutation/](https://leetcode-cn.com/problems/next-permutation/)
## 找到可以修改的位置和可以修改的对象
```go

func nextPermutation(nums []int) {
	hasChanged := false
	n := len(nums)
	if n == 0 {
		return
	}
	st := n - 2
	for st >= 0 {
		if nums[st] < nums[st+1] {
			hasChanged = true
			break
		}
		st--
	}
	if hasChanged {
		end := n - 1
		for end >= 0 {
			if nums[st] < nums[end] {
				nums[st], nums[end] = nums[end], nums[st]
				break
			}
			end--
		}
	}
	reverse(nums[st+1:])
}
func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
```
>执行用时：4 ms, 在所有 Go 提交中击败了69.02%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了46.90%的用户

简单地说就是从后往前，找到第一个满足升序的位置，即 `nums[i] < nums[i+1]` ,这个时候的位置 `i` 就是之后需要修改的地方。

然后我们就再从后往前，找到第一个满足 `i<j && nums[i] < nums[j]` 的位置，这个时候交换`nums[i]` `nums[j]`，并且将 `nums[i]` 后的数组翻转就可以了。

为什么可以这样呢？第一步好理解，就是找到尽量靠后的可以交换的位置。因为找到了就说明后面`nums[i+1:]`是非升序序列，没有更大的排列了。

第二步就有点不好理解了，简单来说就是为了让交换尽可能小，我们要找到尽可能靠后的 第一个 比 可以交换的位置 大的元素。为什么呢？还是为了找到下一个排列。简单来说就是我们找到了第一个升序的位置，这个位置就是我们可以尽可能小的修改的位置，但是这个位置上的数也应该尽可能小，使得之后生成的序列尽可能小。而我们知道，后面的部分是非升序序列，所以必然是越靠后的值越小。这样就变相找到了交换后这个位置上的最小的值。

交换完之后还没有完，因为我们知道`nums[i+1:]`是一个非升序序列，也就是说越靠前的越大，那么我们就把这后半部分翻转了，就可以得到后半部分的最小排列了。

整个逻辑流程就是这样了，主要是第二步不太好想到。