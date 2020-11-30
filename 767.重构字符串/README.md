# 767. 重构字符串
[https://leetcode-cn.com/problems/reorganize-string/](https://leetcode-cn.com/problems/reorganize-string/) 
## 优先排最大的和次大的
```go
func reorganizeString(S string) string {
	n := len(S)
	type pair struct {
		char  rune
		count int
	}
	occurTimesMap := make(map[rune]int)
	for _, v := range S {
		occurTimesMap[v]++
	}
	pairs := []pair{}
	for k, v := range occurTimesMap {
		if v > (n+1)/2 {
			return ""
		}
		pairs = append(pairs, pair{char: k, count: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})
	res := strings.Builder{}
	for len(pairs) != 0 {
		if len(pairs) > 1 {
			res.WriteRune(pairs[0].char)
			res.WriteRune(pairs[1].char)
			pairs[0].count--
			pairs[1].count--
			sort.Slice(pairs, func(i, j int) bool {
				return pairs[i].count > pairs[j].count
			})
			for len(pairs) > 0 && pairs[len(pairs)-1].count == 0 {
				pairs = pairs[0 : len(pairs)-1]
			}
		} else {
			res.WriteRune(pairs[0].char)
			pairs[0].count--
			if pairs[0].count == 0 {
				pairs = pairs[0:0]
			} else {
				return ""
			}
		}
	}
	return res.String()
}
```
>执行用时: 0 ms
内存消耗: 2.3 MB

简单地说就是先统计出现的次数，然后再排序。之后就是不断地塞进去次数最多的和次数次多的，并且同时减少它们的出现次数，然后再排一次序。最后就是如果出到最后一个字母了都还没有完全出去的话就返回空。

其实这个返回空的话可以直接先进行判断。为什么会最后一个字母还不能完全排序的情况？因为最后一个字母占据了一半还多的次数。因此提前判断最多的数是不是出现的次数大于$\frac{n+1}{2}$就可以了

## 优化的思路
从这个出去的角度分析，这个要动态地维护最大和次大，因此可以用堆排序来加速。这个题的数据量小所以每次重新排序的时间消耗可以忽略不计。但是如果次数多了就得用堆排序了。

简单地说就是每次最大和次大出堆，然后把最大-1和次大-1都重新入堆(大于0的时候)。直到出完为止（可以只出最大不出次大，但是这个时候最大的count应该为1不然就返回空）。