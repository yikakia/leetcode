# 1122. 数组的相对排序
[https://leetcode-cn.com/problems/relative-sort-array/description/](https://leetcode-cn.com/problems/relative-sort-array/description/) 
## 自定义排序
```go
func relativeSortArray(arr1 []int, arr2 []int) []int {
	priority := make(map[int]int, len(arr2))
	for i, v := range arr2 {
		priority[v] = i
	}

	n := len(arr1)
	l, r := 0, n-1
	quicksort(arr1, l, r, priority)
	return arr1
}

// 返回 a 是不是小于 b
func less(a, b int, priority map[int]int) bool {
	p1, ok1 := priority[a]
	p2, ok2 := priority[b]
	if ok1 && ok2 {
		return p1 < p2
	} else if ok1 {
		return true
	} else if ok2 {
		return false
	}
	return a < b
}
func equal(a, b int, priority map[int]int) bool {
	p1, ok1 := priority[a]
	p2, ok2 := priority[b]
	if !ok1 && !ok2 {
		return a == b
	}
	return p1 == p2
}

// 返回 a 比 b 大
func larger(a, b int, priority map[int]int) bool {
	p1, ok1 := priority[a]
	p2, ok2 := priority[b]
	if ok1 && ok2 {
		return p1 > p2
	} else if ok1 {
		return false
	} else if ok2 {
		return true
	}
	return a > b
}
func quicksort(arr []int, l, r int, priority map[int]int) {
	if r-l < 1 {
		return
	}
	pivot := arr[l]
	start := l
	end := r
	l++
	for l < r {
		for (less(arr[l], pivot, priority) || equal(arr[l], pivot, priority)) && l < r {
			l++
		}
		for (larger(arr[r], pivot, priority) || equal(arr[r], pivot, priority)) && l < r {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
		}
	}
	if larger(arr[start], arr[l], priority) {
		arr[start], arr[l] = arr[l], arr[start]
	}
	quicksort(arr, start, l-1, priority)
	quicksort(arr, r, end, priority)
}
```
>执行用时：4 ms, 在所有 Go 提交中击败了5.51%的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了30.89%的用户

自定义排序，用了快排来加速。

不过还是比不过把两部分分开计算，然后用调库排序的效率。