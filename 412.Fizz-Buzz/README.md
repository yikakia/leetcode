# 412.Fizz Buzz
[https://leetcode-cn.com/problems/fizz-buzz](https://leetcode-cn.com/problems/fizz-buzz) 
## 原题
写一个程序，输出从 1 到 *n* 数字的字符串表示。

1. 如果 *n* 是3的倍数，输出“Fizz”；

2. 如果 *n* 是5的倍数，输出“Buzz”；

3.如果 *n* 同时是3和5的倍数，输出 “FizzBuzz”。

 **示例：** 

```
n = 15,

返回:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]

```
 
**标签**
`数学` `字符串` `模拟` 


## 模拟
```go
func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := range ret {
		num := i + 1
		if (num)%3 != 0 && (num)%5 != 0 {
			ret[i] = fmt.Sprint(i + 1)
		}
		if (num)%3 == 0 {
			ret[i] = "Fizz"
		}
		if (num)%5 == 0 {
			ret[i] += "Buzz"
		}
	}
	return ret
}
```
>执行用时：4 ms
内存消耗：3.5 MB

就简单模拟一下即可。