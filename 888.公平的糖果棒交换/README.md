# 888.公平的糖果棒交换
[https://leetcode-cn.com/problems/fair-candy-swap/](https://leetcode-cn.com/problems/fair-candy-swap/) 
## 原题
爱丽丝和鲍勃有不同大小的糖果棒：`A[i]` 是爱丽丝拥有的第 `i` 根糖果棒的大小，`B[j]` 是鲍勃拥有的第 `j` 根糖果棒的大小。

因为他们是朋友，所以他们想交换一根糖果棒，这样交换后，他们都有相同的糖果总量。*（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）*

返回一个整数数组 `ans`，其中 `ans[0]` 是爱丽丝必须交换的糖果棒的大小，`ans[1]` 是 Bob 必须交换的糖果棒的大小。

如果有多个答案，你可以返回其中任何一个。保证答案存在。

 

**示例 1：** 

```

输入：A = [1,1], B = [2,2]
输出：[1,2]

```
**示例 2：** 

```

输入：A = [1,2], B = [2,3]
输出：[1,2]

```
**示例 3：** 

```

输入：A = [2], B = [1,3]
输出：[2,3]

```
**示例 4：** 

```

输入：A = [1,2,5], B = [2,4]
输出：[5,4]

```
 

**提示：** 
- `1 <= A.length <= 10000`
- `1 <= B.length <= 10000`
- `1 <= A[i] <= 100000`
- `1 <= B[i] <= 100000`
- 保证爱丽丝与鲍勃的糖果总量不同。
- 答案肯定存在。
 
**标签**
`数组` 


## 排序后查找
```go
func sum(nums ...int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}

func fairCandySwap(A []int, B []int) []int {
	sort.Ints(A)
	sort.Ints(B)
	sa, sb := sum(A...), sum(B...)
	changed := false
	if sa > sb {
		sa, sb = sb, sa
		A, B = B, A
		changed = true
	}
	delta := sb - sa
	for _, num := range A {
		want := num + delta/2
		index := sort.Search(len(B), func(i int) bool { return B[i] >= want })
		if index < len(B) && B[index] == want {
			if changed {
				return []int{want, num}
			}
			return []int{num, want}
		}
	}
	return []int{}
}
```
>执行用时: 112 ms
内存消耗: 6.9 MB

想了想，其实不用交换，没有这个必要。
```go
func sum(nums ...int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}

func fairCandySwap(A []int, B []int) []int {
	sort.Ints(B)
	sa, sb := sum(A...), sum(B...)
	delta := sb - sa
	for _, num := range A {
		want := num + delta/2
		index := sort.Search(len(B), func(i int) bool { return B[i] >= want })
		if index < len(B) && B[index] == want {
			return []int{num, want}
		}
	}
	return []int{}
}
```
>执行用时: 80 ms
内存消耗: 6.9 MB

不用交换的话连 `A` 的排序都省下来了。还不错。

## 储存结果
```go
func sum(nums ...int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}

func fairCandySwap(A []int, B []int) []int {
	sa, sb := sum(A...), sum(B...)

	ma, mb := map[int]bool{}, map[int]bool{}
	for _, v := range A {
		ma[v] = true
	}
	for _, v := range B {
		mb[v] = true
	}
	delta := (sb - sa) / 2
	for _, a := range A {
		if _, ok := mb[a+delta]; ok {
			return []int{a, a + delta}
		}
	}
	return nil
}
```
>执行用时: 88 ms
内存消耗: 7.2 MB

空间换时间的操作。不过好像效果并不明显？时间比起二分查找而言消耗更多了。应该是取 `map` 的操作多耗了点时间吧。
