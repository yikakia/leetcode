# 1722.执行交换操作后的最小汉明距离
[https://leetcode-cn.com/problems/minimize-hamming-distance-after-swap-operations/](https://leetcode-cn.com/problems/minimize-hamming-distance-after-swap-operations/) 
## 原题
给你两个整数数组 `source` 和 `target` ，长度都是 `n` 。还有一个数组 `allowedSwaps` ，其中每个 `allowedSwaps[i] = [a<sub>i</sub>, b<sub>i</sub>]` 表示你可以交换数组 `source` 中下标为 `a<sub>i</sub>` 和 `b<sub>i</sub>`（**下标从 0 开始** ）的两个元素。注意，你可以按 **任意**  顺序 **多次**  交换一对特定下标指向的元素。

相同长度的两个数组 `source` 和 `target` 间的 **汉明距离**  是元素不同的下标数量。形式上，其值等于满足 `source[i] != target[i]` （**下标从 0 开始** ）的下标 `i`（`0 <= i <= n-1`）的数量。

在对数组 `source` 执行 **任意**  数量的交换操作后，返回 `source` 和 `target` 间的 **最小汉明距离**  。

 

**示例 1：** 

```
输入：source = [1,2,3,4], target = [2,1,4,5], allowedSwaps = [[0,1],[2,3]]
输出：1
解释：source 可以按下述方式转换：
- 交换下标 0 和 1 指向的元素：source = [2,1,3,4]
- 交换下标 2 和 3 指向的元素：source = [2,1,4,3]
source 和 target 间的汉明距离是 1 ，二者有 1 处元素不同，在下标 3 。

```
**示例 2：** 

```
输入：source = [1,2,3,4], target = [1,3,2,4], allowedSwaps = []
输出：2
解释：不能对 source 执行交换操作。
source 和 target 间的汉明距离是 2 ，二者有 2 处元素不同，在下标 1 和下标 2 。
```
**示例 3：** 

```
输入：source = [5,1,2,4,3], target = [1,5,4,2,3], allowedSwaps = [[0,4],[4,2],[1,3],[1,4]]
输出：0

```
 

**提示：** 
- `n == source.length == target.length`
- `1 <= n <= 10^5`
- `1 <= source[i], target[i] <= 10^5`
- `0 <= allowedSwaps.length <= 10^5`
- `allowedSwaps[i].length == 2`
- `0 <= a<sub>i</sub>, b<sub>i</sub> <= n - 1`
- `a<sub>i</sub> != b<sub>i</sub>`
 
**标签**
`贪心算法` `深度优先搜索` `并查集` 


## 并查集
```go

func calHan(source []int, target []int) (count int) {
	for i := range source {
		if source[i] != target[i] {
			count++
		}
	}
	return
}
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	if len(allowedSwaps) == 0 {
		return calHan(source, target)
	}

	fa := make([]int, len(source))
	for i := range fa {
		fa[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for i := range allowedSwaps {
		a, b := allowedSwaps[i][0], allowedSwaps[i][1]
		fa[find(a)] = find(b)
	}

	counts := 0
	s := make(map[int]map[int]int)
	for i := range source {
		fa := find(i)
		if s[fa] == nil {
			s[fa] = make(map[int]int)
		}
		s[fa][source[i]]++
	}
	for i := range target {
		fa := find(i)
		if s[fa][target[i]] > 0 {
			counts++
			s[fa][target[i]]--
		}
	}

	return len(source) - counts
}
```
>执行用时: 264 ms
内存消耗: 31.5 MB

简单来说，什么情况下能够交换？当然是在`allowedSwaps`中的一对元素啊。然后这一对`a b`可以交换，那一对`b c`也可以交换，那么`a c`自然也是可以交换的。

于是思路就转变成了先通过并查集统计这个下标`index`属于哪个集合，然后统计这个集合中的各个字母出现的次数（因为到时候替换的时候是看这个字母，能够`cover`掉多少就`cover`掉多少）。因此可以构建 `map[int][26]int`的哈希表，但是我不知道怎么的出错了，于是改成了用`map[int]map[int]int`来统计。

最后计算的时候就是在当的这个字符还能够替换的时候更新可以替换的数目即可。当然也可以直接定义一个变量初始值为`len(source)`，然后更新的时候就是自减，这也是完全`ok`的。