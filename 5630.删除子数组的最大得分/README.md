# 5630. 删除子数组的最大得分
[https://leetcode-cn.com/problems/maximum-erasure-value/](https://leetcode-cn.com/problems/maximum-erasure-value/) 
## 原题
给你一个正整数数组 `nums` ，请你从中删除一个含有 **若干不同元素**  的子数组。删除子数组的 **得分**  就是子数组各元素之 **和**  。
返回 **只删除一个**  子数组可获得的 **最大得分** 。
如果数组 `b` 是数组 `a` 的一个连续子序列，即如果它等于 `a[l],a[l+1],...,a[r]` ，那么它就是 `a` 的一个子数组。
 
**示例 1：** 
```
输入：nums = [4,2,4,5,6]
输出：17
解释：最优子数组是 [2,4,5,6]
```
**示例 2：** 
```
输入：nums = [5,2,1,2,5,2,1,2,5]
输出：8
解释：最优子数组是 [5,2,1] 或 [1,2,5]
```
 
**提示：** 
- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^4`


## 用递归直接处理
```go
func reformatNumber(number string) string {
	res := make([]rune, 0, len(number))
	for _, char := range number {
		if char != ' ' && char != '-' {
			res = append(res, char)
		}
	}
	if len(res) < 3 {
		return string(res)
	} else if len(res) == 4 {
		tmp := res[2:]
		res = []rune(string(res[:2]))
		res = append(res, '-')
		res = append(res, tmp...)
		return string(res)
	}
	tmp := res[3:]
	res = []rune(string(res[:3]))
	res = append(res, '-')
	res = append(res, []rune(reformatNumber(string(tmp)))...)
	if res[len(res)-1] == '-' {
		res = res[:len(res)-1]
	}
	return string(res)
}
func test() {
	type TestType struct {
		number string
		want   string
	}
	ts := []TestType{
		TestType{
			number: "1-23-45 6",
			want:   "123-456",
		},
		TestType{
			number: "123 4-567",
			want:   "123-45-67",
		},
		TestType{
			number: "123 4-5678",
			want:   "123-456-78",
		},
		TestType{
			number: "12",
			want:   "12",
		},
		TestType{
			number: "--17-5 229 35-39475 ",
			want:   "175-229-353-94-75",
		},
	}
	for i, t := range ts {
		get := reformatNumber(t.number)
		if t.want != get {
			// 填充输出格式
			fmt.Printf("i:%d %+v get:%v\n", i, t, get)
		}
	}
	fmt.Println("Finished Test!")

}
```
>执行用时: 0 ms
内存消耗: 3 MB

就是先删空格和破折号，再根据题意拆分。为了简单于是就用了递归。
