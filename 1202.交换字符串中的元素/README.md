# 1202.交换字符串中的元素
[https://leetcode-cn.com/problems/smallest-string-with-swaps/](https://leetcode-cn.com/problems/smallest-string-with-swaps/) 
## 原题
给你一个字符串 `s`，以及该字符串中的一些「索引对」数组 `pairs`，其中 `pairs[i] = [a, b]` 表示字符串中的两个索引（编号从 0 开始）。

你可以 **任意多次交换**  在 `pairs` 中任意一对索引处的字符。

返回在经过若干次交换后，`s` 可以变成的按字典序最小的字符串。

 

**示例 1:** 

```
输入：s = "dcab", pairs = [[0,3],[1,2]]
输出："bacd"
解释： 
交换 s[0] 和 s[3], s = "bcad"
交换 s[1] 和 s[2], s = "bacd"

```
**示例 2：** 

```
输入：s = "dcab", pairs = [[0,3],[1,2],[0,2]]
输出："abcd"
解释：
交换 s[0] 和 s[3], s = "bcad"
交换 s[0] 和 s[2], s = "acbd"
交换 s[1] 和 s[2], s = "abcd"
```
**示例 3：** 

```
输入：s = "cba", pairs = [[0,1],[1,2]]
输出："abc"
解释：
交换 s[0] 和 s[1], s = "bca"
交换 s[1] 和 s[2], s = "bac"
交换 s[0] 和 s[1], s = "abc"

```
 

**提示：** 
- `1 <= s.length <= 10^5`
- `0 <= pairs.length <= 10^5`
- `0 <= pairs[i][0], pairs[i][1] < s.length`
- `s` 中只含有小写英文字母
 
**标签**
`并查集` `数组` 


## 并查集，优化了合并
```go
func smallestStringWithSwaps(s string, pairs [][]int) string {
	if len(pairs) == 0 {
		return s
	}
	n := len(s)

	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(from, to int) {
		fFrom, fTo := find(from), find(to)
		if fFrom == fTo {
			return
		}
		if rank[fTo] < rank[fFrom] {
			fFrom, fTo = fTo, from
		}
		rank[fTo] += rank[from]
		fa[fFrom] = fTo
	}

	for _, pair := range pairs {
		merge(pair[0], pair[1])
	}

	set := make(map[int][]byte)
	for i := range s {
		f := find(i)
		set[f] = append(set[f], s[i])
	}

	for _, bytes := range set {
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
	}

	res := make([]byte, n)
	for i := range res {
		f := find(i)
		res[i] = set[f][0]
		set[f] = set[f][1:]
	}
	return string(res)
}
```
>执行用时: 112 ms
内存消耗: 18.7 MB


今天有事,于是就用官方的题解了。这几天的并查集太多了，做了三道了吧，其实大体思路都是一样的，就是根据 `pairs` 来合并 `fa`。

不过合并的时候最好还能优化一下，简单地说就是把小的集合合并到大的集合里面（`rank[x]`就是记录指向的`x`的元素个数）。这样在合并的时候可以减小时间复杂度，代价是需要额外的$O(n)$的空间来储存。

另外就是并查集的最后一个字，集的定义一定要清楚。这个是根据题意来的。一般来说在进行了 `union` 操作之后才开始处理对应的集合。一般来讲都是 `map[fa[find(x)]]` 新增一个元素，即应该是 `map[int]slice` 这样的类型。

当然也可以是`map`嵌套，或者说二维`map`。不过 `Go` 里面做`map`的嵌套的时候要注意第二维的 `map` 是否为空。即如下，要先有一个判空的操作，再进行添加。因为这个和二维的数组一样，声明的指向的是`nil`，必须要初始化后才可以进行操作。
```go
sets := make(map[int]map[int]int)
for i := range fa {
    f := find(i)
    if set[f] == nil {
        set[f] = make(map[int]int)
    }
    // add elements in set
}
```

并查集大概就是这样了。这个 `rank` 是之前看过优化，但是没有实际写过，这次看了样例倒是清楚了，感觉这个优化倒是比想象中的简单一点。
