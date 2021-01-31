# 839.相似字符串组
[https://leetcode-cn.com/problems/similar-string-groups/](https://leetcode-cn.com/problems/similar-string-groups/) 
## 原题
如果交换字符串 `X` 中的两个不同位置的字母，使得它和字符串 `Y` 相等，那么称 `X` 和 `Y` 两个字符串相似。如果这两个字符串本身是相等的，那它们也是相似的。

例如，`"tars"` 和 `"rats"` 是相似的 (交换 `0` 与 `2` 的位置)； `"rats"` 和 `"arts"` 也是相似的，但是 `"star"` 不与 `"tars"`，`"rats"`，或 `"arts"` 相似。

总之，它们通过相似性形成了两个关联组：`{"tars", "rats", "arts"}` 和 `{"star"}`。注意，`"tars"` 和 `"arts"` 是在同一组中，即使它们并不相似。形式上，对每个组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。

给你一个字符串列表 `strs`。列表中的每个字符串都是 `strs` 中其它所有字符串的一个字母异位词。请问 `strs` 中有多少个相似字符串组？

 

**示例 1：** 

```

输入：strs = ["tars","rats","arts","star"]
输出：2

```
**示例 2：** 

```

输入：strs = ["omv","ovm"]
输出：1

```
 

**提示：** 
- `1 <= strs.length <= 100`
- `1 <= strs[i].length <= 1000`
- `sum(strs[i].length) <= 2 * 10^4`
- `strs[i]` 只包含小写字母。
- `strs` 中的所有单词都具有相同的长度，且是彼此的字母异位词。
 

**备注：** 

      字母异位词（anagram），一种把某个字符串的字母的位置（顺序）加以改换所形成的新词。

 
**标签**
`深度优先搜索` `并查集` `图` 


## 并查集
```go
func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := newUnionFind(n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if isSame(strs[i], strs[j]) {
				uf.union(i, j)
			}
		}
	}
	return uf.count
}

type unionFind struct {
	fa, rank []int
	count    int
}

func newUnionFind(n int) *unionFind {
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &unionFind{fa, rank, n}
}

func (uf *unionFind) find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.find(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.fa[fy] = fx
	uf.rank[fx] += uf.rank[fy]
	uf.count--
}

func isSame(a, b string) bool {
	if a == b {
		return true
	}
	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
		if count > 2 {
			return false
		}
	}
	return true
}
```
>执行用时: 4 ms
内存消耗: 4.2 MB

也是标准的并查集。不过判断有没有边的话需要判断这两个字符串是否相似。因为题目说了字符串的长度相等了，所以就判断字符串长度了。总的来说，要相似的话需要满足
1. 长度相等
2. 不同字符数要小于等于2

这个判断的条件还是比较简单的。

