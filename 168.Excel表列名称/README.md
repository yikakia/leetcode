# 168.Excel表列名称
[https://leetcode-cn.com/problems/excel-sheet-column-title](https://leetcode-cn.com/problems/excel-sheet-column-title) 
## 原题
给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

```
    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB 
    ...

```
 **示例 1:** 

```
输入: 1
输出: "A"

```
 **示例 2:** 

```
输入: 28
输出: "AB"

```
 **示例 3:** 

```
输入: 701
输出: "ZY"

```
 
**标签**
`数学` `字符串` 


## 
```go
func convertToTitle(columnNumber int) (res string) {
	for columnNumber > 0 {
		a0 := (columnNumber - 1) % 26
		res = string('A'+byte(a0)) + res
		columnNumber = (columnNumber - a0) / 26
	}
	return
}
```
>执行用时: 0 ms
内存消耗: 1.9 MB

坑点在于 是 26 进制，但是范围是 `1~26` 。