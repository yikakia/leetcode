# 剑指 Offer 45. 把数组排成最小的数
[https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/](https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/)

## 自定义排序
```go
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})
	res := strings.Builder{}
	for i := 0; i < len(nums); i++ {
		res.WriteString(fmt.Sprint(nums[i]))
	}
	return res.String()
}
func compare(a, b int) bool {
	str1 := fmt.Sprintf("%d%d", a, b)
	str2 := fmt.Sprintf("%d%d", b, a)
	if str1 < str2 {
		return true
	}
	return false
}
```

>执行用时：8 ms, 在所有 Go 提交中击败了29.96%的用户
内存消耗：2.7 MB, 在所有 Go 提交中击败了52.34%的用户

简单地说就是利用字典序来排序。当`a+b < b+a`时，就认为 a 更小。这样不断地比较排序，最后按照顺序生成一个字符串就是了。
