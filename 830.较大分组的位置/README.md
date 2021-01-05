# 830. 较大分组的位置
[https://leetcode-cn.com/problems/positions-of-large-groups/](https://leetcode-cn.com/problems/positions-of-large-groups/) 
## 原题
在一个由小写字母构成的字符串 `s` 中，包含由一些连续的相同字符所构成的分组。
例如，在字符串 `s = "abbxxxxzyy"` 中，就含有 `"a"`, `"bb"`, `"xxxx"`, `"z"` 和 `"yy"` 这样的一些分组。
分组可以用区间 `[start, end]` 表示，其中 `start` 和 `end` 分别表示该分组的起始和终止位置的下标。上例中的 `"xxxx"` 分组用区间表示为 `[3,6]` 。
我们称所有包含大于或等于三个连续字符的分组为 **较大分组**  。
找到每一个 **较大分组**  的区间，**按起始位置下标递增顺序排序后** ，返回结果。
 
**示例 1：** 
```
输入：s = "abbxxxxzzy"
输出：[[3,6]]
解释："xxxx" 是一个起始于 3 且终止于 6 的较大分组。
```
**示例 2：** 
```
输入：s = "abc"
输出：[]
解释："a","b" 和 "c" 均不是符合要求的较大分组。
```
**示例 3：** 
```
输入：s = "abcdddeeeeaabbbcd"
输出：[[3,5],[6,9],[12,14]]
解释：较大分组为 "ddd", "eeee" 和 "bbb"
```
**示例 4：** 
```
输入：s = "aba"
输出：[]
```
 
**提示：** 
- `1 <= s.length <= 1000`
- `s` 仅含小写英文字母


## 
```go

func largeGroupPositions(s string) [][]int {
	res := [][]int{}

	cur := []int{0, 0}
	curRune := '0'
	count := 1
	for i, ch := range s {
		if ch != curRune {
			if cur[1]-cur[0] >= 2 {
				res = append(res, cur)
			}
			curRune = ch
			cur = []int{0, 0}
			count = 1
			continue
		} else {
			count++
			if count == 3 {
				cur[0] = i - 2
				cur[1] = i
			} else if count > 3 {
				cur[1] = i
			}
		}
	}
	if cur[1]-cur[0] >= 2 {
		if len(res) == 0 {
			res = append(res, cur)
		} else if res[len(res)-1][0] != cur[0] {
			res = append(res, cur)
		}
	}

	return res
}
```
>执行用时: 0 ms
内存消耗: 2.6 MB

这个题的思路就是记录区间，然后满足的就把它放进去。不过如果不想特殊判断最后一个是满足条件的情况的话，那么就
- 要么在后面加一个非小写字母的字符 ` s+="0" `
- 要么就多循环一次 `for i := 0;i <= len(s); i++`
    在循环里面进行判断 `if i==len(s)||i-start>=2`

总之就是要多判断一次尾节点，不过可以通过循环来让尾节点的特殊一般化即

```go
func largeGroupPositions(s string) [][]int {
    start := 0
    var r [][]int
    for i := 0; i <= len(s); i++ {
        if i == len(s) || s[start] != s[i] {
            if i - start >= 3 {
                r = append(r, []int{start, i-1})
            }
            start = i
        }
    }
    return r
}
```
