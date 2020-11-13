# 剑指 Offer 67. 把字符串转换成整数
[https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/](https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/)
##
```go

func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	res := 0

	for len(str) > 0 && str[0] == ' ' {
		str = str[1:]
	}
	if len(str) == 0 {
		return 0
	}

	hasNeg := false
	if isNeg(str[0]) || isPos(str[0]) {
		if isNeg(str[0]) {
			hasNeg = true
		}
		str = str[1:]
	}
	if len(str) == 0 {
		return 0
	}

	for len(str) > 0 && str[0] == '0' {
		str = str[1:]
	}
	if len(str) == 0 {
		return 0
	}

	numcont := 0
	for numcont < len(str) && isNumber(str[numcont]) {
		numcont++
	}
	str = str[:numcont]
	if len(str) == 0 {
		return 0
	}
	if hasNeg {
		minInt := "2147483648"
		if nminInt := len(minInt); numcont > nminInt {
			return math.MinInt32
		} else if numcont == nminInt && str > minInt {
			return math.MinInt32
		}
		res = atoi(str)
		res = -res
	} else {
		maxInt := "2147483647"
		if nmaxInt := len(maxInt); numcont > nmaxInt {
			return math.MaxInt32
		} else if numcont == nmaxInt && str > maxInt {
			return math.MaxInt32
		}
		res = atoi(str)
	}
	return res
}
func atoi(str string) int {
	res := 0
	for i := range str {
		res *= 10
		res += int(str[i] - '0')
	}
	return res
}
func isPos(x byte) bool {
	return x == '+'
}
func isNeg(x byte) bool {
	return x == '-'
}
func isNumber(x byte) bool {
	return x >= '0' && x <= '9'
}
```
>执行用时：4 ms, 在所有 Go 提交中击败了53.89%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了68.33%的用户

思路简单，就是边界条件很多。 

1. 要去掉所有的空格
2. 去掉符号
3. 去掉前导 0 
4. 计算最长的数字，并检测溢出
5. 根据 2. 中的符号判断，看要不要加上负号

我这里主要的问题就是是直接对字符串进行的操作，所以1~3的每一步做完都要判断一次是否需要返回空。