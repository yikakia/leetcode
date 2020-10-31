# 剑指 Offer 40. 最小的k个数
[https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

## 先排序，然后返回前 k 个数
```go

func getLeastNumbers(arr []int, k int) []int {
	fastsort(arr)
	return arr[:k]
}

func fastsort(nums []int) {
	if len(nums) < 2 {
		return
	}
	pivot := nums[0]
	fast, slow := len(nums)-1, 1
	for slow < fast {
		for nums[slow] <= pivot && slow < fast {
			slow++
		}
		for nums[fast] >= pivot && slow < fast {
			fast--
		}
		if slow < fast {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	if nums[0] > nums[slow] {
		nums[0], nums[slow] = nums[slow], nums[0]
	}
	fastsort(nums[:slow])
	fastsort(nums[fast:])
	return
}
```
>执行用时：36 ms, 在所有 Go 提交中击败了56.62%的用户
内存消耗：6.5 MB, 在所有 Go 提交中击败了45.33%的用户

简单地说 对于 Top N 问题，解决方法是维护一个大根堆，或者排序后输出前 k 个。排序一般来说更简单。

因为快排的特性，所以刚好可以适应并发。所以也写了一个并发类型的快排,希望能够快一点
```go
func getLeastNumbers(arr []int, k int) []int {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go fastsort(arr, &wg)
	wg.Wait()
	return arr[:k]
}

func fastsort(nums []int, wg *sync.WaitGroup) {
	if len(nums) < 2 {
		wg.Done()
		return
	}
	pivot := nums[0]
	fast, slow := len(nums)-1, 1
	for slow < fast {
		for nums[slow] <= pivot && slow < fast {
			slow++
		}
		for nums[fast] >= pivot && slow < fast {
			fast--
		}
		if slow < fast {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	if nums[0] > nums[slow] {
		nums[0], nums[slow] = nums[slow], nums[0]
	}
	wg.Add(1)
	go fastsort(nums[:slow], wg)
	wg.Add(1)
	go fastsort(nums[fast:], wg)

	wg.Done()
}

```
>执行用时：
80 ms, 在所有 Go 提交中击败了16.65%的用户
内存消耗：8.1 MB, 在所有 Go 提交中击败了5.08%的用户

实际上反而更加耗时间和耗内存了。想了想应该是数据量还不够大，如果较小的情况下要切换协程本来就需要很多时间。所以应该有个判断，当剩下的数量比较短的时候才并发，不然就单线程。

```go
func getLeastNumbers(arr []int, k int) []int {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go fastsort(arr, &wg)
	wg.Wait()
	return arr[:k]
}

func fastsort(nums []int, wg *sync.WaitGroup) {
	if len(nums) < 2 {
		wg.Done()
		return
	}
	pivot := nums[0]
	fast, slow := len(nums)-1, 1
	for slow < fast {
		for nums[slow] <= pivot && slow < fast {
			slow++
		}
		for nums[fast] >= pivot && slow < fast {
			fast--
		}
		if slow < fast {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	if nums[0] > nums[slow] {
		nums[0], nums[slow] = nums[slow], nums[0]
	}
	if len(nums) > 10 {
		wg.Add(1)
		go fastsort(nums[:slow], wg)
		wg.Add(1)
		go fastsort(nums[fast:], wg)
	} else {
		wg.Add(1)
		fastsort(nums[:slow], wg)
		wg.Add(1)
		fastsort(nums[fast:], wg)
	}

	wg.Done()
}
```
>执行用时：56 ms, 在所有 Go 提交中击败了18.92%的用户
内存消耗：7.1 MB, 在所有 Go 提交中击败了5.08%的用户

好像效果不错，耗时砍半了。把`if len(nums) > 10`修改为`if len(nums) > 30`后

>执行用时：36 ms, 在所有 Go 提交中击败了56.62%的用户
内存消耗：6.9 MB, 在所有 Go 提交中击败了5.08%的用户

嗯 时间和单线程差不多了 修改为 `50` 后

>执行用时：44 ms, 在所有 Go 提交中击败了38.08%的用户
内存消耗：6.6 MB, 在所有 Go 提交中击败了11.95%的用户

可以看出这个优化的效果不是越大越好，而是尽可能贴近数据的时候效果才足够明显。而内存占用越来越少也是可以理解的，毕竟协程本来就要占一定的空间。总的来说就是并不是并发就一定比单线程来得好，还是要看具体的应用上怎么样。

## TODO: 堆排序
