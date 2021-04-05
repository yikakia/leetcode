# 781.森林中的兔子
[https://leetcode-cn.com/problems/rabbits-in-forest](https://leetcode-cn.com/problems/rabbits-in-forest) 
## 原题
森林中，每个兔子都有颜色。其中一些兔子（可能是全部）告诉你还有多少其他的兔子和自己有相同的颜色。我们将这些回答放在 `answers` 数组里。

返回森林中兔子的最少数量。

```

示例:
输入: answers = [1, 1, 2]
输出: 5
解释:
两只回答了 "1" 的兔子可能有相同的颜色，设为红色。
之后回答了 "2" 的兔子不会是红色，否则他们的回答会相互矛盾。
设回答了 "2" 的兔子为蓝色。
此外，森林中还应有另外 2 只蓝色兔子的回答没有包含在数组中。
因此森林中兔子的最少数量是 5: 3 只回答的和 2 只没有回答的。

输入: answers = [10, 10, 10]
输出: 11

输入: answers = []
输出: 0

```
 **说明:** 
-  `answers` 的长度最大为 `1000` 。
-  `answers[i]` 是在 `[0, 999]` 范围内的整数。
 
**标签**
`哈希表` `数学` 


## 记录之后向上取整
```go
func numRabbits(answers []int) int {
	res := 0
	m := map[int]int{}
	for _, num := range answers {
		m[num]++
	}
	for k, v := range m {
		if k+1 >= v {
			res += k + 1
		} else {
			nogroups := v % (k + 1)
			if nogroups > 0 {
				res += v - nogroups + k + 1
			} else if nogroups == 0 {
				res += v
			}
		}
	}
	return res
}
```
>执行用时: 0 ms
内存消耗: 2.8 MB

简单地说就是对一有几个的说法分成不同的小组，然后计算它们的向上取整的结果即可。