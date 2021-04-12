# 179.最大数
[https://leetcode-cn.com/problems/largest-number](https://leetcode-cn.com/problems/largest-number) 
## 原题
给定一组非负整数 `nums` ，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

 **注意：** 输出结果可能非常大，所以你需要返回一个字符串而不是整数。

 

 **示例 1：** 

```

输入：nums = [10,2]
输出："210"
```
 **示例 2：** 

```

输入：nums = [3,30,34,5,9]
输出："9534330"

```
 **示例 3：** 

```

输入：nums = [1]
输出："1"

```
 **示例 4：** 

```

输入：nums = [10]
输出："10"

```
 

 **提示：** 
-  `1 <= nums.length <= 100` 
-  `0 <= nums[i] <= 10^9` 
 
**标签**
`排序` 


## 局部贪心
```go
func largestNumber(nums []int) string {
	res := ""
	l := list.New()
	for _, num := range nums {
		l.PushBack(num)
	}

	for l.Len() > 0 {
		var find *list.Element
		max := ""
		for e := l.Front(); e != nil; e = e.Next() {
			if find == nil {
				find = e
				max = strconv.Itoa(e.Value.(int))
				continue
			}
			if num := strconv.Itoa(e.Value.(int)); num+max > max+num {
				max = num
				find = e
			}
		}
		res += max
		l.Remove(find)
	}
	if strings.TrimLeft(res, "0") == "" {
		return "0"
	}
	return res
}
```
>执行用时: 16 ms
内存消耗: 3.1 MB

这里用的是局部贪心。找到当前组合后最大的元素，那么前面的那个自然在最后的结果里面应该在更加前面。为了方便删除我就用了链表来实现。


## 排序
```go
func largestNumber(nums []int) string {
	strNums := make([]string, 0, len(nums))
	for _, num := range nums {
		strNums = append(strNums, strconv.Itoa(num))
	}
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})
	if strNums[0] == "0" {
		return "0"
	}
	return strings.Join(strNums, "")
}
```
>执行用时: 4 ms
内存消耗: 2.4 MB

这里的整体思路和之前的是一样的，但是转化为了排序。其实我之前的写法也相当于是排序了，但是相当于是选择排序，时间复杂度是 $O(n^2)$而用库函数的话就是快排的$O(nlog_n)$了