# 1. 两数之和
[https://leetcode-cn.com/problems/two-sum/](https://leetcode-cn.com/problems/two-sum/) 
## 哈希表
```go
func twoSum(nums []int, target int) []int {
	res := []int{0, 0}
	want := make(map[int]int, 0)
	for i, v := range nums {
		_, exist := want[v]
		if exist {
			res[0] = want[v]
			res[1] = i
			break
		} else {
			want[target-v] = i
		}
	}
	return res
}
```
>执行用时: 92 ms
内存消耗: 9 MB

好像现在是样例给得特别卡人，所以不好过了。简单来说下思路吧，简单的说就是开一个map来储存未来需要什么数。然后每次就判断现在这个数是不是被需要的，如果被需要就退出循环返回下标，不然就把当前的的数在未来需要的数存进去。
