# 5629. 重新格式化电话号码
[https://leetcode-cn.com/problems/reformat-phone-number/](https://leetcode-cn.com/problems/reformat-phone-number/) 
## 原题
给你一个字符串形式的电话号码 `number` 。`number` 由数字、空格 `' '`、和破折号 `'-'` 组成。
请你按下述方式重新格式化电话号码。
- 首先，**删除**  所有的空格和破折号。
- 其次，将数组从左到右 **每 3 个一组**  分块，**直到 ** 剩下 4 个或更少数字。剩下的数字将按下述规定再分块：
	
	- 2 个数字：单个含 2 个数字的块。
	- 3 个数字：单个含 3 个数字的块。
	- 4 个数字：两个分别含 2 个数字的块。
	
	
最后用破折号将这些块连接起来。注意，重新格式化过程中 **不应该**  生成仅含 1 个数字的块，并且 **最多**  生成两个含 2 个数字的块。
返回格式化后的电话号码。
 
**示例 1：** 
```
输入：number = "1-23-45 6"
输出："123-456"
解释：数字是 "123456"
步骤 1：共有超过 4 个数字，所以先取 3 个数字分为一组。第 1 个块是 "123" 。
步骤 2：剩下 3 个数字，将它们放入单个含 3 个数字的块。第 2 个块是 "456" 。
连接这些块后得到 "123-456" 。
```
**示例 2：** 
```
输入：number = "123 4-567"
输出："123-45-67"
解释：数字是 "1234567".
步骤 1：共有超过 4 个数字，所以先取 3 个数字分为一组。第 1 个块是 "123" 。
步骤 2：剩下 4 个数字，所以将它们分成两个含 2 个数字的块。这 2 块分别是 "45" 和 "67" 。
连接这些块后得到 "123-45-67" 。
```
**示例 3：** 
```
输入：number = "123 4-5678"
输出："123-456-78"
解释：数字是 "12345678" 。
步骤 1：第 1 个块 "123" 。
步骤 2：第 2 个块 "456" 。
步骤 3：剩下 2 个数字，将它们放入单个含 2 个数字的块。第 3 个块是 "78" 。
连接这些块后得到 "123-456-78" 。
```
**示例 4：** 
```
输入：number = "12"
输出："12"
```
**示例 5：** 
```
输入：number = "--17-5 229 35-39475 "
输出："175-229-353-94-75"
```
 
**提示：** 
- `2 <= number.length <= 100`
- `number` 由数字和字符 `'-'` 及 `' '` 组成。
- `number` 中至少含 **2**  个数字。


## 
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

按照规则来就行。不过为了简单所以用了递归。
