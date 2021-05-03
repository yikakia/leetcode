# 7.整数反转
[https://leetcode-cn.com/problems/reverse-integer](https://leetcode-cn.com/problems/reverse-integer) 
## 原题
给你一个 32 位的有符号整数 `x` ，返回将 `x` 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围  `[−2^31,  2^31 − 1]` ，就返回 0。
 **假设环境不允许存储 64 位整数（有符号或无符号）。** 

 

 **示例 1：** 

```

输入：x = 123
输出：321

```
 **示例 2：** 

```

输入：x = -123
输出：-321

```
 **示例 3：** 

```

输入：x = 120
输出：21

```
 **示例 4：** 

```

输入：x = 0
输出：0

```
 

 **提示：** 
-  `-2^31 <= x <= 2^31 - 1` 
 
**标签**
`数学` 


## 探测是否溢出
```go
func reverse(x int) int {
	var newX int32 = 0
	for x != 0 {
		X := x % 10
		x /= 10
		tNewX := newX * 10
		if tNewX/10 != newX {
			return 0
		}
		newX = tNewX

		tNewX = newX + int32(X)
		if tNewX-int32(X) != newX {
			return 0
		}
		newX = tNewX
	}
	return int(newX)
}
```
>执行用时: 0 ms
内存消耗: 2.1 MB

就是利用了溢出后计算不正常的特点，如果计算后和目标预期不同，则说明发生了溢出，这个时候直接返回 `0` 即可，不然的话就继续按照规则反转即可。