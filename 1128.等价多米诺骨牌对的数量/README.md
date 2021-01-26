# 1128.等价多米诺骨牌对的数量
[https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/](https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/) 
## 原题
给你一个由一些多米诺骨牌组成的列表 `dominoes`。

如果其中某一张多米诺骨牌可以通过旋转 `0` 度或 `180` 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。

形式上，`dominoes[i] = [a, b]` 和 `dominoes[j] = [c, d]` 等价的前提是 `a==c` 且 `b==d`，或是 `a==d` 且 `b==c`。

在 `0 <= i < j < dominoes.length` 的前提下，找出满足 `dominoes[i]` 和 `dominoes[j]` 等价的骨牌对 `(i, j)` 的数量。

 

**示例：** 

```
输入：dominoes = [[1,2],[2,1],[3,4],[5,6]]
输出：1

```
 

**提示：** 
- `1 <= dominoes.length <= 40000`
- `1 <= dominoes[i][j] <= 9`
 
**标签**
`数组` 


## 暴力模拟
```go
func numEquivDominoPairs(dominoes [][]int) int {
	for i := range dominoes {
		if dominoes[i][0] > dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}
	}
	sort.Slice(dominoes, func(i, j int) bool {
		if dominoes[i][0] < dominoes[j][0] {
			return true
		} else if dominoes[i][0] == dominoes[j][0] && dominoes[i][1] < dominoes[j][1] {
			return true
		}
		return false
	})

	res := 0
	for i := range dominoes {
		for j := i + 1; j < len(dominoes); j++ {
			if dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1] {
				res++
			} else {
				break
			}
		}
	}
	return res
}
```
>执行用时: 100 ms
内存消耗: 8 MB

先调整大小，然后排序，这样在找的时候遇到了不符合的就跳过到后面。

## 标识出现次数
```go
func numEquivDominoPairs(dominoes [][]int) int {
	var m [10][10]int
	res := 0
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		res += m[d[0]][d[1]]
		m[d[0]][d[1]]++
	}
	return res
}
```
>执行用时: 32 ms
内存消耗: 7.6 MB


注意到题目的取值范围是`1 <= dominoes[i][j] <= 9`，那么就可以直接创建`map`或者是固定范围的数组来记录每一组牌的出现次数`count`，最后计算$\frac{count*(count-1)}{2}$即可。当然也可以每次出现的时候就加上当前的出现次数。从 0 开始，因为要计算的是一对，所以要减去自己出现的那一次。