# 剑指 Offer 20. 表示数值的字符串
[https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)

## 

```go
func isNumber(s string) bool {
	newS := strings.TrimSpace(s)
	if len(newS) == 0 {
		return false
	}
	return isE(newS, 'e') ||
		isE(newS, 'E') ||
		isInteger(newS, true) ||
		isFloat(newS)
}

func isE(s string, c byte) bool {
	//有e
	e := strings.IndexByte(s, c)
	if e >= 0 {
		if e == 0 {
			return false
		}
		if e == len(s)-1 {
			return false
		}
		return (isInteger(s[0:e], true) || isFloat(s[0:e])) && isInteger(strings.TrimSpace(s[e+1:]), true)
	}
	return false
}

func isInteger(s string, flag bool) bool {
	//flag代表是不是严格判断
	//如果flag为false，则+、-能通过
	//如果flag为true，则+、-后必须有数字

	plus := strings.IndexByte(s, '+')
	if plus >= 0 {
		if plus > 0 {
			return false
		}
		if flag {
			if len(s) == 1 {
				return false
			}
		}
		return allNumber(s[1:])
	}

	minus := strings.IndexByte(s, '-')
	if minus >= 0 {
		if minus > 0 {
			return false
		}
		if flag {
			if len(s) == 1 {
				return false
			}
		}
		return allNumber(s[1:])
	}
	return allNumber(s)
}

func isFloat(s string) bool {
	if len(s) <= 1 {
		return false
	}
	dot := strings.IndexByte(s, '.')
	if dot >= 0 {
		if dot == 0 {
			return allNumber(s[dot+1:])
		}
		if dot == len(s)-1 {
			return isInteger(s[0:dot], true)
		}
		return isInteger(s[0:dot], false) && allNumber(s[dot+1:])
	}
	return false
}

func allNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
```
代码来自[https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/solution/biao-shi-shu-zhi-de-zi-fu-chuan-by-irenebyron-2/](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/solution/biao-shi-shu-zhi-de-zi-fu-chuan-by-irenebyron-2/)


写不出来，所以就复制了别人的。虽然编译原理讲过状态自动机，自己也写过，但是不看题解真的想不出来能用状态自动机来实现。接下来还是得看看怎么自己实现一次。还有就是补习一下怎么实现。