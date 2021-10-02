# 405.数字转换为十六进制数
[https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal](https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal) 
## 原题
给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用 <a href="https://baike.baidu.com/item/%E8%A1%A5%E7%A0%81/6854613?fr=aladdin">补码运算</a> 方法。

 **注意:** 
- 十六进制中所有字母( `a-f` )都必须是小写。
- 十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符 `'0'` 来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。 
- 给定的数确保在32位有符号整数范围内。
-  **不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。** 
 **示例 1：** 

```

输入:
26

输出:
"1a"

```
 **示例 2：** 

```

输入:
-1

输出:
"ffffffff"

```
 
**标签**
`位运算` `数学` 


## 按位转换
```go
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	sb := &strings.Builder{}
	for i := 7; i >= 0; i-- {
		val := num >> (4 * i) & 0xf
		if val > 0 || sb.Len() > 0 {
			sb.WriteByte(getByte(val))
		}
	}
	return sb.String()
}

func getByte(a int) (ret byte) {
	if a < 10 {
		ret = '0' + byte(a)
	} else {
		ret = 'a' + byte(a-10)
	}
	return
}
```
>执行用时：0 ms
内存消耗：1.9 MB

四位一组取出数字，然后再进行转换。